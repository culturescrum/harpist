package platform

import (
	"io"
	"log"
)

// HarpistLogger defines a wrapper around l.Logger
func HarpistLogger(w io.Writer, env string) *log.Logger {
	var logger = log.New(w, "HARPIST: ", log.Ldate|log.Ltime)
	return logger
}
