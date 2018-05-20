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

// HarpistConfig sets the platform-level Config
// TODO: add functionality for application-level Config
var HarpistConfig = GetConfig()

// Config defines the structure for configuration for all cmds
type Config struct {
	Environment string
}

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
func HarpistLogger(w io.Writer, env string) *log.Logger {
	var logger = log.New(w, "HARPIST: ", log.Ldate|log.Ltime)
	return logger
}

// TODO: implement GetLogger(); should pull from environment default values

// HarpistDB platform database
var HarpistDB, _ = GetDb()

// GetDb only returns development sqlite at the moment
func GetDb() (*gorm.DB, error) {
	dbFilename := fmt.Sprintf("%v.db", HarpistConfig.Environment)
	db, err := gorm.Open("sqlite3", dbFilename)

	return db, err
}
