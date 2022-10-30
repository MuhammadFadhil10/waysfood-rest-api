package cartdto

type CreateCartRequest struct {
	ProductID int `json:"productId"  gorm:"type: int"`
	UserID    int `json:"userId"  gorm:"type: int"`
}	
