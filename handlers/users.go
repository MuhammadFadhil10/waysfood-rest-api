package handlers

import (
	"encoding/json"
	"fmt"
	dto "go-batch2/dto/result"
	usersdto "go-batch2/dto/users"
	"go-batch2/models"
	"go-batch2/repositories"
	"net/http"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handler struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handler {
	return &handler{UserRepository}
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(usersdto.CreateUserRequest)

	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	userModel := models.User{
		FullName: request.FullName,
		Email:    request.Email,
		Password: request.Password,
	}

	user, err := h.UserRepository.CreateUser(userModel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: user}
	json.NewEncoder(w).Encode(response)
}

func (h *handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	role := r.URL.Query()["role"]

	var (
		users   []models.User
		userErr error
	)

	if len(role) > 0 && role[0] == "partner" {
		users, userErr = h.UserRepository.GetPartners(role[0])
	} else {
		users, userErr = h.UserRepository.GetUsers()
	}

	if userErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "Failed", Message: userErr.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: convertUsersResponse(users)}
	json.NewEncoder(w).Encode(response)

}

func (h *handler) FindUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := h.UserRepository.FindUserById(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: user}
	json.NewEncoder(w).Encode(response)
}

func (h *handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	user, err := h.UserRepository.GetProfile(userId)

	fmt.Println(user)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: user}
	json.NewEncoder(w).Encode(response)
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var request usersdto.UpdateUserRequest

	dataUpload := r.Context().Value("dataFile")
	var filename string
	if dataUpload != "" {
		filename = dataUpload.(string)
		request = usersdto.UpdateUserRequest{
			FullName: r.FormValue("fullname"),
			Image:    os.Getenv("UPLOAD_path_name") + filename,
			Email:    r.FormValue("email"),
			Phone:    r.FormValue("phone"),
			Location: r.FormValue("location"),
		}
	} else {
		request = usersdto.UpdateUserRequest{
			FullName: r.FormValue("fullname"),
			Email:    r.FormValue("email"),
			Phone:    r.FormValue("phone"),
			Location: r.FormValue("location"),
		}
	}

	userModel := models.User{}
	// hashedPassword, _ := bcryptpkg.HashingPassword(request.Password)

	if request.FullName != "" {
		userModel.FullName = request.FullName
	}

	if request.Image != "" {
		userModel.Image = request.Image
	}

	if request.Email != "" {
		userModel.Email = request.Email
	}
	if request.Phone != "" {
		userModel.Phone = request.Phone
	}
	if request.Location != "" {
		userModel.Location = request.Location
	}

	// if request.Password != "" {
	// 	userModel.Password = hashedPassword
	// }

	user, err := h.UserRepository.UpdateUser(userModel, id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: user}

	json.NewEncoder(w).Encode(response)

}

// delete
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user := models.User{}

	deletedUser, err := h.UserRepository.DeleteUser(user, id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: deletedUser}
	json.NewEncoder(w).Encode(response)
}

// convert many users response
func convertUsersResponse(u []models.User) []models.UsersProfileResponse {

	var resp []models.UsersProfileResponse

	for _, item := range u {
		resp = append(resp, models.UsersProfileResponse{
			ID:       item.ID,
			FullName: item.FullName,
			Image:    item.Image,
			Email:    item.Email,
			Phone:    item.Phone,
			Location: item.Location,
		})
	}

	return resp
}
