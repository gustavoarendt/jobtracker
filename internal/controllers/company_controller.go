package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gustavoarendt/jobtracker/internal/database"
	"github.com/gustavoarendt/jobtracker/internal/dto"
	"github.com/gustavoarendt/jobtracker/internal/entities"
)

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	var company dto.CreateCompanyInputModel
	err := json.NewDecoder(r.Body).Decode(&company)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	c, err := entities.NewCompany(company.Name, company.Description, company.Website_url, company.Linkedin_url, company.Image_url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err = database.NewCompany(database.DB).Create(c)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func GetCompanies(w http.ResponseWriter, r *http.Request) {
	companies, err := database.NewCompany(database.DB).FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var companiesViewModel []dto.CompanyViewModel
	for _, c := range companies {
		companiesViewModel = append(companiesViewModel, dto.CompanyViewModel{
			ID:           c.ID,
			Name:         c.Name,
			Description:  c.Description,
			Website_url:  c.Website_url,
			Linkedin_url: c.Linkedin_url,
			Image_url:    c.Image_url,
			Created_at:   c.Created_at.String(),
		})
	}
	json.NewEncoder(w).Encode(companiesViewModel)
}

func GetCompany(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	company, err := database.NewCompany(database.DB).FindById(idUint)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(dto.CompanyViewModel{
		ID:           company.ID,
		Name:         company.Name,
		Description:  company.Description,
		Website_url:  company.Website_url,
		Linkedin_url: company.Linkedin_url,
		Image_url:    company.Image_url,
		Created_at:   company.Created_at.String(),
	})
}

func UpdateCompany(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var company dto.CreateCompanyInputModel
	err := json.NewDecoder(r.Body).Decode(&company)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	c, err := entities.NewCompany(company.Name, company.Description, company.Website_url, company.Linkedin_url, company.Image_url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	c.ID = idUint

	err = database.NewCompany(database.DB).Update(c)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func DeleteCompany(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := database.NewCompany(database.DB).Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
