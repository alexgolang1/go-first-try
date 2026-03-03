package main

import (
	"back-api/internal/app/db"
	"back-api/internal/pkg/app"
)

func main() {
	database := db.InitDB()
	a := app.New(database)
	a.Run()
}
