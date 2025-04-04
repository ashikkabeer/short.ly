package main

import (
	"fmt"
	"log"

	"github.com/ashikkabeer/short.ly/cmd/api"
	"github.com/ashikkabeer/short.ly/config/db"
	"github.com/ashikkabeer/short.ly/internal/handler"
)

func main() {

	// starting the database
	err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database Connected")
	
	// initialize the handler with repository and service
	err = handler.InitHandler()
	if err != nil {
		log.Fatal("Failed to initialize handler:", err)
	}
	
	fmt.Println("Hello, Welcome to Short.ly")
	
	// setting up the API Routes
	r := api.SetupAPI()

	r.Run()
}
