package repositories

import (
	"go-batch2/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProducts() ([]models.Product, error)
	GetProductByID(ID int) (models.Product, error)
	GetProductByPartner(userID int) ([]models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	UpdateProduct(product models.Product, ID int) (models.Product, error)
	DeleteProduct(product models.Product, ID int) (models.Product, error)
}

func RepositoryProduct(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("User").Find(&products).Error

	return products, err
}

func (r *repository) GetProductByID(ID int) (models.Product, error) {
	var product models.Product
	err := r.db.Preload("User").First(&product, ID).Error

	return product, err
}

func (r *repository) GetProductByPartner(userID int) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("User").Where("user_id = ?", userID).Find(&products).Error

	return products, err
}

func (r *repository) CreateProduct(product models.Product) (models.Product, error) {
	err := r.db.Create(&product).Error
	return product, err
}

func (r *repository) UpdateProduct(Product models.Product, ID int) (models.Product, error) {
	err := r.db.Model(&Product).Where("id=? ", ID).Updates(&Product).Error

	return Product, err
}

func (r *repository) DeleteProduct(Product models.Product, ID int) (models.Product, error) {
	err := r.db.Delete(&Product).Error

	return Product, err
}
