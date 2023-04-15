package main

import (
	"os"

	"github.com/ferytell/go-jwt/database"
	"github.com/ferytell/go-jwt/initializer"
	"github.com/ferytell/go-jwt/router"
)

func init() {
	initializer.LoadEnvVar()
}

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(os.Getenv("PORT"))

}
