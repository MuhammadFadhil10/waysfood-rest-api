package models

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
	Phone    string `json:"phone"`
	Location string `json:"location"`
	Image    string `json:"image"`
	Role     string `json:"role"`
	Gender   string `json:"gender"`
}

type UsersProfileResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName"`
	Image    string `json:"image"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Location string `json:"location"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
