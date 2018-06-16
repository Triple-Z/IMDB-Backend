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

func CreateTitleASSQLString(queryString string, tconst string, isAdult string, startYearStart string, startYearEnd string, endYearStart string, endYearEnd string, runtimeMinStart string, runtimeMinEnd string) (SQLCountString string, SQLQueryString string) {

	var isFirstCondition = true

	basicSQLQueryString := "select id, tconst, Title_type, Primary_title, Original_title, Is_adult, Start_year, End_year, Runtime_minutes, Genres, Create_date, Last_updated from title_basics"

	basicSQLCountString := "select count(*) from title_basics"

	if tconst != "" {
		SQLQueryString = basicSQLQueryString + " where tconst = '" + tconst + "'"
		SQLCountString = basicSQLCountString + " where tconst = '" + tconst + "'"
		return SQLCountString, SQLQueryString
	}

	if queryString != "" {
		if isFirstCondition {
			SQLQueryString = basicSQLQueryString + " where match(Title_type, Primary_title, Original_title, Genres) against ('" + queryString + "' in boolean mode)"
			SQLCountString = basicSQLCountString + " where match(Title_type, Primary_title, Original_title, Genres) against ('" + queryString + "' in boolean mode)"
			isFirstCondition = false
		} else {
			SQLQueryString = SQLQueryString + " and match(Title_type, Primary_title, Original_title, Genres) against ('" + queryString + "' in boolean mode)"
			SQLCountString = SQLCountString + " and match(Title_type, Primary_title, Original_title, Genres) against ('" + queryString + "' in boolean mode)"
		}
	}

	if isAdult != "" {
		if isFirstCondition {
			SQLQueryString = basicSQLQueryString + " where Is_adult = " + isAdult
			SQLCountString = basicSQLCountString + " where Is_adult = " + isAdult
			isFirstCondition = false
		} else {
			SQLQueryString = SQLQueryString + " and Is_adult = " + isAdult
			SQLCountString = SQLCountString + " and Is_adult = " + isAdult
		}
	}
	// ----

	if startYearStart != "" {
		if isFirstCondition {
			SQLQueryString = basicSQLQueryString + " where Start_year >= " + startYearStart
			SQLCountString = basicSQLCountString + " where Start_year >= " + startYearStart
			isFirstCondition = false
		} else {
			SQLQueryString = SQLQueryString + " and Start_year >= " + startYearStart
			SQLCountString = SQLCountString + " and Start_year >= " + startYearStart
		}
	}

	if startYearEnd != "" {
		if isFirstCondition {
			SQLQueryString = basicSQLQueryString + " where Start_year <= " + startYearEnd
			SQLCountString = basicSQLCountString + " where Start_year <= " + startYearEnd
			isFirstCondition = false
		} else {
			SQLQueryString = SQLQueryString + " and Start_year <= " + startYearEnd
			SQLCountString = SQLCountString + " and Start_year <= " + startYearEnd
		}
	}
	// ----

	if endYearStart != "" {
		if isFirstCondition {
			SQLQueryString = basicSQLQueryString + " where End_year >= " + endYearStart
			SQLCountString = basicSQLCountString + " where End_year >= " + endYearStart
			isFirstCondition = false
		} else {
			SQLQueryString = SQLQueryString + " and End_year >= " + endYearStart
			SQLCountString = SQLCountString + " and End_year >= " + endYearStart
		}
	}

	if endYearEnd != "" {
		if isFirstCondition {
			SQLQueryString = basicSQLQueryString + " where End_year <= " + endYearEnd
			SQLCountString = basicSQLCountString + " where End_year <= " + endYearEnd
			isFirstCondition = false
		} else {
			SQLQueryString = SQLQueryString + " and End_year <= " + endYearEnd
			SQLCountString = SQLCountString + " and End_year <= " + endYearEnd
		}
	}
	// ----

	if runtimeMinStart != "" {
		if isFirstCondition {
			SQLQueryString = basicSQLQueryString + " where Runtime_minutes >= " + runtimeMinStart
			SQLCountString = basicSQLCountString + " where Runtime_minutes >= " + runtimeMinStart
			isFirstCondition = false
		} else {
			SQLQueryString = SQLQueryString + " and Runtime_minutes >= " + runtimeMinStart
			SQLCountString = SQLCountString + " and Runtime_minutes >= " + runtimeMinStart
		}
	}

	if runtimeMinEnd != "" {
		if isFirstCondition {
			SQLQueryString = basicSQLQueryString + " where Runtime_minutes <= " + runtimeMinEnd
			SQLCountString = basicSQLCountString + " where Runtime_minutes <= " + runtimeMinEnd
			isFirstCondition = false
		} else {
			SQLQueryString = SQLQueryString + " and Runtime_minutes <= " + runtimeMinEnd
			SQLCountString = SQLCountString + " and Runtime_minutes <= " + runtimeMinEnd
		}
	}

	if isFirstCondition {
		SQLQueryString = basicSQLQueryString
		SQLCountString = basicSQLCountString
	}

	return SQLCountString, SQLQueryString
}

func CreateNameASSQLString(queryString string, nconst string, birthYearStart string, birthYearEnd string, deathYearStart string, deathYearEnd string) (SQLCountString string, SQLQueryString string) {

	var isFirstCondition = true

	basicSQLQueryString := "select id, nconst, Primary_name, Birth_year, Death_year, Primary_profession, Known_for_titles, Create_date, Last_updated from name_basics"

	basicSQLCountString := "select count(*) from name_basics"

	if nconst != "" {
		SQLQueryString = basicSQLQueryString + " where nconst = '" + nconst + "'"
		SQLCountString = basicSQLCountString + " where nconst = '" + nconst + "'"
		return SQLCountString, SQLQueryString
	}

	if queryString != "" {
		if isFirstCondition {
			SQLQueryString = basicSQLQueryString + " where match(Primary_name, Primary_profession) against ('" + queryString + "' in boolean mode)"
			SQLCountString = basicSQLCountString + " where match(Primary_name, Primary_profession) against ('" + queryString + "' in boolean mode)"
			isFirstCondition = false
		} else {
			SQLQueryString = SQLQueryString + " and match(Primary_name, Primary_profession) against ('" + queryString + "' in boolean mode)"
			SQLCountString = SQLCountString + " and match(Primary_name, Primary_profession) against ('" + queryString + "' in boolean mode)"
		}
	}

	if birthYearStart != "" {
		if isFirstCondition {
			SQLQueryString = basicSQLQueryString + " where Birth_year >= " + birthYearStart
			SQLCountString = basicSQLCountString + " where Birth_year >= " + birthYearStart
			isFirstCondition = false
		} else {
			SQLQueryString = SQLQueryString + " and Birth_year >= " + birthYearStart
			SQLCountString = SQLCountString + " and Birth_year >= " + birthYearStart
		}
	}

	if birthYearEnd != "" {
		if isFirstCondition {
			SQLQueryString = basicSQLQueryString + " where Birth_year <= " + birthYearEnd
			SQLCountString = basicSQLCountString + " where Birth_year <= " + birthYearEnd
			isFirstCondition = false
		} else {
			SQLQueryString = SQLQueryString + " and Birth_year <= " + birthYearEnd
			SQLCountString = SQLCountString + " and Birth_year <= " + birthYearEnd
		}
	}
	// ----

	if deathYearStart != "" {
		if isFirstCondition {
			SQLQueryString = basicSQLQueryString + " where Death_year >= " + deathYearStart
			SQLCountString = basicSQLCountString + " where Death_year >= " + deathYearStart
			isFirstCondition = false
		} else {
			SQLQueryString = SQLQueryString + " and Death_year >= " + deathYearStart
			SQLCountString = SQLCountString + " and Death_year >= " + deathYearStart
		}
	}

	if deathYearEnd != "" {
		if isFirstCondition {
			SQLQueryString = basicSQLQueryString + " where Death_year <= " + deathYearEnd
			SQLCountString = basicSQLCountString + " where Death_year <= " + deathYearEnd
			isFirstCondition = false
		} else {
			SQLQueryString = SQLQueryString + " and Death_year <= " + deathYearEnd
			SQLCountString = SQLCountString + " and Death_year <= " + deathYearEnd
		}
	}

	if isFirstCondition {
		SQLQueryString = basicSQLQueryString
		SQLCountString = basicSQLCountString
	}

	return SQLCountString, SQLQueryString
}
