package main

import (
	// "flag"
	"fmt"

	"github.com/culturescrum/harpist/models"
)

func populateExamples() {
	var exampleGroup = models.PlayGroup{Name: "Example Group"}
	db.FirstOrCreate(&exampleGroup, models.PlayGroup{ID: 1})
	adminUser := models.User{}
	db.Where(models.User{ID: 1}).First(&adminUser)
	exampleGroup.Owner = adminUser
	exampleGroup.AddAdmin(adminUser)
	exampleGroup.AddMember(adminUser)
	db.Save(&exampleGroup)

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
	logger.Printf("Storyteller: %v\n", storyteller.Name)

	gameAdmin := models.User{ID: 2}
	db.First(&gameAdmin)
	logger.Printf("Game Admin: %v\n", gameAdmin.Name)

	player := models.User{ID: 4}
	db.First(&player)
	logger.Printf("Player: %v\n", player.Name)

	game := models.Game{Name: "Example Game", Owner: storyteller}
	game.AddAdmin(gameAdmin)
	game.AddMember(player)

	character := models.Character{Name: "Example Character", Owner: player}
	game.AddCharacter(character)

	db.Save(&game)

}
