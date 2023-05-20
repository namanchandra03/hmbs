package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"hms/database"
	"hms/server"
	"log"
)

func main() {

	done := make(chan int)
	srv := server.SetUpRoutes()

	if err := database.ConnectAndMigrate(
		"localhost",
		"5433",
		"postgres",
		"local",
		"local",
		database.SSLModeDisable); err != nil {
		logrus.Panicf("Failed to initialize and migrate database with error: %+v", err)
	}
	logrus.Print("migration successful!!")
	go func() {
		err := srv.Start(":3000")
		if err != nil {
			log.Fatalf("Failed to run server with error: %+v", err)
		}

	}()

	fmt.Println("Server Running at port :3000")
	<-done
}
