package models

type Transaction struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	Qty       int                  `json:"qty"`
	UsersID   int                  `json:"user_id"`
	Users     UsersProfileResponse `json:"users"`
	Status    string               `json:"status"`
	ProductID int                  `json:"product_id" gorm:"type: int"`
	Product   ProductResponse      `json:"product"`
}
