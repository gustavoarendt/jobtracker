package dto

type CreateCompanyInputModel struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Website_url  string `json:"website_url"`
	Linkedin_url string `json:"linkedin_url"`
	Image_url    string `json:"image_url"`
}

type CompanyViewModel struct {
	ID           uint64 `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Website_url  string `json:"website_url"`
	Linkedin_url string `json:"linkedin_url"`
	Image_url    string `json:"image_url"`
	Created_at   string `json:"created_at"`
}
