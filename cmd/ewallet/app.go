package ewallet

import (
	"fmt"
	http2 "github.com/theisaachome/eWallet-platform/internal/interfaces/http"
	"log"
	"net/http"
	"os"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variable SERVER_ADDRESS and SERVER_PORT not set")
	}
}

func StartApplication() {
	sanityCheck()
	// router instance
	routers := http2.GetWalletRouter()

	// get server instance and start
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), routers))
}
