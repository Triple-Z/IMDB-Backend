package app

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
)

func DbInit() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)?parseTime=true&charset=utf8")
	checkNormalError(err)

	err = db.Ping()
	if err != nil {
		log.Fatalln("Cannot connect to database !")
		panic(err.Error())
	}

	log.Println("Connect to database successfully ~")

	return db
}

var db = DbInit()

