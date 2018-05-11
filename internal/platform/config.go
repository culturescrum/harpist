package platform

import (
	"os"
	//  "fmt"
	//  "runtime"
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
