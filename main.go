package main

import (
	"github.com/theisaachome/eWallet-platform/cmd/ewallet"
	"github.com/theisaachome/eWallet-platform/pkg/utils/logger"
)

func main() {
	logger.Info("Starting EWallet Platform")
	ewallet.StartApplication()
}
