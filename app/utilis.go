package app

import (
	"github.com/go-sql-driver/mysql"
	"log"
)

func checkSQLError(err error) uint16 {
	if driverErr, ok := err.(*mysql.MySQLError); ok {
		log.Fatalln(err.Error())
		return driverErr.Number
	} else {
		return 0
	}
}

func checkNormalError(err error) bool {
	if err != nil {
		log.Fatalln(err.Error())
		return true
	} else {
		return false
	}
}
