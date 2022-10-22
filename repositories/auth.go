package repositories

import (
	"go-batch2/models"

	"gorm.io/gorm"
)

// import (
// 	"waysfood/models"

// 	"gorm.io/gorm"
// )

// type repository struct {
// 	db *gorm.DB
// }

type AuthRepository interface {
	Register(user models.User) (models.User, error)
	Login(email string) (models.User, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Register(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *repository) Login(email string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "email=?", email).Error

	return user, err
}