package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func DbInit() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/imdb?parseTime=true&charset=utf8")
	checkNormalError(err)

	err = db.Ping()
	if err != nil {
		log.Fatalln("Cannot connect to database !")
		panic(err.Error())
	}

	log.Println("Connect to database successfully ~")

	//db.Exec("use imdb")

	return db
}

var db = DbInit()
