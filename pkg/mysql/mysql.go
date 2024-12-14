package mysql

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func DataBaseinit() {
	var err error

	// env database relation
	var dbconfig = os.Getenv("DB_CHARNTIME")
	var mst = os.Getenv("DB_MST")

	// env DATABASE HOST
	var dbuser = os.Getenv("DB_USER")
	var dbpass = os.Getenv("DB_PASS")
	var dbhost = os.Getenv("DB_HOST")
	var dbport = os.Getenv("DB_PORT")

	var mysqlconfig = dbuser + ":" + dbpass + "@tcp(" + dbhost + ":" + dbport + ")/"
	var gormConfig = &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}}

	DB, err = gorm.Open(mysql.Open(fmt.Sprintf("%s%s%s", mysqlconfig, mst, dbconfig)), gormConfig)
	if err != nil {
		fmt.Println("connect to database mst failed")
		panic(err)
	}

	// Retrieve the underlying sql.DB and configure the connection pool
	sqlDB, err := DB.DB()
	if err != nil {
		fmt.Println("get underlying sql.DB failed")
		panic(err)
	}
	// Set the connection pool parameters
	sqlDB.SetConnMaxIdleTime(time.Hour)
	sqlDB.SetConnMaxLifetime(24 * time.Hour)
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(200)

	fmt.Println("connected to database")
}

// setup multidatabase

// DB.Use(dbresolver.Register(dbresolver.Config{
// 	Sources: []gorm.Dialector{mysql.Open(fmt.Sprintf("%s%s%s", mysqlconfig, usr, dbconfig))},
// }, &models.Myuser{}, "usr21").Register(dbresolver.Config{
// 	Sources: []gorm.Dialector{mysql.Open(fmt.Sprintf("%s%s%s", mysqlconfig, img, dbconfig))},
// }, &models.ImageUser{}, &models.ImagePenghuni{}, &models.ImageStnk{}, "img21").
// 	SetConnMaxIdleTime(time.Hour).
// 	SetConnMaxLifetime(24 * time.Hour).
// 	SetMaxIdleConns(100).
// 	SetMaxOpenConns(200),
// )
