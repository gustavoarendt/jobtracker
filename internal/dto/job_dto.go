package dto

type CreateJobInputModel struct {
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Status          string  `json:"status"`
	Currency        string  `json:"currency"`
	Language        string  `json:"language"`
	Id_company      uint64  `json:"id_company"`
	Id_user         uint64  `json:"id_user"`
	Expected_salary float64 `json:"expected_salary"`
	Interest        int     `json:"interest"`
}

type JobViewModel struct {
	ID              uint64  `json:"id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Status          string  `json:"status"`
	Currency        string  `json:"currency"`
	Language        string  `json:"language"`
	Id_company      uint64  `json:"id_company"`
	Expected_salary float64 `json:"expected_salary"`
	Interest        int     `json:"interest"`
	Created_at      string  `json:"created_at"`
	Updated_at      string  `json:"updated_at"`
}
