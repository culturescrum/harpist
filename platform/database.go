package platform

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// needed for gorm init
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// HarpistDB platform database
var HarpistDB, _ = GetDb()

// GetDb only returns development sqlite at the moment
func GetDb() (*gorm.DB, error) {
	dbFilename := fmt.Sprintf("%v.db", HarpistConfig.Environment)
	db, err := gorm.Open("sqlite3", dbFilename)

	return db, err
}
