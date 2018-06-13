package app

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"math"
	"net/http"
	"time"
	"database/sql"
	"log"
)

func FuzzySearchTitles(c *gin.Context) {
	var (
		queryPage int
		prevPage  int
		nextPage  int
		totalRows int
		title     TitleBasicsSQL
		titles    []TitleBasicsSQL
		nullScore sql.NullFloat64
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

	row := db.QueryRow("select count(*) from title_basics where match (Title_type, Primary_title, Original_title, Genres) against (? in natural language mode)", queryString)
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

	rows, err := db.Query("select *, match(Primary_title, Original_title) against (? in natural language mode) as rel_title, match(Genres) against (? in natural language mode) as rel_genres, match(Title_type) against (? in natural language mode) as rel_titleType from title_basics where match(Title_type, Primary_title, Original_title, Genres) against (? in natural language mode) order by (rel_title*3)+(rel_genres*2)+(rel_titleType) DESC limit ?, ?", queryString, queryString, queryString, queryString, startRow, rowsTitlePerPage)
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
		err := rows.Scan(&title.Id, &title.TConst, &title.TitleType, &title.PrimaryTitle, &title.OriginalTitle, &title.IsAdult, &title.StartYear, &title.EndYear, &title.RuntimeMinutes, &title.Genres, &title.CreateDate, &title.LastUpdated, &nullScore, &nullScore, &nullScore)
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

	if len(titles) == 0 && FuzzyBrutalForceSearch {
		// Fulltext search failed
		// Brute-force method

		row := db.QueryRow("select count(*) from title_basics	where Primary_title like concat('%', ?, '%') or  Original_title like concat('%', ?, '%') or  Genres like concat('%', ?, '%') or Title_type like concat('%', ?, '%')", queryString, queryString, queryString, queryString)
		row.Scan(&totalRows)
		log.Println(totalRows)

		totalPages = int(math.Ceil(float64(totalRows) / rowsTitlePerPage))

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

		rows, err := db.Query("select * from title_basics 	where Primary_title like concat('%', ?, '%') or  Original_title like concat('%', ?, '%') or  Genres like concat('%', ?, '%') or Title_type like concat('%', ?, '%') order by Primary_title like concat(?, '%') desc, ifnull(nullif(instr(Primary_title, concat(' ', ?)), 0), 99999), ifnull(nullif(instr(Primary_title, ?), 0), 99999), Primary_title, Original_title like concat(?, '%') desc, ifnull(nullif(instr(Original_title, concat(' ', ?)), 0), 99999), ifnull(nullif(instr(Original_title, ?), 0), 99999), Original_title,	Genres like concat(?, '%') desc, ifnull(nullif(instr(Genres, concat(' ', ?)), 0), 99999), ifnull(nullif(instr(Genres, ?), 0), 99999), Genres, Title_type like concat(?, '%') desc, ifnull(nullif(instr(Title_type, concat(' ', ?)), 0), 99999), ifnull(nullif(instr(Title_type, ?), 0), 99999), Title_type limit ?, ?",
			queryString, queryString, queryString, queryString, queryString, queryString, queryString, queryString, queryString, queryString, queryString, queryString, queryString, queryString, queryString, queryString, startRow, rowsTitlePerPage)
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

	var (
		queryPage int
		prevPage  int
		nextPage  int
		totalRows int
		nullScore sql.NullFloat64
		name      NameBasicsSQL
		names     []NameBasicsSQL
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

	row := db.QueryRow("select count(*) from title_basics where match (Title_type, Primary_title, Original_title, Genres) against (? in natural language mode)", queryString)
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

	rows, err := db.Query("select *, match(Primary_name) against (? in natural language mode) as rel_name, match(Primary_profession) against(? in natural language mode) as rel_profession from name_basics where match (Primary_name, Primary_profession) against (? in natural language mode) order by (rel_name*2)+(rel_profession) desc limit ?, ?",queryString, queryString, queryString, startId, rowsNamePerPage)

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
		err := rows.Scan(&name.Id, &name.NConst, &name.PrimaryName, &name.BirthYear, &name.DeathYear, &name.PrimaryProfession, &name.KnownForTitles, &name.CreateDate, &name.LastUpdated, &nullScore, &nullScore)
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

	if len(names) == 0 && FuzzyBrutalForceSearch {
		// Fulltext search failed
		// Brute-force method

		row := db.QueryRow("select count(*) from title_basics	where Primary_title like concat('%', ?, '%') or  Original_title like concat('%', ?, '%') or  Genres like concat('%', ?, '%') or Title_type like concat('%', ?, '%')", queryString, queryString, queryString, queryString)
		row.Scan(&totalRows)
		log.Println(totalRows)

		totalPages = int(math.Ceil(float64(totalRows) / rowsTitlePerPage))

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

		rows, err := db.Query("select * from name_basics where Primary_name like concat('%', ?, '%') or  Primary_profession like concat('%', ?, '%') order by Primary_name like concat(?, '%') desc, ifnull(nullif(instr(Primary_name, concat(' ', ?)), 0), 99999), ifnull(nullif(instr(Primary_name, ?), 0), 99999), Primary_name, Primary_profession like concat(?, '%') desc, ifnull(nullif(instr(Primary_profession, concat(' ', ?)), 0), 99999), ifnull(nullif(instr(Primary_profession, ?), 0), 99999), Primary_profession limit ?, ?", queryString, queryString, queryString, queryString, queryString, queryString, queryString, queryString, startRow, rowsTitlePerPage)
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
			//log.Print(title)
		}

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
