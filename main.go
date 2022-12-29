package main

import (
	"github.com/tryhd/dbo-test/app/config"
	"github.com/tryhd/dbo-test/app/database"
	"github.com/tryhd/dbo-test/app/router"
)

func main() {
	db := config.SetupDatabaseConnection()

	database.Migrator(db)
	router.Router()
}
