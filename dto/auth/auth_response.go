package authdto

type AuthResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}
