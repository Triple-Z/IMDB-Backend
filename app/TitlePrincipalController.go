package app

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"net/http"
	"time"
)

type TitleDetailSQL struct {
	Ordering sql.NullInt64
	PrimaryName sql.NullString
	Category    sql.NullString
	Job         sql.NullString
	Characters  sql.NullString
}

type TitlePrincipalSQL struct {
	Id          sql.NullInt64
	TConst      sql.NullString
	Ordering    sql.NullInt64
	NConst      sql.NullString
	Category    sql.NullString
	Job         sql.NullString
	Characters  sql.NullString
	CreateDate  mysql.NullTime
	LastUpdated mysql.NullTime
}

type TitlePrincipal struct {
	TConst     string `form:"tconst" json:"tconst" binding:"required"`
	Ordering   int    `form:"ordering" json:"ordering" binding:"required"`
	NConst     string `form:"nconst" json:"nconst" binding:"required"`
	Category   string `form:"category" json:"category"`
	Job        string `form:"job" json:"job"`
	Characters string `form:"characters" json:"characters"`
}

func ReadTitleDetails(c *gin.Context) {
	var (
		titleDetailSQL TitleDetailSQL
		titleDetailSQLs []TitleDetailSQL
	)

	id := c.Params.ByName("id")

	// No need for pagination

	rows, err := db.Query("select title_principals.Ordering, name_basics.Primary_name, title_principals.Category. title_principals.Job, title_principals.Characters from title_basics join title_principals on title_basics.tconst = title_principals.tconst join name_basics on title_principals.nconst = name_basics.nconst where title_basics.id=? order by Ordering", id)

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
		err := rows.Scan(&titleDetailSQL.Ordering, &titleDetailSQL.PrimaryName, &titleDetailSQL.Category, &titleDetailSQL.Job, &titleDetailSQL.Characters)
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
		titleDetailSQLs = append(titleDetailSQLs, titleDetailSQL)
		//log.Print(titleDetailSQL)
	}

	c.JSON(http.StatusOK, gin.H{
		"server_time": time.Now(),
		"count":       len(titleDetailSQLs),
		"data":        titleDetailSQLs,
	})

}

func ReadSinglePrincipal(c *gin.Context) {
	var titlePrincipalSQL TitlePrincipalSQL

	id := c.Params.ByName("id")

	row := db.QueryRow("select id, tconst, Ordering, nconst, Category, Job, Characters, Create_date, Last_updated from title_principals where id = ?", id)

	err := row.Scan(&titlePrincipalSQL.Id, &titlePrincipalSQL.TConst, &titlePrincipalSQL.Ordering, &titlePrincipalSQL.NConst, &titlePrincipalSQL.Category, &titlePrincipalSQL.Job, &titlePrincipalSQL.Characters, &titlePrincipalSQL.CreateDate, &titlePrincipalSQL.LastUpdated)

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
		"data":        titlePrincipalSQL,
	})

}

func CreatePrincipal(c *gin.Context) {
	var (
		titlePrincipal    TitlePrincipal
		titlePrincipalSQL TitlePrincipalSQL
	)

	if err := c.ShouldBind(&titlePrincipal); err == nil {

		stmt, err := db.Prepare("insert into title_principals(tconst, Ordering, nconst, Category, Job, Characters) VALUES (?, ?, ?, ?, ?, ?)")
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

		titlePrincipalSQL = ValidateTitlePrincipalStruct(&titlePrincipal)

		status, err := stmt.Exec(titlePrincipal.TConst, titlePrincipal.Ordering, titlePrincipal.NConst, titlePrincipalSQL.Category, titlePrincipalSQL.Job, titlePrincipalSQL.Characters)
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

func UpdatePrincipal(c *gin.Context) {
	id := c.Params.ByName("id")

	var (
		titlePrincipal    TitlePrincipal
		titlePrincipalSQL TitlePrincipalSQL
	)

	if err := c.ShouldBind(&titlePrincipal); err == nil {

		stmt, err := db.Prepare("update title_principals set tconst=?, Ordering=?, nconst=?, Category=?, Job=?, Characters=? where id=?")
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

		titlePrincipalSQL = ValidateTitlePrincipalStruct(&titlePrincipal)

		status, err := stmt.Exec(titlePrincipal.TConst, titlePrincipal.Ordering, titlePrincipal.NConst, titlePrincipalSQL.Category, titlePrincipalSQL.Job, titlePrincipalSQL.Characters, id)
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

func DeletePrincipal(c *gin.Context) {
	id := c.Params.ByName("id")

	stmt, err := db.Prepare("delete from title_principals where id = ?")
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
