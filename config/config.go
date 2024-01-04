package config

import (
	"fmt"
	"fundhub-api/campaign"
	"fundhub-api/helper"
	"fundhub-api/transaction"
	"fundhub-api/user"

	"github.com/labstack/gommon/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsnMaster := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		helper.GoDotEnvVariable("DB_USER"),
		helper.GoDotEnvVariable("DB_PASSWORD"),
		helper.GoDotEnvVariable("DB_HOST"),
		helper.GoDotEnvVariable("DB_PORT"),
		helper.GoDotEnvVariable("DB_NAME"),
	)

	// dsnMaster := helper.GoDotEnvVariable("DATABASE_URL")
	db, err := gorm.Open(mysql.Open(dsnMaster), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	if err = db.AutoMigrate(
		&user.User{},
		&campaign.Campaign{},
		&campaign.CampaignImage{},
		&transaction.Transaction{},
	); err != nil {
		log.Fatal(err)
	}

	log.Info("success connect to database")

	return db
}
