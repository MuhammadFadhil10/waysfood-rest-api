package handlers

import (
	"encoding/json"
	"fmt"
	dto "go-batch2/dto/result"
	"go-batch2/models"
	"go-batch2/repositories"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerCart struct {
	CartRepository repositories.CartRepository
}

func HandlerCart(CartRepository repositories.CartRepository) *handlerCart {
	return &handlerCart{CartRepository}
}

func (h *handlerCart) AddToCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	productId, _ := strconv.Atoi(mux.Vars(r)["productID"])

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	var cartModels models.Cart
	if err := json.NewDecoder(r.Body).Decode(&cartModels); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var cart models.Cart
	cartExist, err := h.CartRepository.GetChartByUser(userId, productId)

	if err == nil {
		cartModels.UsersID = userId
		cartModels.ProductID = productId
		cartModels.Qty = cartExist.Qty + 1
		cartModels.Price = cartModels.Price + cartExist.Price
		// update cart
		cart, err = h.CartRepository.UpdateCartQty(cartModels, userId, productId)
		fmt.Println(cart)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

	} else {
		cartModels.UsersID = userId
		cartModels.ProductID = productId
		cartModels.Qty = 1

		fmt.Println(cartModels)

		cart, err = h.CartRepository.AddToCart(cartModels)
		fmt.Println(cart)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: cart}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerCart) GetChartByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	cart, err := h.CartRepository.GetChartByUserID(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: cart}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerCart) DeleteChartByQty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	productId, _ := strconv.Atoi(mux.Vars(r)["productID"])

	cartItem, _ := h.CartRepository.GetChartByUser(userId, productId)

	var cartUpdate models.Cart

	var cart models.Cart
	var updateErr error

	cartUpdate.Qty = cartItem.Qty - 1
	cartUpdate.Price = cartItem.Price - (cartItem.Price / cartItem.Qty)
	cart, updateErr = h.CartRepository.UpdateCartQty(cartUpdate, userId, productId)

	if updateErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "Failed", Message: updateErr.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: cart}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCart) DeleteChartByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	productId, _ := strconv.Atoi(mux.Vars(r)["productID"])

	cartItem, _ := h.CartRepository.GetChartByUser(userId, productId)

	cartDelete, err := h.CartRepository.DeleteCartByID(cartItem, cartItem.ID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: cartDelete}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCart) DeleteAllCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// cartItem, _ := h.CartRepository.GetChartByUser(userId, productId)

	cart := models.Cart{}

	cartDelete, err := h.CartRepository.DeleteAllCart(cart, userId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: cartDelete}
	json.NewEncoder(w).Encode(response)
}
