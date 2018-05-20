package main

import (
	"fmt"

	"gopkg.in/harpist.v0"
	// "github.com/go-harpist/harpist"
)

// shorthand
var (
	config = harpist.HarpistConfig
	db     = harpist.HarpistDB
	logger = harpist.HarpistLogger
)

func main() {
	err := ParseArgs()()

	if err != nil {
		fmt.Println(err)
	}
}
