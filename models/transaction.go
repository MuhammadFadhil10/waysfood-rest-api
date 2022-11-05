package models

type Transaction struct {
	ID         int                  `json:"id" gorm:"primary_key:auto_increment"`
	Qty        int                  `json:"qty"`
	TotalPrice int                  `json:"totalPrice"`
	BuyerID    int                  `json:"buyer_id"`
	Buyer      UsersProfileResponse `json:"userOrder" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	SellerID   int                  `json:"seller_id"`
	Seller     UsersProfileResponse `json:"seller" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Status     string               `json:"status"`
	// ProductID int                  `json:"product_id" gorm:"type: int"`
	// Product   ProductResponse      `json:"order" gorm:"foreignKey:product_id;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}

// gorm:"foreignKey:product_id;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"
