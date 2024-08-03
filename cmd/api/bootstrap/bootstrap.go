package bootstrap

import (
	"context"
	"errors"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/javiertelioz/template-clean-architecture-go/internal/application/use_cases"
	"github.com/javiertelioz/template-clean-architecture-go/internal/infrastructure/proto/user"
	"github.com/javiertelioz/template-clean-architecture-go/internal/infrastructure/repositories"
	"github.com/javiertelioz/template-clean-architecture-go/internal/infrastructure/services"
	"github.com/javiertelioz/template-clean-architecture-go/internal/presentation/controllers"
	"github.com/javiertelioz/template-clean-architecture-go/internal/presentation/routes"
)

func Run() {
	// Crear el servidor HTTP
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: service(),
	}

	// Crear el servidor gRPC
	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, &services.UserServiceServer{})
	reflection.Register(grpcServer)

	// Canal para recibir señales del sistema operativo
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// Contexto para controlar el cierre de los servidores
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Función para cerrar ambos servidores
	go func() {
		<-sig

		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		err := httpServer.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}

		grpcServer.GracefulStop()
		serverStopCtx()
	}()

	// Iniciar el servidor HTTP
	go func() {
		log.Printf("Starting HTTP server on: http://localhost%s\n", httpServer.Addr)
		err := httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	// Iniciar el servidor gRPC
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Printf("Starting gRPC server on: %s\n", lis.Addr())
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Iniciar el servidor gRPC-Gateway
	go func() {
		err := runGRPCGatewayServer()
		if err != nil {
			log.Fatal(err)
		}
	}()

	<-serverCtx.Done()
}

func service() http.Handler {
	userRepository := repositories.NewInMemoryUserRepository()

	createUserUseCase := use_cases.NewCreateUserUseCase(userRepository)
	getUserByIDUseCase := use_cases.NewGetUserByIDUseCase(userRepository)

	userController := controllers.NewUserController(*createUserUseCase, *getUserByIDUseCase)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/ping"))

	r.Mount("/api/v1/users", routes.UserRoutes(userController))

	return r
}

func runGRPCGatewayServer() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := user.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		return err
	}

	log.Println("Starting gRPC-Gateway server on :8081")
	return http.ListenAndServe(":8081", mux)
}
