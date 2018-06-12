package app

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"math"
	"net/http"
	"strconv"
	"time"
)

type NameBasicsSQL struct {
	Id                sql.NullInt64
	NConst            sql.NullString
	PrimaryName       sql.NullString
	BirthYear         sql.NullInt64
	DeathYear         sql.NullInt64
	PrimaryProfession sql.NullString
	KnownForTitles    sql.NullString
	CreateDate        mysql.NullTime
	LastUpdated       mysql.NullTime
}

type NameBasics struct {
	NConst            string `form:"nconst" json:"nconst" binding:"required"`
	PrimaryName       string `form:"primary_name" json:"primary_name" binding:"required"`
	BirthYear         int    `form:"birth_year" json:"birth_year"`
	DeathYear         int    `form:"death_year" json:"death_year"`
	PrimaryProfession string `form:"primary_profession" json:"primary_profession"`
	KnownForTitles    string `form:"known_for_titles" json:"known_for_titles"`
}

func ReadAllNames(c *gin.Context) {

	var (
		queryPage int
		prevPage  int
		nextPage  int
		totalRows int
		name      NameBasicsSQL
		names     []NameBasicsSQL
	)

	queryPageString := c.Query("page")
	if queryPageString == "" {
		queryPage = 1
	} else {
		tem, err := strconv.Atoi(queryPageString)
		checkNormalError(err)
		queryPage = tem
	}

	row := db.QueryRow("select count(*) from name_basics")
	row.Scan(&totalRows)

	totalPages := int(math.Ceil(float64(totalRows) / rowsNamePerPage))

	if queryPage == totalPages {
		nextPage = -1
	} else {
		nextPage = queryPage + 1
	}

	if queryPage == 1 {
		prevPage = -1
	} else {
		prevPage = queryPage - 1
	}

	startId := (queryPage - 1) * rowsNamePerPage

	rows, err := db.Query("select id, nconst, Primary_name, Birth_year, Death_year, Primary_profession, Known_for_titles, Create_date, Last_updated from name_basics where id > ? limit ?", startId, rowsNamePerPage)

	if errCode := checkSQLError(err); errCode != 0 {
		switch errCode {
		case 1:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Unknown error",
			})
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
		return
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&name.Id, &name.NConst, &name.PrimaryName, &name.BirthYear, &name.DeathYear, &name.PrimaryProfession, &name.KnownForTitles, &name.CreateDate, &name.LastUpdated)
		if errCode := checkSQLError(err); errCode != 0 {
			switch errCode {
			case 1:
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Unknown error",
				})
			default:
				c.JSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
				})
			}
			return
		}
		names = append(names, name)
		//log.Print(name)
	}

	c.JSON(http.StatusOK, gin.H{
		"start_id":    startId,
		"count":       len(names),
		"cur_page":    queryPage,
		"next_page":   nextPage,
		"prev_page":   prevPage,
		"total_page":  totalPages,
		"server_time": time.Now(),
		"data":        names,
	})

}

func ReadSingleName(c *gin.Context) {
	var (
		name NameBasicsSQL
	)

	id := c.Params.ByName("id")

	row := db.QueryRow("select id, nconst, Primary_name, Birth_year, Death_year, Primary_profession, Known_for_titles, Create_date, Last_updated from name_basics where id = ?", id)

	err := row.Scan(&name.Id, &name.NConst, &name.PrimaryName, &name.BirthYear, &name.DeathYear, &name.PrimaryProfession, &name.KnownForTitles, &name.CreateDate, &name.LastUpdated)

	if errCode := checkSQLError(err); errCode != 0 {
		switch errCode {
		case 1:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Unknown error",
			})
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"server_time": time.Now(),
		"data":        name,
	})
}

func CreateName(c *gin.Context) {
	var (
		name    NameBasics
		nameSQL NameBasicsSQL
	)

	if err := c.ShouldBind(&name); err == nil {

		stmt, err := db.Prepare("insert into name_basics(nconst, Primary_name, Birth_year, Death_year, Primary_profession, Known_for_titles)  values (?, ?, ?, ?, ?, ?)")
		if errCode := checkSQLError(err); errCode != 0 {
			switch errCode {
			case 1:
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Unknown error",
				})
			default:
				c.JSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
				})
			}
			return
		}
		defer stmt.Close()

		nameSQL = ValidateNameStruct(&name)

		status, err := stmt.Exec(name.NConst, name.PrimaryName, nameSQL.BirthYear, nameSQL.DeathYear, nameSQL.PrimaryProfession, nameSQL.KnownForTitles)
		if errCode := checkSQLError(err); errCode != 0 {
			switch errCode {
			case 1:
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Unknown error",
				})
			default:
				c.JSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
				})
			}
			return
		}

		insertId, err := status.LastInsertId()
		checkSQLError(err)

		c.JSON(http.StatusOK, gin.H{
			"server_time": time.Now(),
			"insert_id":   insertId,
		})

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
}

func UpdateName(c *gin.Context) {
	id := c.Params.ByName("id")

	var (
		name    NameBasics
		nameSQL NameBasicsSQL
	)

	if err := c.ShouldBind(&name); err == nil {

		stmt, err := db.Prepare("update name_basics set nconst = ?, Primary_name = ?, Birth_year = ?, Death_year = ?, Primary_profession = ?, Known_for_titles = ? where id = ?")
		if errCode := checkSQLError(err); errCode != 0 {
			switch errCode {
			case 1:
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Unknown error",
				})
			default:
				c.JSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
				})
			}
			return
		}
		defer stmt.Close()

		nameSQL = ValidateNameStruct(&name)

		status, err := stmt.Exec(name.NConst, name.PrimaryName, nameSQL.BirthYear, nameSQL.DeathYear, nameSQL.PrimaryProfession, nameSQL.KnownForTitles, id)
		if errCode := checkSQLError(err); errCode != 0 {
			switch errCode {
			case 1:
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Unknown error",
				})
			default:
				c.JSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
				})
			}
			return
		}

		rowsAffected, err := status.RowsAffected()
		checkSQLError(err)

		if rowsAffected != 0 {
			c.JSON(http.StatusOK, gin.H{
				"server_time": time.Now(),
				"updated_id":  id,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("No update action detected with id: %s", id),
			})
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
}

func DeleteName(c *gin.Context) {

	id := c.Params.ByName("id")

	stmt, err := db.Prepare("delete from name_basics where id = ?")
	if errCode := checkSQLError(err); errCode != 0 {
		switch errCode {
		case 1:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Unknown error",
			})
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
		return
	}
	defer stmt.Close()

	status, err := stmt.Exec(id)
	if errCode := checkSQLError(err); errCode != 0 {
		switch errCode {
		case 1:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Unknown error",
			})
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
		return
	}

	rowsAffected, err := status.RowsAffected()
	checkSQLError(err)

	if rowsAffected != 0 {
		c.Status(http.StatusNoContent)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("No delete action detected with id: %s", id),
		})
	}

}
