package repositories

import (
	"fmt"
	"go-batch2/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	ShowTransaction() ([]models.Transaction, error)
	GetTransactionByID(ID int) (models.Transaction, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	UpdateTransaction(transaction models.Transaction, ID int) (models.Transaction, error)
	DeleteTransaction(transaction models.Transaction, ID int) (models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) ShowTransaction() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("Users").Preload("Product").Find(&transactions).Error

	return transactions, err
}

func (r *repository) GetTransactionByID(ID int) (models.Transaction, error) {
	var transactions models.Transaction
	err := r.db.Preload("Users").Preload("Product").First(&transactions, ID).Error

	return transactions, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error

	if err != nil {
		fmt.Println(err)
	}

	return transaction, err
}

func (r *repository) UpdateTransaction(transaction models.Transaction, ID int) (models.Transaction, error) {
	err := r.db.Model(&transaction).Where("id=?", ID).Updates(&transaction).Error

	return transaction, err
}

func (r *repository) DeleteTransaction(transaction models.Transaction, ID int) (models.Transaction, error) {
	err := r.db.Delete(&transaction).Error

	return transaction, err
}
