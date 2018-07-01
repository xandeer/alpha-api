package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	. "github.com/xandeer/alpha-api/config"
	. "github.com/xandeer/alpha-api/dao"
	. "github.com/xandeer/alpha-api/models"
)

var config = Config{}
var dao = DAO{}
var latestOperateVersion int

func InitOperateVersion() error {
	version, err := dao.GetLatestOperateVersion()
	latestOperateVersion = version
	return err
}

// GET list of items
func AllItems(w http.ResponseWriter, r *http.Request) {
	items, err := dao.FindItemsAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, items)
}

// GET an item by its ID
func FindItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	item, err := dao.FindItemById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Item ID")
		return
	}
	respondWithJson(w, http.StatusOK, item)
}

// POST a new item
func CreateItem(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.InsertItem(item); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	insertOperate(item, INSERT)
	respondWithJson(w, http.StatusCreated, item)
}

// PUT update an existing item
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.UpdateItem(item); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	insertOperate(item, UPDATE)
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing item
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	item, err := dao.FindItemById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Item ID")
		return
	}
	item.Removed = true
	if err := dao.UpdateItem(item); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	insertOperate(item, DELETE)
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func OperateVersion(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, http.StatusOK, latestOperateVersion)
}

func PullOperates(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var version int
	var err error
	if version, err = strconv.Atoi(params["version"]); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Operate version")
		return
	}
	operates, err := dao.FindOperatesGreaterThan(version)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Operate version")
		return
	}
	var ret []Operate

	for _, o1 := range operates {
		for j, o2 := range ret {
			if o1.Data.ID == o2.Data.ID {
				ret = append(ret[:j], ret[j+1:]...)
			}
		}
		ret = append(ret, o1)
	}
	respondWithJson(w, http.StatusOK, ret)
}

func insertOperate(item Item, operateType OperateType) {
	var operate Operate
	operate.Data = item
	operate.Type = operateType
	latestOperateVersion++
	operate.Version = latestOperateVersion

	dao.InsertOperate(operate)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}
