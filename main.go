package main

import (
	"Appointy/Tasks/todo/db"
	"Appointy/Tasks/todo/router"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("local.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	fmt.Println("Server started at http://localhost:" + os.Getenv("PORT"))
	DB, _ := db.Connect()

	h := db.New(DB)

	router := router.CreateRouter(h)
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
