package bootstrap

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/javiertelioz/template-clean-architecture-go/internal/application/use_cases"
	"github.com/javiertelioz/template-clean-architecture-go/internal/infrastructure/repositories"
	"github.com/javiertelioz/template-clean-architecture-go/internal/presentation/controllers"
	"github.com/javiertelioz/template-clean-architecture-go/internal/presentation/routes"
)

func Run() {
	server := http.Server{
		Addr:    ":8080",
		Handler: ServerService(),
	}

	log.Printf("Starting server on: http://localhost:%s\n", server.Addr)

	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}

	<-serverCtx.Done()
}

func ServerService() http.Handler {

	userRepository := repositories.NewInMemoryUserRepository()

	createUserUseCase := use_cases.NewCreateUserUseCase(userRepository)
	getUserByUseCase := use_cases.NewGetUsesUseCase(userRepository)
	getUserByIDUseCase := use_cases.NewGetUserByIDUseCase(userRepository)
	updateUserByIDUseCase := use_cases.NewUpdateUserByIDUseCase(userRepository)
	deleteUserByIDUseCase := use_cases.NewDeleteUserByIDUseCase(userRepository)

	userController := controllers.NewUserController(
		*createUserUseCase,
		*getUserByUseCase,
		*getUserByIDUseCase,
		*updateUserByIDUseCase,
		*deleteUserByIDUseCase,
	)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/ping"))

	r.Mount("/api/v1/users", routes.UserRoutes(userController))

	return r
}
