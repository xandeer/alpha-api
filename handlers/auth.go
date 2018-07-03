package handlers

import (
	"encoding/json"
	"net/http"

	. "github.com/xandeer/alpha-api/models"
	"github.com/xandeer/alpha-api/auth"
)

func Signin(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	userM, err := dao.FindUserByName(user.Name)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user name")
		return
	}

	if user.Password != userM.Password {
		respondWithError(w, http.StatusForbidden, "Invalid user name or password")
		return
	}

	ss, err := auth.CreateToken(user.Name)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error while Signing Token")
		return
	}

	respondWithJson(w, http.StatusOK, ss)
}
