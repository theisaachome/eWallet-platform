package http

import (
	"github.com/gorilla/mux"
	"github.com/theisaachome/eWallet-platform/internal/app/auth"
	appUser "github.com/theisaachome/eWallet-platform/internal/app/user"
	appWallet "github.com/theisaachome/eWallet-platform/internal/app/wallet"
	domainUser "github.com/theisaachome/eWallet-platform/internal/domain/user"
	"github.com/theisaachome/eWallet-platform/internal/domain/wallet"
	"github.com/theisaachome/eWallet-platform/internal/infrastructure/db/postgres"
	"github.com/theisaachome/eWallet-platform/internal/interfaces/http/handlers"
	"github.com/theisaachome/eWallet-platform/pkg/security/jwt"
	"os"
)

func GetWalletRouter() *mux.Router {

	router := mux.NewRouter()

	// Wire service and handler
	db := postgres.NewPostgresDB()
	// repos wiring
	userRepo := domainUser.NewRepositoryDb(db)
	walletRepo := wallet.NewRepositoryDb(db)

	// Wire service and handler
	walletService := appWallet.NewService(walletRepo)
	userService := appUser.NewService(userRepo, walletService)
	// jwt service
	jwtService := jwt.NewJwtService(os.Getenv("JWT_SECRET"))
	userHandler := handlers.NewUserHandler(userService)

	authService := auth.NewAuthService(userRepo, walletService, jwtService)
	authHandler := handlers.NewAuthHandler(authService)

	// Setup routes
	api := router.PathPrefix("/wallet/api").Subrouter()
	api.HandleFunc("/users", userHandler.NewUser).Methods("POST")
	// Auth routers
	api.HandleFunc("/auth/register", authHandler.Register).Methods("POST")
	api.HandleFunc("/auth/login", authHandler.Login).Methods("POST")

	return router
}
