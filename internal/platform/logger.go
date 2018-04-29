package platform

import (
  "log"
  "bytes"
)

// HarpistLogger defines a wrapper around l.Logger
func HarpistLogger(buf *bytes.Buffer, env string) *log.Logger {
  var logger = log.New(buf, "HARPIST: ", log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
  logger.Printf("Initializing for %v\n", env)
  return logger
}
