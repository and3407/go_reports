package main

import (
	"github.com/and3407/go_reports/app/db"
	"github.com/and3407/go_reports/app/http/server"
)

func main() {
	db.MigrateUp()
	db.Init()
	server.Init()
}