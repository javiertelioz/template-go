package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/javiertelioz/template-clean-architecture-go/internal/presentation/controllers"
	"net/http"
)

func UserRoutes(controller *controllers.UserController) http.Handler {
	r := chi.NewRouter()

	r.Post("/", controller.CreateUser)
	r.Get("/", controller.GetUsers)
	r.Get("/{id}", controller.GetUserByID)
	r.Put("/{id}", controller.UpdateUserByID)
	r.Delete("/{id}", controller.DeleteUser)

	return r
}
