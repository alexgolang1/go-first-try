package main

import (
	"back-api/internal/app/repository"
	"back-api/internal/pkg/app"
)

func main() {
	database := repository.InitDB()
	rdb := repository.InitRDB()
	a := app.New(database, rdb)
	a.Run()
}
