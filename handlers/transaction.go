package handlers

import (
	"encoding/json"
	"fmt"
	dto "go-batch2/dto/result"
	transactiondto "go-batch2/dto/transaction"
	"go-batch2/models"
	"go-batch2/repositories"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository}
}

// func (h *handlerTransaction) GetAllTransaction(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	transaction, err := h.TransactionRepository.ShowTransaction()
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Status: "Success", Data: convertTransactionResponse(transaction)}
// 	json.NewEncoder(w).Encode(response)
// }

// func (h *handlerTransaction) GetTransactionByID(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])

// 	var transaction models.Transaction
// 	transaction, err := h.TransactionRepository.GetTransactionByID(id)
// 	if err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Status: "Success", Data: transaction}
// 	json.NewEncoder(w).Encode(response)
// }

func (h *handlerTransaction) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	request := new(transactiondto.CreateTransactionRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// var c repositories.CartRepository

	// fmt.Println(c)

	userCart, err := h.TransactionRepository.FindChartByUserID(userId)

	fmt.Println(err)
	fmt.Println(userCart)
	fmt.Println(len(userCart))

	var order models.Order
	// var addOrder models.Order

	for _, c := range userCart {
		order.ProductID = c.ProductID
		order.BuyerID = userId
		order.SellerID = request.SellerID
		err := h.TransactionRepository.CreateTransactionOrder(order)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	transaction := models.Transaction{
		BuyerID:  userId,
		SellerID: request.SellerID,
		Status:   request.Status,
		Qty:      request.Qty,
	}

	validation := validator.New()
	validateErr := validation.Struct(request)
	if validateErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "Failed", Message: validateErr.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	transaction, err = h.TransactionRepository.CreateTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	transaction, _ = h.TransactionRepository.GetTransactionByID(transaction.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: transaction}
	json.NewEncoder(w).Encode(response)
}

// func (h *handlerTransaction) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	request := new(transactiondto.UpdateTransactionRequest)
// 	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])

// 	transaction := models.Transaction{}

// 	transaction.ID = id

// 	if request.Status != "" {
// 		transaction.Status = request.Status
// 	}

// 	if request.ProductID != 0 {
// 		transaction.ProductID = request.ProductID
// 	}

// 	if request.Qty != 0 {
// 		transaction.Qty = request.Qty
// 	}

// 	data, err := h.TransactionRepository.UpdateTransaction(transaction, id)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Status: "Success", Data: data}
// 	json.NewEncoder(w).Encode(response)
// }

// func (h *handlerTransaction) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])

// 	transaction, err := h.TransactionRepository.GetTransactionByID(id)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	data, err := h.TransactionRepository.DeleteTransaction(transaction, id)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Status: "Success", Data: data}
// 	json.NewEncoder(w).Encode(response)
// }

// func convertTransactionResponse(u []models.Transaction) []transactiondto.TransactionResponse {

// 	var products []models.ProductResponse
// 	var resp []transactiondto.TransactionResponse

// 	for _, r := range u {
// 		products = append(products, r.Product)
// 		resp = append(resp, transactiondto.TransactionResponse{
// 			ID:      r.ID,
// 			Users:   r.Users,
// 			Status:  r.Status,
// 			Product: products,
// 		})
// 	}
// 	return resp
// }
