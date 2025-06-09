package http

import (
	"github.com/gorilla/mux"
	appUser "github.com/theisaachome/eWallet-platform/internal/app/user"
	domainUser "github.com/theisaachome/eWallet-platform/internal/domain/user"
	"github.com/theisaachome/eWallet-platform/internal/infrastructure/db/postgres"
	"github.com/theisaachome/eWallet-platform/internal/interfaces/http/handlers"
)

func GetWalletRouter() *mux.Router {

	router := mux.NewRouter()

	// Wire service and handler
	db := postgres.NewPostgresDB()
	userRepo := domainUser.NewRepositoryDb(db)
	userService := appUser.NewService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Setup routes
	api := router.PathPrefix("/ewallet/api").Subrouter()
	api.HandleFunc("/users", userHandler.NewUser).Methods("POST")

	return router
}
