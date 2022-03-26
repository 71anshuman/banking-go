package main

import (
	"github.com/71anshuman/banking-go/app"
	"github.com/71anshuman/banking-go/logger"
)

func main() {
	logger.Info("Starting our application")
	app.Start()
}
