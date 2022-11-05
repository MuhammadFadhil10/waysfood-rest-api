package models

type Order struct {
	ID            int                  `json:"id" gorm:"primary_key:auto_increment"`
	ProductID     int                  `json:"product_id" gorm:"type: int"`
	Products      ProductResponse      `json:"product" gorm:"foreignKey:product_id;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	BuyerID       int                  `json:"buyer_id"`
	Buyer         UsersProfileResponse `json:"userOrder" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	SellerID      int                  `json:"seller_id"`
	Seller        UsersProfileResponse `json:"seller" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	TransactionID int                  `json:"transactionId" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Transaction   Transaction          `json:"transaction" gorm:"foreignKey:transaction_id;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}
