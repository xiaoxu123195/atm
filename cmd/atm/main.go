package main

import (
	"fmt"
	"os"

	"github.com/xiaoxu123195/atm/pkg/app"
	"github.com/xiaoxu123195/atm/pkg/i18n"
)

const (
	VERSION        = "1.0.0"
	REPOSITORY_URL = "https://github.com/xiaoxu123195/atm"
)

func main() {
	// Initialize i18n system
	i18n.Init()

	// Create and run application
	application := app.NewApp(VERSION, REPOSITORY_URL)

	if err := application.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
