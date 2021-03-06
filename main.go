package harpist

import (
	"fmt"
	"log"
	"os"
	//  "runtime"

	"github.com/jinzhu/gorm"
	"gopkg.in/harpist.v0/models"
	// needed for gorm init
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	// HarpistConfig sets the platform-level Config
	HarpistConfig = GetConfig()
	// HarpistDB platform database
	HarpistDB, _ = GetDb()
	// HarpistLogger defines a central logger with environment pivots
	HarpistLogger = GetLogger()
)

func init() {
	var db = HarpistDB
	db.SetLogger(HarpistLogger)
	for _, model := range models.CoreModels {
		db.AutoMigrate(model)
	}
}

// TODO: add functionality for application-level Config

// GetConfig returns the default configuration structure
func GetConfig() Config {
	env := os.Getenv("HARPIST_ENV")
	if len(env) == 0 {
		env = "dev"
	}
	config := Config{Environment: env}

	return config

}

// GetLogger returns a standard logger instance for any project to use.
func GetLogger() *log.Logger {

	lfn := fmt.Sprintf("harpist.%v.log", HarpistConfig.Environment)
	var _, lcheck = os.Stat(lfn)

	// create file if not exists
	if os.IsNotExist(lcheck) {
		var file, err = os.Create(lfn)
		if err != nil {
			os.Exit(1)
		}
		file.Close()
	}

	var logfile, err = os.OpenFile(lfn, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error creating log file, something is wrong: %v", err)
		os.Exit(2)
	}
	var logger = log.New(logfile, "HARPIST: ", log.Ldate|log.Ltime)

	return logger
}

// GetDb only returns development sqlite at the moment
func GetDb() (*gorm.DB, error) {
	dbFilename := fmt.Sprintf("%v.db", HarpistConfig.Environment)
	db, err := gorm.Open("sqlite3", dbFilename)

	return db, err
}
