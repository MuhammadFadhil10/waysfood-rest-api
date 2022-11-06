package handlers

import (
	"encoding/json"
	"fmt"
	dto "go-batch2/dto/result"
	transactiondto "go-batch2/dto/transaction"
	"go-batch2/models"
	"go-batch2/repositories"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository}
}

func (h *handlerTransaction) GetTransactionByPartner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	partnerId, _ := strconv.Atoi(mux.Vars(r)["partnerId"])

	transaction := []models.Transaction{}

	myTransaction, err := h.TransactionRepository.GetTransactionByPartnerID(transaction, partnerId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: h.convertTransactionResponse(myTransaction)}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerTransaction) GetTransactionByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	transaction := []models.Transaction{}

	myTransaction, err := h.TransactionRepository.GetTransactionByUserID(transaction, userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: h.convertTransactionResponse(myTransaction)}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerTransaction) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var c = coreapi.Client{
		ServerKey: os.Getenv("SERVER_KEY"),
		ClientKey: os.Getenv("CLIENT_KEY"),
	}

	fmt.Println(c)

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	request := new(transactiondto.CreateTransactionRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// create unique transaction id
	var TransIdIsMatch = false
	var TransactionId int
	for !TransIdIsMatch {
		TransactionId = userId + request.SellerID + rand.Intn(10000) - rand.Intn(100)
		transactionData, _ := h.TransactionRepository.GetTransactionByID(TransactionId)
		if transactionData.ID == 0 {
			TransIdIsMatch = true
		}
	}

	transaction := models.Transaction{
		BuyerID:    userId,
		SellerID:   request.SellerID,
		Status:     request.Status,
		Qty:        request.Qty,
		TotalPrice: request.TotalPrice,
	}

	validation := validator.New()
	validateErr := validation.Struct(request)
	if validateErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "Failed", Message: validateErr.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var createTransactionErr error
	transaction, createTransactionErr = h.TransactionRepository.CreateTransaction(transaction)
	if createTransactionErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "Failed", Message: createTransactionErr.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// 1. Initiate Snap client
	var s = snap.Client{}
	s.New(c.ServerKey, midtrans.Sandbox)
	key := os.Getenv("SERVER_KEY")
	// Use to midtrans.Production if you want Production Environment (accept real transaction).
	fmt.Println("ypour key server:", key)
	// 2. Initiate Snap request param
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.TotalPrice),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: transaction.Buyer.FullName,
			Email: transaction.Buyer.Email,
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, _ := s.CreateTransaction(req)

	// add order product list
	userCart, err := h.TransactionRepository.FindChartByUserID(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var order models.Order

	for _, c := range userCart {
		fmt.Println("transaction id:", transaction.ID)
		order.ProductID = c.ProductID
		order.BuyerID = userId
		order.SellerID = request.SellerID
		order.TransactionID = transaction.ID
		err := h.TransactionRepository.CreateTransactionOrder(order)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	var cart models.Cart

	h.TransactionRepository.DeleteFromCart(cart,userId)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: snapResp}
	json.NewEncoder(w).Encode(response)
}

// NOTIFICATION
func (h *handlerTransaction) Notification(w http.ResponseWriter, r *http.Request) {
	var notificationPayload map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&notificationPayload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	transactionStatus := notificationPayload["transaction_status"].(string)
	fraudStatus := notificationPayload["fraud_status"].(string)
	orderId := notificationPayload["order_id"].(string)

	if transactionStatus == "capture" {
		if fraudStatus == "challenge" {
			// TODO set transaction status on your database to 'challenge'
			// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
			h.TransactionRepository.UpdateTransaction("pending", orderId)
		} else if fraudStatus == "accept" {
			// TODO set transaction status on your database to 'success'
			h.TransactionRepository.UpdateTransaction("success", orderId)
		}
	} else if transactionStatus == "settlement" {
		// TODO set transaction status on your databaase to 'success'
		h.TransactionRepository.UpdateTransaction("success", orderId)
	} else if transactionStatus == "deny" {
		// TODO you can ignore 'deny', because most of the time it allows payment retries
		// and later can become success
		h.TransactionRepository.UpdateTransaction("failed", orderId)
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
		// TODO set transaction status on your databaase to 'failure'
		h.TransactionRepository.UpdateTransaction("failed", orderId)
	} else if transactionStatus == "pending" {
		// TODO set transaction status on your databaase to 'pending' / waiting payment
		h.TransactionRepository.UpdateTransaction("pending", orderId)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handlerTransaction) convertTransactionResponse(t []models.Transaction) []transactiondto.GetTransactionResponse {
	var order []models.Order
	var resp []transactiondto.GetTransactionResponse

	// var orderList [] models.Order

	for _, item := range t {
		test, err := h.TransactionRepository.GetTransactionProducts(order, item.ID)
		fmt.Println(err)
		resp = append(resp, transactiondto.GetTransactionResponse{
			ID:        item.ID,
			Qty:       item.Qty,
			Buyer:     item.Buyer,
			Seller:    item.Seller,
			Status:    item.Status,
			OrderList: test,
			TotalPrice: item.TotalPrice,
		})
	}

	fmt.Println(resp)

	return resp

}
