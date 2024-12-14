package database

import (
	"fmt"
	"kredit-plus/internal/models"
	"kredit-plus/pkg/mysql"
	// "gorm.io/plugin/dbresolver"
)

func RunMigration() {
	var err error
	// main migration
	err = mysql.DB.AutoMigrate(
		&models.Consumer{},
		&models.CreditLimit{},
		&models.Patner{},
		&models.TransactionDetail{},
		// &models.MyUser{},
	)

	if err != nil {
		fmt.Println(err)
		panic("DB Migration Failed ")
	}

	fmt.Println("All Migration Success")
}

//sample handling multi database
// err = mysql.DB.Clauses(dbresolver.Use("usr21")).AutoMigrate(&models.Myuser{})

// if err != nil {
// 	fmt.Println(err)
// 	panic("DB Migration Failed")
// }

// err = mysql.DB.Clauses(dbresolver.Use("img21")).AutoMigrate(&models.ImageUser{}, &models.ImageStnk{}, &models.ImageSubPenghuni{})
// if err != nil {
// 	fmt.Println(err)
// 	panic("DB Migration Failed")
// }
