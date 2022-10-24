package models

type Product struct {
	ID     int                  `json:"id" gorm:"primary_key:auto_increment"`
	Title  string               `json:"title" form:"title" gorm:"type: varchar(255)"`
	Price  int                  `json:"price" form:"price" gorm:"type: int"`
	Image  string               `json:"image" form:"image" gorm:"type: varchar(255)"`
	Qty    int                  `json:"qty" form:"qty"`
	UserID int                  `json:"user_id" form:"user_id"`
	User   UsersProfileResponse `json:"user" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}

type ProductResponse struct {
	ID     int                  `json:"id"`
	Title  string               `json:"name"`
	Price  int                  `json:"price"`
	Image  string               `json:"image"`
	Qty    int                  `json:"qty"`
	UserID int                  `json:"-"`
	User   UsersProfileResponse `json:"user"`
}

type ProductUserResponse struct {
	ID     int    `json:"id"`
	Title  string `json:"name"`
	Price  int    `json:"price"`
	Image  string `json:"image"`
	Qty    int    `json:"qty"`
	UserID int    `json:"-"`
}

func (ProductResponse) TableName() string {
	return "products"
}

func (ProductUserResponse) TableName() string {
	return "products"
}
