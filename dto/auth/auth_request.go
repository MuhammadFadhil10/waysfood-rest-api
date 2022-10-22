package authdto

type AuthRequest struct {
	Email    string `gorm:"type: varchar(255)" form:"email" json:"email" validate:"required"`
	Password string `gorm:"type: varchar(255)" form:"password" json:"password" validate:"required"`
}