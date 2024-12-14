package db

import (
	"fmt"
	"log"

	// "github.com/and3407/go_reports/app/db/models"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Connect gorm.DB

func Init() {
	Connect = connectPG()
}

func connectPG() gorm.DB {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	user := viper.Get("DB_USER").(string)
	password := viper.Get("DB_PASSWORD").(string)
	host := viper.Get("DB_HOST").(string)
	port := viper.Get("DB_PORT").(string)
	dbName := viper.Get("DB_NAME").(string)


	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		user,
		password,
		host,
		port,
		dbName,
	)

	fmt.Println(url)

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	// db.AutoMigrate(&models.Report{})
	// db.AutoMigrate(&models.Task{})

	return *db
}