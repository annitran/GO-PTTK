package main

import (
	"GO-PTTK/config"
	"GO-PTTK/migrations"
	"GO-PTTK/routes"
)

func main() {
	config.ConnectDB()
	migrations.Migrate()
	routes.SetupRouter().Run(":8080")
}
