package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gustavoarendt/jobtracker/internal/database"
	"github.com/gustavoarendt/jobtracker/internal/dto"
	"github.com/gustavoarendt/jobtracker/internal/entities"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInputModel
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := entities.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err = database.NewUser(database.DB).CreateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user dto.LoginInputModel
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := database.NewUser(database.DB).FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if !u.ComparePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, token, _ := database.Jwt.Encode(map[string]interface{}{
		"sub": u.ID,
		"exp": time.Now().Add(time.Minute * time.Duration(database.JwtExpiresIn)).Unix(),
	})

	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: token,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}
