package handlers

import (
	"encoding/json"
	"fmt"
	authdto "go-batch2/dto/auth"
	dto "go-batch2/dto/result"
	usersdto "go-batch2/dto/users"
	"go-batch2/models"
	bcryptpkg "go-batch2/pkg/bcrypt"
	jwttoken "go-batch2/pkg/jwt"
	"go-batch2/repositories"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) Register(w http.ResponseWriter, r *http.Request) {
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

	// check if user email already exist
	userExist, emailErr := h.AuthRepository.Login(request.Email)

	if emailErr == nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: "Email " + userExist.Email + " already exist!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	hashedPassword, _ := bcryptpkg.HashingPassword(request.Password)
	userModel := models.User{
		Email:    request.Email,
		Password: hashedPassword,
		FullName: request.FullName,
		Gender:   request.Gender,
		Phone:    request.Phone,
		Role:     request.Role,
		Image:    os.Getenv("UPLOAD_PATH_NAME") + "default_profile.png", // set default profile photo
	}

	user, err := h.AuthRepository.Register(userModel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: user}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerAuth) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("COntent-Type", "application/json")

	request := new(authdto.AuthRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	userModel := models.User{
		Email:    request.Email,
		Password: request.Password,
	}

	user, err := h.AuthRepository.Login(userModel.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: "Email not registered!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	if isPasswordMatch := bcryptpkg.CheckPasswordHash(userModel.Password, user.Password); !isPasswordMatch {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: "Wrong password!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	generateToken := jwt.MapClaims{}

	generateToken["id"] = user.ID
	generateToken["exp"] = time.Now().Add(time.Hour * 3).Unix()

	token, err := jwttoken.CreateToken(&generateToken)
	if err != nil {
		log.Println(err)
		fmt.Println("Unauthorize")
		return
	}

	AuthResponse := authdto.AuthResponse{
		ID: user.ID,
		Email:    user.Email,
		Role:     user.Role,
		Token:    token,
	}

	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Status: "Success", Data: AuthResponse}
	json.NewEncoder(w).Encode(response)

}
