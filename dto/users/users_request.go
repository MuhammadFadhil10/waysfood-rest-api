package usersdto

type CreateUserRequest struct {
	Email    string `json:"email" form:"email" gorm:"type: varchar(255)" validate:"required"`
	Password string `json:"password" form:"password" gorm:"type: varchar(255)" validate:"required"`
	FullName string `json:"fullname" form:"fullname" gorm:"type: varchar(255)" validate:"required"`
	Gender   string `json:"gender" form:"gender" gorm:"type: varchar(255)" validate:"required"`
	Phone    string `json:"phone" form:"phone" gorm:"type: varchar(255)" validate:"required"`
	Role     string `json:"role" form:"role" gorm:"type: varchar(255)" validate:"required"`
}

type UpdateUserRequest struct {
	Email    string `json:"email" form:"email" gorm:"type: varchar(255)"`
	Password string `json:"password" form:"password" gorm:"type: varchar(255)"`
	FullName string `json:"fullName" form:"name" gorm:"type: varchar(255)"`
	Gender   string `json:"gender" form:"gender" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" form:"phone" gorm:"type: varchar(255)"`
	Location string `json:"location" form:"location" gorm:"type: varchar(255)"`
	Image    string `json:"image" form:"image" gorm:"type: varchar(255)"`
}