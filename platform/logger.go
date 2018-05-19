package platform

import (
	"io"
	"log"
)

// HarpistLogger defines a wrapper around l.Logger
// TODO: make this an exported var that calls GetLogger()
func HarpistLogger(w io.Writer, env string) *log.Logger {
	var logger = log.New(w, "HARPIST: ", log.Ldate|log.Ltime)
	return logger
}

// TODO: implement GetLogger(); should pull from environment default values
