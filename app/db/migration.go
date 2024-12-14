package db

import (
	"fmt"
	"log"
	
	"github.com/spf13/viper"
	"github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/pkger"
)

func MigrateUp() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	user := viper.Get("DB_USER").(string)
	password := viper.Get("DB_PASSWORD").(string)
	host := viper.Get("DB_HOST").(string)
	port := viper.Get("DB_PORT").(string)
	dbName := viper.Get("DB_NAME").(string)

	pgUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		password,
		host,
		port,
		dbName,
	)

	sourceUrl := "pkger:///app/db/migrations"

    m, err := migrate.New(sourceUrl, pgUrl)

	if err != nil {
		log.Println("Migrate error", err)

		return
	}

    m.Up()
}