package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-todo/config"
	"go-todo/routes"
	"log"
)

//TODO
// dsn, createTodo

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()

	r := gin.Default()
	routes.RegisterRotes(r)

	err = r.Run(":8080")
	if err != nil {
		return
	}
}
