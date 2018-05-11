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
	"github.com/culturescrum/harpist/models"
)

var (
	initCmd = flag.NewFlagSet("init", flag.ExitOnError)

	userCmd       = flag.NewFlagSet("user", flag.ExitOnError)
	userAddCmd    = flag.NewFlagSet("add", flag.ExitOnError)
	userRemoveCmd = flag.NewFlagSet("remove", flag.ExitOnError)

	// userInfoCmd = flag.NewFlagSet("info", flag.ExitOnError)

	groupCmd = flag.NewFlagSet("group", flag.ExitOnError)

	gameCmd = flag.NewFlagSet("game", flag.ExitOnError)

	charCmd = flag.NewFlagSet("char", flag.ExitOnError)
)

func parseInitCmd() error {
	err := initCmd.Parse(os.Args[2:])
	if initCmd.Parsed() {
		if len(initCmd.Args()) == 0 {
			initDatabase()
			populateExamples()
			return nil
		}
	}
	return err
}

func initDatabase() {
	var adminUser = models.User{ID: 1}
	db.FirstOrInit(&adminUser, &adminUser)
	adminUser.Name = "Admin User"
	adminUser.LoginInfo.Username = "admin"
	adminUser.SetPassword("password")
	adminUser.EmailAddress = "admin@example.com"
	db.Save(&adminUser)
	var exampleGroup = models.PlayGroup{Name: "Example Group"}
	db.FirstOrCreate(&exampleGroup, models.PlayGroup{ID: 1})
	exampleGroup.Owner = adminUser
	exampleGroup.AddAdmin(adminUser)
	exampleGroup.AddMember(adminUser)
	db.Save(&exampleGroup)
}

func parseUserCmd() error {
	err := userCmd.Parse(os.Args[2:])

	if userCmd.Parsed() {
		switch userCmd.Arg(0) {
		case "add":
			parseUserAddCmd()
		case "remove":
			parseUserRemCmd()
		}
	}
	return err
}

func parseUserAddCmd() error {
	var (
		username *string
		password *string
		name     *string
		email    *string
		err      error
	)

	username = userAddCmd.String("u", "", "username (shorthand)")
	password = userAddCmd.String("p", "", "password (shorthand)")
	name = userAddCmd.String("n", "", "name (shorthand)")
	email = userAddCmd.String("e", "", "email (shorthand)")

	passedArgs := userCmd.Args()

	err = userAddCmd.Parse(passedArgs[1:])

	if err != nil {
		logger.Fatalf("ERROR: %v", err)
	}

	if userAddCmd.Parsed() {
		if *username == "" {
			err = fmt.Errorf("ERROR: Must specify at least a username.")
			logger.Printf("ERROR: No username specified")
			logger.Printf("DEBUG: Args parsed: ")
			for _, arg := range userAddCmd.Args() {
				logger.Printf("%v, ", arg)
			}
			logger.Printf("\n")
			logger.Printf("DEBUG: Flags parsed: ")
			userAddCmd.VisitAll(func(f *flag.Flag) {
				logger.Printf("%v: %v", f.Name, f.Value)
			})
			return err
		}
		logger.Printf("Adding user %v\n", *username)
		user := models.User{
			LoginInfo:    models.LoginInfo{Username: *username},
			Name:         *name,
			EmailAddress: *email,
		}
		user.SetPassword(*password)
		err = db.Where(models.User{LoginInfo: models.LoginInfo{Username: *username}}).FirstOrCreate(&user).Error
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
		logger.Fatalf("ERROR: %v", err)
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
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
	return nil
}
