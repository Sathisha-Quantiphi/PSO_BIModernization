package main

import (
	"fmt"
	"time"
)

type Tag struct {
	Name, Type string
}

type Settings struct {
	NotificationsEnabled bool
}

type User struct {
	Id, Name, LastName, Status string
	Tags                       []*Tag
	*Settings
}

type NotificationsService struct {
}

func main() {
	// Creating channel with type as slice of pointer to user
	usersToUpdate := make(chan []*User)
	userToNotify := make(chan *User)
	newUsers := []*User{
		{Name: "Vishal", Status: "active", Settings: &Settings{NotificationsEnabled: true}},
		{Name: "Jayam", Status: "active", Settings: &Settings{NotificationsEnabled: false}},
		{Name: "Paulraj", Status: "deactive", Settings: &Settings{NotificationsEnabled: true}},
		{Name: "Vivek", Status: "active", Settings: &Settings{NotificationsEnabled: true}},
	}
	existingUsers := []*User{
		{Name: "Veera", Status: "active", Settings: &Settings{NotificationsEnabled: true}},
		{Name: "Gowtham", Status: "active", Settings: &Settings{NotificationsEnabled: true}},
		{Name: "Brani", Status: "active", Settings: &Settings{NotificationsEnabled: true}},
	}

	go filterNewUsersByStatus(usersToUpdate, newUsers)
	go updateUsers(usersToUpdate, userToNotify, existingUsers)
	notifyUsers(userToNotify, existingUsers)
}

func filterNewUsersByStatus(usersToUpdate chan<- []*User, users []*User) {
	defer close(usersToUpdate)
	filteredUsers := []*User{}
	for _, user := range users {
		if user.Status == "active" && user.Settings.NotificationsEnabled {
			filteredUsers = append(filteredUsers, user)
		}
	}

	usersToUpdate <- filteredUsers
}

func updateUsers(usersToUpdate <-chan []*User, userToNotify chan<- *User, users []*User) {
	defer close(userToNotify)
	for _, user := range users {
		user.Tags = append(user.Tags, &Tag{Name: "UserNotified", Type: "Notifications"})
	}
	for _, v := range <-usersToUpdate {
		fmt.Println("\n usersToUpdate in usersToUpdate func", v)
	}
	newUsers := <-usersToUpdate

	for _, user := range newUsers {
		time.Sleep(1 * time.Second)
		user.Tags = append(user.Tags, &Tag{Name: "NewNotification", Type: "Notifications"})
		userToNotify <- user
	}
}

func notifyUsers(userToNotify <-chan *User, users []*User) {
	service := &NotificationsService{}
	for _, user := range users {
		service.SendEmailNotification(user, "Tags", "A new tag has been added to your profile!!")
	}

	for user := range userToNotify {
		service.SendEmailNotification(user, "Tags", "You got your first tag!!")
	}
}

func (n *NotificationsService) SendEmailNotification(user *User, title, message string) {
	fmt.Printf("Email Notification Sent to %v, Hi %s, %s\n", user, user.Name, message)
}
