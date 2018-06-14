package app

import (
	"github.com/go-sql-driver/mysql"
	"log"
)

func checkSQLError(err error) uint16 {
	// check SQL driver error
	if driverErr, ok := err.(*mysql.MySQLError); ok {
		log.Fatalln(err.Error())
		return driverErr.Number
	}

	// check normal & SQL error
	if err != nil {
		log.Fatalln(err.Error())
		return 1
	}

	// Normal, no error
	return 0
}

func checkNormalError(err error) bool {
	if err != nil {
		log.Fatalln(err.Error())
		return true
	} else {
		return false
	}
}
