package main

import (
	"github.com/ferytell/go-jwt/database"
	_ "github.com/ferytell/go-jwt/docs"
	"github.com/ferytell/go-jwt/initializer"
	"github.com/ferytell/go-jwt/router"
)

// @title 	Tag Service API
// @version	1.0
// @description A Tag service API in Go using Gin framework
// @host localhost:8000
// @securityDefinitions.apiKey BearerAuth
// @in header
// @name Authorization
// @BasePath /

func init() {
	// Env Varible
	initializer.LoadEnvVar()
}

func main() {
	// Database
	database.StartDB()
	var PORT = ":8000"

	// Start Router
	r := router.StartApp()
	r.Run(PORT)

}
