package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	
	var appUser *user.User
	
	appUser, err := user.New(userFirstName, userLastName, userBirthdate)
	
	if err != nil {
		fmt.Println(err)
		return
	}
	
	admin := user.NewAdmin("test@example.com", "test123")
	
	admin.OutputUserData()
	admin.ClearUserName()

	// ... do something awesome with that gathered data!

	appUser.OutputUserData()
	appUser.ClearUserName()
	appUser.OutputUserData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

