package platform

import (
  "os"
//  "fmt"
//  "runtime"
  )

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
  config := Config{Environment: env,}

  return config

}
