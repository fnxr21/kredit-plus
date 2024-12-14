package database

import (
	"fmt"
	"kredit-plus/internal/models"
	"kredit-plus/pkg/mysql"
)

func RunMigration() {
	var err error
	// main migration
	err = mysql.DB.AutoMigrate(
		&models.Consumer{},
		&models.CreditLimit{},
		&models.Partner{},
		&models.TransactionDetail{},
		&models.MyUser{},
	)

	if err != nil {
		fmt.Println(err)
		panic("DB Migration Failed ")
	}

	fmt.Println("All Migration Success")
}
