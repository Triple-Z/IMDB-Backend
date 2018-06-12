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

type TitleBasicsSQL struct {
	Id             sql.NullInt64
	TConst         sql.NullString
	TitleType      sql.NullString
	PrimaryTitle   sql.NullString
	OriginalTitle  sql.NullString
	IsAdult        sql.NullBool
	StartYear      sql.NullInt64
	EndYear        sql.NullInt64
	RuntimeMinutes sql.NullInt64
	Genres         sql.NullString
	CreateDate     mysql.NullTime
	LastUpdated    mysql.NullTime
}

type TitleBasics struct {
	TConst         string `form:"tconst" json:"tconst" binding:"required"`
	TitleType      string `form:"title_type" json:"title_type" binding:"required"`
	PrimaryTitle   string `form:"primary_title" json:"primary_title" binding:"required"`
	OriginalTitle  string `form:"original_title" json:"original_title"`
	IsAdult        *bool  `form:"is_adult" json:"is_adult" binding:"exists"`
	StartYear      int    `form:"start_year" json:"start_year"`
	EndYear        int    `form:"end_year" json:"end_year"`
	RuntimeMinutes int    `form:"runtime_minutes" json:"runtime_minutes"`
	Genres         string `form:"genres" json:"genres"`
}

func ReadAllTitles(c *gin.Context) {
	var (
		queryPage int
		prevPage  int
		nextPage  int
		totalRows int
		title     TitleBasicsSQL
		titles    []TitleBasicsSQL
	)

	queryPageString := c.Query("page")
	if queryPageString == "" {
		queryPage = 1
	} else {
		tem, err := strconv.Atoi(queryPageString)
		checkNormalError(err)
		queryPage = tem
	}

	row := db.QueryRow("select count(*) from test_title_basics")
	row.Scan(&totalRows)

	totalPages := int(math.Ceil(float64(totalRows) / rowsPerPage))

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

	startId := (queryPage - 1) * rowsPerPage

	rows, err := db.Query("select id, tconst, titleType, primaryTitle, originalTitle, isAdult, startYear, endYear, runtimeMinutes, genres, createDate, lastUpdated from test_title_basics where id > ? limit ?", startId, rowsPerPage)

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
		err := rows.Scan(&title.Id, &title.TConst, &title.TitleType, &title.PrimaryTitle, &title.OriginalTitle, &title.IsAdult, &title.StartYear, &title.EndYear, &title.RuntimeMinutes, &title.Genres, &title.CreateDate, &title.LastUpdated)
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
		titles = append(titles, title)
		//log.Print(title)
	}

	c.JSON(http.StatusOK, gin.H{
		"start_id":    startId,
		"count":       len(titles),
		"cur_page":    queryPage,
		"next_page":   nextPage,
		"prev_page":   prevPage,
		"total_page":  totalPages,
		"server_time": time.Now(),
		"data":        titles,
	})
}

func ReadSingleTitle(c *gin.Context) {
	var (
		title TitleBasicsSQL
	)

	id := c.Params.ByName("id")

	row := db.QueryRow("select id, tconst, titleType, primaryTitle, originalTitle, isAdult, startYear, endYear, runtimeMinutes, genres, createDate, lastUpdated from test_title_basics where id = ?", id)

	err := row.Scan(&title.Id, &title.TConst, &title.TitleType, &title.PrimaryTitle, &title.OriginalTitle, &title.IsAdult, &title.StartYear, &title.EndYear, &title.RuntimeMinutes, &title.Genres, &title.CreateDate, &title.LastUpdated)

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
		"data":        title,
	})
}

func CreateTitle(c *gin.Context) {
	var (
		title    TitleBasics
		titleSQL TitleBasicsSQL
	)

	if err := c.ShouldBind(&title); err == nil {

		stmt, err := db.Prepare("insert into test_title_basics(tconst, titleType, primaryTitle, originalTitle, isAdult, startYear, endYear, runtimeMinutes, genres) values (?, ?, ?, ?, ?, ?, ?, ?, ?)")
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

		titleSQL = ValidateTitleStruct(&title)

		status, err := stmt.Exec(title.TConst, title.TitleType, title.PrimaryTitle, titleSQL.OriginalTitle, title.IsAdult, titleSQL.StartYear, titleSQL.EndYear, titleSQL.RuntimeMinutes, titleSQL.Genres)
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

func UpdateTitle(c *gin.Context) {

	id := c.Params.ByName("id")

	var (
		title    TitleBasics
		titleSQL TitleBasicsSQL
	)

	if err := c.ShouldBind(&title); err == nil {

		stmt, err := db.Prepare("update test_title_basics set tconst = ?, titleType = ?, primaryTitle = ?, originalTitle = ?, isAdult = ?, startYear = ?, endYear = ?, runtimeMinutes = ?, genres = ? where id = ?;")
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

		titleSQL = ValidateTitleStruct(&title)

		status, err := stmt.Exec(title.TConst, title.TitleType, title.PrimaryTitle, titleSQL.OriginalTitle, title.IsAdult, titleSQL.StartYear, titleSQL.EndYear, titleSQL.RuntimeMinutes, titleSQL.Genres, id)
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

		rows_affected, err := status.RowsAffected()
		checkSQLError(err)

		if rows_affected != 0 {
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

func DeleteTitle(c *gin.Context) {

	id := c.Params.ByName("id")

	stmt, err := db.Prepare("delete from test_title_basics where id = ?")
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

	rows_affected, err := status.RowsAffected()
	checkSQLError(err)

	if rows_affected != 0 {
		c.Status(http.StatusNoContent)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("No delete action detected with id: %s", id),
		})
	}

}
