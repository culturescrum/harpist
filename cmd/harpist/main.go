package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/culturescrum/harpist/internal/platform"
	"github.com/culturescrum/harpist/models"
)

var config = platform.GetConfig()

var db *gorm.DB
var logger *log.Logger

func init() {
	var lfn = "harpist.log"
	var _, lcheck = os.Stat(lfn)

	// create file if not exists
	if os.IsNotExist(lcheck) {
		var file, err = os.Create(lfn)
		if err != nil {
			os.Exit(1)
		}
		file.Close()
	}

	var logfile, err = os.OpenFile(lfn, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error creating log file, something is wrong: %v", err)
		os.Exit(2)
	}

	logger = platform.HarpistLogger(logfile, config.Environment)
	logger.Printf("Logger initialized.")

	logger.Printf("Initalizing for environment: %v", config.Environment)

	logger.Printf("Initializing database.")
	dbp, err := platform.GetDb()
	if err != nil {
		logger.Fatalf("error: %v", err)
		os.Exit(1)
	}
	db = dbp
	db.SetLogger(logger)
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.PlayGroup{})
	db.AutoMigrate(&models.Game{})
	db.AutoMigrate(&models.Character{})
	initDatabase()
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

func populateExamples() {
	loopNums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, i := range loopNums {
		logger.Printf("Creating user: User %v", i)
		var user = models.User{
			Name:         fmt.Sprintf("User %v", i),
			EmailAddress: fmt.Sprintf("user%v@example.com", i),
		}
		user.LoginInfo.Username = fmt.Sprintf("username%v", i)
		user.SetPassword("password")
		db.FirstOrCreate(&user, models.User{LoginInfo: models.LoginInfo{Username: user.LoginInfo.Username}})
		exampleGroup := models.PlayGroup{}
		db.First(&exampleGroup)
		logger.Printf("Adding `%v` to group `%v`", user.Username(), exampleGroup.MemberName())
		exampleGroup.AddMember(user)
		db.Save(&exampleGroup)
	}

	storyteller := models.User{ID: 3}
	db.First(&storyteller)
	fmt.Printf("Storyteller: %v\n", storyteller.Name)

	gameAdmin := models.User{ID: 2}
	db.First(&gameAdmin)
	fmt.Printf("Game Admin: %v\n", gameAdmin.Name)

	player := models.User{ID: 4}
	db.First(&player)
	fmt.Printf("Player: %v\n", player.Name)

	game := models.Game{Name: "Example Game", Owner: storyteller}
	game.AddAdmin(gameAdmin)
	game.AddMember(player)

	character := models.Character{Name: "Example Character", Owner: player}
	game.AddCharacter(character)

	db.Save(&game)

}

func main() {
	populateExamples()
}
