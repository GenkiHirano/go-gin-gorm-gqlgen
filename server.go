package main

import (
	"log"
	"os"

	"github.com/GenkiHirano/gqlgen-tutorial/db"
	"github.com/GenkiHirano/gqlgen-tutorial/router"
	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func initServer() *gin.Engine {
	server := gin.Default()
	router.Setup(server)
	return server
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db.InitDB()
	defer db.CloseDB()

	server := initServer()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(server.Run(":8000"))
}
