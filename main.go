package harpist

import (
	"fmt"
	"io"
	"log"
	"os"
	//  "runtime"

	"github.com/jinzhu/gorm"
	// needed for gorm init
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	// HarpistConfig sets the platform-level Config
	HarpistConfig = GetConfig()
	// HarpistDB platform database
	HarpistDB, _  = GetDb()
	HarpistLogger = GetLogger()
)

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

// HarpistLogger defines a wrapper around l.Logger
// TODO: make this an exported var that calls GetLogger()
func oldHarpistLogger(w io.Writer, env string) *log.Logger {
	var logger = log.New(w, "HARPIST: ", log.Ldate|log.Ltime)
	return logger
}

// TODO: implement GetLogger(); should pull from environment default values
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
	logger := log.New(logfile, "HARPIST: ", log.Ldate|log.Ltime)

	return logger
}

// GetDb only returns development sqlite at the moment
func GetDb() (*gorm.DB, error) {
	dbFilename := fmt.Sprintf("%v.db", HarpistConfig.Environment)
	db, err := gorm.Open("sqlite3", dbFilename)

	return db, err
}
