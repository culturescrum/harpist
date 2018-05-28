package harpist

import (
	"database/sql"
	"log"
	"os"
	"testing"
)

// TestInit tests database initialization
func TestInit(t *testing.T) {
	var config interface{} = HarpistConfig
	var db interface{} = HarpistDB.DB()
	var logger interface{} = HarpistLogger

	cv, ok := config.(Config)
	if !ok {
		t.Error(
			"For", config,
			"expected Config, got", cv,
		)
	}

	dv, ok := db.(*sql.DB)
	if !ok {
		t.Error(
			"For", db,
			"expected sql.DB, got", dv,
		)

	}

	lv, ok := logger.(*log.Logger)

	if !ok {
		t.Error(
			"For", logger,
			"expected log.Logger, got", lv,
		)
	}

	_, lcheck := os.Stat("harpist.dev.log")
	if os.IsNotExist(lcheck) {
		t.Error("Logger did not initialize new log file")
	}

	HarpistConfig = Config{Environment: "test"}
	HarpistLogger = GetLogger()
	HarpistDB, _ = GetDb()
	_, ltcheck := os.Stat("harpist.test.log")
	if os.IsNotExist(ltcheck) {
		t.Error("New Logger instance did not pick up environment change")
	}
}
