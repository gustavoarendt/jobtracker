package dto

type CreateUserInputModel struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginInputModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
