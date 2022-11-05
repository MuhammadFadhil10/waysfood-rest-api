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

type GetTransactionResponse struct {
	ID        int                         `json:"id" gorm:"primary_key:auto_increment"`
	Qty       int                         `json:"qty"`
	Buyer     models.UsersProfileResponse `json:"buyer"`
	Seller    models.UsersProfileResponse `json:"seller" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Status    string                      `json:"status"`
	OrderList []models.Order              `json:"orderList"`
}

// type Transaction struct {
// 	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
// 	Qty       int                  `json:"qty"`
// 	BuyerID   int                  `json:"buyer_id"`
// 	Buyer     UsersProfileResponse `json:"userOrder" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
// 	SellerID  int                  `json:"seller_id"`
// 	Seller    UsersProfileResponse `json:"seller" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
// 	Status    string               `json:"status"`
// 	// ProductID int                  `json:"product_id" gorm:"type: int"`
// 	// Product   ProductResponse      `json:"order" gorm:"foreignKey:product_id;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
// }
