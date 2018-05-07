package platform

import (
	"github.com/jinzhu/gorm"
	// needed for gorm init
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// GetDb only returns development sqlite at the moment
func GetDb() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "development.db")

	return db, err
}
