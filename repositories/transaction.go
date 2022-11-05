package repositories

import (
	"fmt"
	"go-batch2/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	ShowTransaction() ([]models.Transaction, error)
	GetTransactionByID(ID int) (models.Transaction, error)
	GetTransactionByUserID(transaction []models.Transaction, userID int) ([]models.Transaction, error)
	GetTransactionProducts(order []models.Order, transactionID int) ([]models.Order, error)
	GetTransactionByPartnerID(transaction []models.Transaction, sellerId int) ([]models.Transaction, error)
	CreateTransactionOrder(order models.Order) error
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	FindChartByUserID(userID int) ([]models.Cart, error)
	UpdateTransaction(transaction models.Transaction, ID int) (models.Transaction, error)
	DeleteTransaction(transaction models.Transaction, ID int) (models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) ShowTransaction() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("Users").Preload("Products").Find(&transactions).Error

	return transactions, err
}

func (r *repository) GetTransactionByID(ID int) (models.Transaction, error) {
	var transactions models.Transaction
	err := r.db.Preload("Buyer").Preload("Seller").First(&transactions, ID).Error

	return transactions, err
}

func (r *repository) GetTransactionByUserID(transaction []models.Transaction, userID int) ([]models.Transaction, error) {
	err := r.db.Preload("Buyer").Preload("Seller").Where("buyer_id=?", userID).Find(&transaction).Error

	return transaction, err
}

func (r *repository) GetTransactionByPartnerID(transaction []models.Transaction, sellerId int) ([]models.Transaction, error) {
	err := r.db.Preload("Buyer").Preload("Seller").Where("seller_id=?", sellerId).Find(&transaction).Error

	return transaction, err
}

func (r *repository) GetTransactionProducts(order []models.Order, transactionID int) ([]models.Order, error) {
	err := r.db.Preload("Products").Preload("Buyer").Preload("Seller").Where("transaction_id=?", transactionID).Find(&order).Error
	fmt.Println(err)
	return order, err
}

func (r *repository) FindChartByUserID(userID int) ([]models.Cart, error) {
	var cart []models.Cart
	err := r.db.Preload("Users").Preload("Products.User").Where("users_id=?", userID).Find(&cart).Error
	return cart, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error

	if err != nil {
		fmt.Println(err)
	}

	return transaction, err
}

func (r *repository) CreateTransactionOrder(order models.Order) error {
	err := r.db.Create(&order).Error

	return err
}

func (r *repository) UpdateTransaction(transaction models.Transaction, ID int) (models.Transaction, error) {
	err := r.db.Model(&transaction).Where("id=?", ID).Updates(&transaction).Error

	return transaction, err
}

func (r *repository) DeleteTransaction(transaction models.Transaction, ID int) (models.Transaction, error) {
	err := r.db.Delete(&transaction).Error

	return transaction, err
}
