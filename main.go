package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/LamichhaneBibek/multiDB/database"
	"github.com/LamichhaneBibek/multiDB/model"
	"github.com/LamichhaneBibek/multiDB/repo"
)

func init() {
	// Initialize database connections during program startup.
	database.InitDBConnections()
}

func main() {
	users := []model.User{
		{UserName: "user1", Password: "password1"},
		{UserName: "user2", Password: "password2"},
		{UserName: "user3", Password: "password3"},
		{UserName: "user4", Password: "password4"},
		{UserName: "user5", Password: "password5"},
		{UserName: "user6", Password: "password6"},
		{UserName: "user7", Password: "password7"},
		{UserName: "user8", Password: "password8"},
		{UserName: "user9", Password: "password9"},
		{UserName: "user10", Password: "password10"},
		{UserName: "user11", Password: "password11"},
		{UserName: "user12", Password: "password12"},
		{UserName: "user13", Password: "password13"},
		{UserName: "user14", Password: "password14"},
		{UserName: "user15", Password: "password15"},
	}

	// Create a new UserRepository with a specified connection name
	repo := repo.NewUserRepo("TEST_SQLITE_CON") // TEST_POSTGRES_CON, TEST_MYSQL_CON, TEST_SQLITE_CON

	// save users
	err := repo.Save(users...)
	if err != nil {
		return
	}

	// get all users
	users, err = repo.FindAll()
	if err != nil {
		return
	}
	// print users
	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}

	// Set up a channel to listen for OS signals (e.g., Ctrl+C)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	// Wait for a signal (e.g., Ctrl+C) to gracefully exit the program
	<-c

	// Close database connections when the program terminates.
	defer database.CloseDBConnections()
}
