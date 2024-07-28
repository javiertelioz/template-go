package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/javiertelioz/template-clean-architecture-go/internal/application/dto/user"
	"github.com/javiertelioz/template-clean-architecture-go/internal/application/use_cases"
)

type UserController struct {
	createUserUseCase  use_cases.CreateUserUseCase
	getUserByIDUseCase use_cases.GetUserByIDUseCase
}

func NewUserController(
	createUserUseCase use_cases.CreateUserUseCase,
	getUserByIDUseCase use_cases.GetUserByIDUseCase,
) *UserController {
	return &UserController{
		createUserUseCase:  createUserUseCase,
		getUserByIDUseCase: getUserByIDUseCase,
	}
}

func (ctl *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userDto user.CreateUserDTO

	json.NewDecoder(r.Body).Decode(&userDto)

	err := ctl.createUserUseCase.Execute(userDto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (ctl *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")

	userDto, err := ctl.getUserByIDUseCase.Execute(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userDto)
}
