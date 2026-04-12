package dto

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterResponse struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
