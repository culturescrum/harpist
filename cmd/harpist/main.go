package main

import (
	"fmt"
	"os"

	"gopkg.in/harpist.v0"
	// "github.com/go-harpist/harpist"
	"gopkg.in/harpist.v0/models"
)

// shorthand
var (
	config = harpist.HarpistConfig
	db     = harpist.HarpistDB
	logger = harpist.HarpistLogger
)

func init() {

	if config.Environment != "prod" {
		logger.Printf("Initalizing for environment: %v", config.Environment)
	}

	dbp, err := harpist.GetDb()
	if err != nil {
		logger.Fatalf("error: %v", err)
		os.Exit(1)
	}

	db = dbp
	db.SetLogger(logger)
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.PlayGroup{})
	db.AutoMigrate(&models.Game{})
	db.AutoMigrate(&models.Character{})
}

func main() {
	err := ParseArgs()()

	if err != nil {
		fmt.Println(err)
	}
}
