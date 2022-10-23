package transactiondto

import "go-batch2/models"

type TransactionResponse struct {
	ID     int                         `json:"id"`
	Users  models.UsersProfileResponse `json:"userOrder"`
	Status string                      `json:"status"`
	// Qty       int                  `json:"qty"`
	// UsersID   int                  `json:"user_id"`
	// ProductID int                  `json:"product_id" gorm:"type: int"`
	Product []models.ProductResponse `json:"order"`
}
