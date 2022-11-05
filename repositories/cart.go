package repositories

import (
	"go-batch2/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	AddToCart(cart models.Cart) (models.Cart, error)
	GetCartByID(ID int) (models.Cart, error)
	GetChartByUserID(userID int) ([]models.Cart, error)
	GetChartByUser(userID int, productID int) (models.Cart, error)
	GetChartByProductID(productID int) ([]models.Cart, error)
	UpdateCartQty(Cart models.Cart, userID int, productID int) (models.Cart, error)
	DeleteCartByID(Cart models.Cart, ID int) (models.Cart, error)
	DeleteAllCart(Cart models.Cart, userID int) (models.Cart, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AddToCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Create(&cart).Error
	return cart, err
}

func (r *repository) GetCartByID(ID int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.First(&cart, ID).Preload("User").Preload("Products").Error
	return cart, err
}

func (r *repository) GetChartByUser(userID int, productID int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("Users").Preload("Products").Where("users_id = ? and product_id=?", userID, productID).First(&cart).Error
	return cart, err
}

func (r *repository) GetChartByUserID(userID int) ([]models.Cart, error) {
	var cart []models.Cart
	err := r.db.Preload("Users").Preload("Products.User").Where("users_id=?", userID).Find(&cart).Error
	return cart, err
}

func (r *repository) GetChartByProductID(productID int) ([]models.Cart, error) {
	var cart []models.Cart
	err := r.db.Preload("Users").Preload("Products.User").Where("products_id=?", productID).Find(&cart).Error
	return cart, err
}

func (r *repository) UpdateCartQty(Cart models.Cart, userID int, productID int) (models.Cart, error) {
	err := r.db.Model(&Cart).Where("users_id=? and product_id=?", userID, productID).Updates(&Cart).Error
	return Cart, err
}

func (r *repository) DeleteCartByQty(Cart models.Cart, userID int, productID int) (models.Cart, error) {
	err := r.db.Model(&Cart).Where("users_id=? and product_id=?", userID, productID).Updates(&Cart).Error
	return Cart, err
}

func (r *repository) DeleteCartByID(Cart models.Cart, ID int) (models.Cart, error) {
	err := r.db.Delete(&Cart, ID).Error
	return Cart, err
}
func (r *repository) DeleteAllCart(Cart models.Cart, userID int) (models.Cart, error) {
	err := r.db.Preload("User").Preload("Products.User").Where("users_id = ?", userID).Delete(&Cart).Error
	return Cart, err
}
