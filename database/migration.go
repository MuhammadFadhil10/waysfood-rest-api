package database

import (
	"fmt"
	"go-batch2/models"
	"go-batch2/pkg/mysql"
)

func RunMigration() {
	if err := mysql.DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Transaction{}, &models.Cart{}, &models.Order{}); err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
