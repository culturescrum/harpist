package main

import (
	// "bufio"
	"flag"
	"fmt"
	// "log"
	"os"
	//
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	//
	// "github.com/culturescrum/harpist/internal/platform"
	// "github.com/culturescrum/harpist/models"
)

var userCmd = flag.NewFlagSet("user", flag.ExitOnError)

func parseUserCmd() error {
	err := userCmd.Parse(os.Args[2:])
	if userCmd.Parsed() {
		if userCmd.Arg(0) == "add" {
			fmt.Println("Not implemented!")
			os.Exit(0)
		}
	}
	return err
}

var groupCmd = flag.NewFlagSet("group", flag.ExitOnError)

func parseGroupCmd() error {
	err := groupCmd.Parse(os.Args[2:])
	if groupCmd.Parsed() {
		if groupCmd.Arg(0) == "add" {
			fmt.Println("Not implemented!")
			os.Exit(0)
		}
	}
	return err
}

var gameCmd = flag.NewFlagSet("game", flag.ExitOnError)

func parseGameCmd() error {
	err := gameCmd.Parse(os.Args[2:])
	if gameCmd.Parsed() {
		if gameCmd.Arg(0) == "add" {
			fmt.Println("Not implemented!")
			os.Exit(0)
		}
	}
	return err
}

var charCmd = flag.NewFlagSet("char", flag.ExitOnError)

func parseCharCmd() error {
	err := charCmd.Parse(os.Args[2:])
	if charCmd.Parsed() {
		if charCmd.Arg(0) == "add" {
			fmt.Println("Not implemented!")
			os.Exit(0)
		}
	}
	return err
}

func ParseArgs() func() error {
	switch os.Args[1] {
	case "user":
		return parseUserCmd
	case "group":
		return parseGroupCmd
	case "game":
		return parseGameCmd
	case "char":
		return parseCharCmd
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
	return nil
}
