package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/javiertelioz/template-clean-architecture-go/internal/application/dto/user"
	"github.com/javiertelioz/template-clean-architecture-go/internal/application/use_cases"
)

type UserController struct {
	createUserUseCase     use_cases.CreateUserUseCase
	getUsersUseCase       use_cases.GetUsesUseCase
	getUserByIDUseCase    use_cases.GetUserByIDUseCase
	updateUserByIDUseCase use_cases.UpdateUserByIDUseCase
	deleteUserByIDUseCase use_cases.DeleteUserByIDUseCase
}

func NewUserController(
	createUserUseCase use_cases.CreateUserUseCase,
	getUsesUseCase use_cases.GetUsesUseCase,
	getUserByIDUseCase use_cases.GetUserByIDUseCase,
	updateUserByIDUseCase use_cases.UpdateUserByIDUseCase,
	deleteUserByIDUseCase use_cases.DeleteUserByIDUseCase,
) *UserController {
	return &UserController{
		createUserUseCase:     createUserUseCase,
		getUsersUseCase:       getUsesUseCase,
		getUserByIDUseCase:    getUserByIDUseCase,
		updateUserByIDUseCase: updateUserByIDUseCase,
		deleteUserByIDUseCase: deleteUserByIDUseCase,
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
	userID := chi.URLParam(r, "id")

	userDto, err := ctl.getUserByIDUseCase.Execute(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(userDto)
}

func (ctl *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := ctl.getUsersUseCase.Execute()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(users)
}

func (ctl *UserController) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")

	var userDto user.UpdateUserDTO

	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	userDto.ID = userID

	err = ctl.updateUserByIDUseCase.Execute(userDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (ctl *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")

	err := ctl.deleteUserByIDUseCase.Execute(userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
