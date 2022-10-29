package authdto

type AuthResponse struct {
	ID int `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}
