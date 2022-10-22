package transactiondto

type TransactionResponse struct {
	ID        int    `json:"id"`
	Qty       int    `json:"qty"`
	UsersID   int    `json:"user_id"`
	Status    string `json:"status"`
	ProductID int    `json:"product_id"`
}
