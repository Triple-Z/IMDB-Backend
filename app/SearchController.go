package app

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"math"
	"net/http"
	"time"
)

func FuzzySearchTitles(c *gin.Context) {
	var (
		queryPage int
		prevPage  int
		nextPage  int
		totalRows int
		title     TitleBasicsSQL
		titles    []TitleBasicsSQL
	)


	queryString := c.Query("q")
	queryPageString := c.Query("page")

	if queryPageString == "" {
		queryPage = 1
	} else {
		tem, err := strconv.Atoi(queryPageString)
		checkNormalError(err)
		queryPage = tem
	}

	row := db.QueryRow("select count(*) from title_basics where match (Title_type, Primary_title, Original_title, Genres) against ('?' in natural language mode)", queryString)
	row.Scan(&totalRows)

	totalPages := int(math.Ceil(float64(totalRows) / rowsTitlePerPage))

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

	startRow := (queryPage - 1) * rowsTitlePerPage

	rows, err := db.Query("select * from title_basics where match (Title_type, Primary_title, Original_title, Genres) against (? in natural language mode) limit ?, ?", queryString, startRow, rowsTitlePerPage)
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
		"start_id":    startRow,
		"count":       len(titles),
		"cur_page":    queryPage,
		"next_page":   nextPage,
		"prev_page":   prevPage,
		"total_page":  totalPages,
		"server_time": time.Now(),
		"data":        titles,
	})
}

func FuzzySearchNames(c *gin.Context) {

}
