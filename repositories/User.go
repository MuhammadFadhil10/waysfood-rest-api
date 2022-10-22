package repositories

import (
	"go-batch2/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user models.User) (models.User ,error)
	GetUsers() ([]models.User, error)
	FindUserById(ID int) (models.User, error) 
	UpdateUser(user models.User, ID int) (models.User, error)
	DeleteUser(user models.User,ID int) (models.User, error)
}

type repository struct {
	db *gorm.DB
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) GetUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Raw("SELECT * FROM users").Scan(&users).Error

  	return users, err
}

func (r *repository) FindUserById(ID int) (models.User, error) {
	var user models.User

	err := r.db.First(&user, ID).Error

	return user, err
}

func (r *repository) UpdateUser(user models.User, ID int) (models.User, error) {

	err := r.db.Model(&user).Where("id=?", ID).Updates(&user).Error

	return user, err
}  

func (r *repository) DeleteUser( user models.User ,ID int) (models.User, error) {
	err := r.db.Delete(&user, ID).Error

	return user, err
}