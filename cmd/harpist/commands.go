package main

import (
	// "bufio"
	"flag"
	"fmt"
	// "log"
	"errors"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
	"syscall"
	//
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	//
	// "gopkg.in/harpist.v0"
	"gopkg.in/harpist.v0/models"
)

var (
	initCmd = flag.NewFlagSet("init", flag.ExitOnError)

	userCmd       = flag.NewFlagSet("user", flag.ExitOnError)
	userAddCmd    = flag.NewFlagSet("add", flag.ExitOnError)
	userRemoveCmd = flag.NewFlagSet("remove", flag.ExitOnError)
	userInfoCmd   = flag.NewFlagSet("info", flag.ExitOnError)
	userSetPWCmd  = flag.NewFlagSet("set-password", flag.ExitOnError)

	groupCmd = flag.NewFlagSet("group", flag.ExitOnError)

	gameCmd = flag.NewFlagSet("game", flag.ExitOnError)

	charCmd = flag.NewFlagSet("char", flag.ExitOnError)
)

func parseInitCmd() error {
	err := initCmd.Parse(os.Args[2:])
	if initCmd.Parsed() {
		if len(initCmd.Args()) == 0 {
			initDatabase()
			if config.Environment == "dev" {
				// In examples.go
				populateExamples()
			}
			return nil
		}
	}
	return err
}

func initDatabase() {
	logger.Printf("Initializing database...")
	var adminUser = models.User{ID: 1}
	db.FirstOrInit(&adminUser, &adminUser)
	if db.NewRecord(adminUser) {
		adminUser.Name = "Admin User"
		adminUser.LoginInfo.Username = "admin"
		adminUser.SetPassword("password")
		adminUser.EmailAddress = "admin@example.com"
		db.Save(&adminUser)
	} else {
		fmt.Printf("Database already initialized for %v environment.", config.Environment)
	}
}

func parseUserCmd() error {
	err := userCmd.Parse(os.Args[2:])

	if userCmd.Parsed() {
		switch userCmd.Arg(0) {
		case "add":
			err = parseUserAddCmd()
		case "remove":
			err = parseUserRemCmd()
		case "info":
			err = parseUserInfoCmd()
		case "set-password":
			err = parseUserSetPWCmd()
		}
	}
	return err
}

func parseUserAddCmd() error {
	var (
		username string
		password *string
		name     string
		email    string
		err      error
	)

	password = userAddCmd.String("p", "", "password")

	passedArgs := userCmd.Args()

	err = userAddCmd.Parse(passedArgs[1:])

	if err != nil {
		logger.Fatalf("ERROR: %v", err)
	}

	if userAddCmd.Parsed() {
		username = userAddCmd.Arg(0)
		if username == "" {
			err = fmt.Errorf("must specify at least a username and email address")
			logger.Printf("ERROR: No username specified")
			return err
		}
		email = userAddCmd.Arg(1)
		if email == "" {
			err = fmt.Errorf("must specify at least a username and email address")
			logger.Printf("ERROR: No email specified")
			return err
		}
		name = strings.Join(userAddCmd.Args()[2:], " ")

		user := models.User{LoginInfo: models.LoginInfo{Username: username}}
		db.Where(user).First(&user)

		if db.NewRecord(user) {
			logger.Printf("Adding user %v\n", username)
			user.Name = name
			user.EmailAddress = email
			user.SetPassword(*password)
			db.Create(&user)
			err = db.Save(&user).Error
		}
	}
	return err
}

func parseUserRemCmd() error {
	var (
		username *string
		id       *uint
		email    *string
		err      error

		user models.User
	)

	username = userRemoveCmd.String("u", "", "username (shorthand)")
	email = userRemoveCmd.String("e", "", "email (shorthand)")
	id = userRemoveCmd.Uint("i", 0, "user id (shorthand)")

	passedArgs := userCmd.Args()

	err = userRemoveCmd.Parse(passedArgs[1:])

	if err != nil {
		logger.Printf("ERROR: %v", err)
	}

	if userRemoveCmd.Parsed() {
		if *id != 0 {
			user = models.User{
				ID: *id,
			}
			err = db.Delete(models.User{}, &user).Error
			return err
		}
		if *username != "" {
			user = models.User{
				LoginInfo: models.LoginInfo{
					Username: *username,
				},
			}
			err = db.Delete(models.User{}, &user).Error
			return err
		}
		if *email != "" {
			user = models.User{
				EmailAddress: *email,
			}
			err = db.Delete(models.User{}, &user).Error
			return err
		}
	}
	return err
}

func parseUserInfoCmd() error {
	var (
		username string
		id       *uint
		email    *string
		err      error

		user = models.User{}
	)

	email = userInfoCmd.String("e", "", "email (shorthand)")
	id = userInfoCmd.Uint("i", 0, "user id (shorthand)")

	passedArgs := userCmd.Args()

	err = userInfoCmd.Parse(passedArgs[1:])
	if err != nil {
		logger.Printf("ERROR: %v", err)
	}

	if userInfoCmd.Parsed() {
		username = userInfoCmd.Arg(0)

		if username == "" {
			if *id != 0 {
				user.ID = *id
			}
			if *email != "" {
				user.EmailAddress = *email
			}
		}
		user.LoginInfo.Username = username

		err = db.Find(&user, user).Error
		if err != nil {
			return err
		}
		fmt.Println("User Details:")
		fmt.Printf("\t- ID: %v\n", user.ID)
		fmt.Printf("\t- Username: %v\n", user.Username())
		fmt.Printf("\t- Name: %v\n", user.MemberName())
		fmt.Printf("\t- Email: %v\n", user.EmailAddress)
		return nil
	}
	return err
}

func parseUserSetPWCmd() error {
	var (
		password     string
		err          error
		bytePassword []byte
		byteConfirm  []byte
		confirm      string
		user         = models.User{}
	)

	passedArgs := userCmd.Args()

	err = userSetPWCmd.Parse(passedArgs[1:])
	if err != nil {
		logger.Printf("ERROR: %v", err)
	}
	if userSetPWCmd.Parsed() {
		user.LoginInfo.Username = userSetPWCmd.Arg(0)
		password = userSetPWCmd.Arg(1)
		err = db.Find(&user, user).Error
		if err != nil {
			return err
		}
		if password == "" {
			fmt.Print("Password: ")
			bytePassword, err = terminal.ReadPassword(int(syscall.Stderr))
			fmt.Println("")
			if err != nil {
				return err
			}
			password = string(bytePassword)
			fmt.Print("Re-enter Password: ")
			byteConfirm, err = terminal.ReadPassword(int(syscall.Stderr))
			fmt.Println("")
			if err != nil {
				return err
			}
			confirm = string(byteConfirm)
			if confirm != password {
				return errors.New("passwords do not match")
			}
		}
		user.SetPassword(password)
		err = db.Save(&user).Error
		return err
	}
	return errors.New("record not found")
}

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
	if len(os.Args) < 2 {
		flag.PrintDefaults()
		os.Exit(0)
	}
	switch os.Args[1] {
	case "init":
		return parseInitCmd
	case "user":
		return parseUserCmd
	case "group":
		return parseGroupCmd
	case "game":
		return parseGameCmd
	case "char":
		return parseCharCmd
	case "help":
		flag.PrintDefaults()
		os.Exit(0)
	default:
		flag.PrintDefaults()
		os.Exit(0)
	}
	return nil
}
