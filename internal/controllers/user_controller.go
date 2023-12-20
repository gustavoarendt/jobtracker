package controllers

import (
	"encoding/json"
	"net/http"

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
		return
	}
	err = database.DB.Create(u).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
