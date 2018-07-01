package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	. "github.com/xandeer/alpha-api/models"
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
		respondWithError(w, http.StatusForbidden, "Invalid use name or password")
		return
	}

	str, _ := json.Marshal(userM)
	key := []byte(str)

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
		Issuer:    "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)

	respondWithJson(w, http.StatusOK, ss)
}
