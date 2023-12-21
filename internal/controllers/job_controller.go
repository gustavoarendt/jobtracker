package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/gustavoarendt/jobtracker/internal/database"
	"github.com/gustavoarendt/jobtracker/internal/dto"
	"github.com/gustavoarendt/jobtracker/internal/entities"
)

func CreateJob(w http.ResponseWriter, r *http.Request) {
	userId := getUserIdFromContext(r)
	var job dto.CreateJobInputModel
	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	j, err := entities.NewJob(job.Name, job.Description, job.Status, job.Currency, job.Language, job.Id_company, userId, job.Expected_salary, job.Interest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err = database.NewJob(database.DB).Create(j)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func GetJobs(w http.ResponseWriter, r *http.Request) {
	userId := getUserIdFromContext(r)
	jobs, err := database.NewJob(database.DB).FindAll(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var jobsViewModel []dto.JobViewModel
	for _, j := range jobs {
		jobsViewModel = append(jobsViewModel, dto.JobViewModel{
			ID:              j.ID,
			Name:            j.Name,
			Description:     j.Description,
			Status:          j.Status,
			Currency:        j.Currency,
			Language:        j.Language,
			Id_company:      j.Id_company,
			Expected_salary: j.Expected_salary,
			Interest:        j.Interest,
			Created_at:      j.Created_at.Format("02/01/2006 15:04:05"),
			Updated_at:      j.Updated_at.Format("02/01/2006 15:04:05"),
		})
	}
	json.NewEncoder(w).Encode(jobsViewModel)
}

func GetJob(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userId := getUserIdFromContext(r)
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	job, err := database.NewJob(database.DB).FindById(idUint)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if job.Id_user != userId {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	json.NewEncoder(w).Encode(job)
}

func UpdateJob(w http.ResponseWriter, r *http.Request) {
	userId := getUserIdFromContext(r)
	var job dto.CreateJobInputModel
	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	j, err := entities.NewJob(job.Name, job.Description, job.Status, job.Currency, job.Language, job.Id_company, userId, job.Expected_salary, job.Interest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err = database.NewJob(database.DB).Update(j)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func DeleteJob(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := database.NewJob(database.DB).Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func getUserIdFromContext(r *http.Request) uint64 {
	_, claims, _ := jwtauth.FromContext(r.Context())
	userIdString := fmt.Sprintf("%v", claims["user_id"])
	userId, _ := strconv.ParseUint(userIdString, 10, 64)
	return userId
}
