package app

import "database/sql"

func ValidateTitleStruct(title *TitleBasics) TitleBasicsSQL {
	var titleSQL TitleBasicsSQL

	if title.OriginalTitle == "" {
		titleSQL.OriginalTitle = sql.NullString{
			String: "",
			Valid:  false,
		}
	} else {
		titleSQL.OriginalTitle = sql.NullString{
			String: title.OriginalTitle,
			Valid:  true,
		}
	}

	if title.StartYear == 0 {
		titleSQL.StartYear = sql.NullInt64{
			Int64: 0,
			Valid: false,
		}
	} else {
		titleSQL.StartYear = sql.NullInt64{
			Int64: int64(title.StartYear),
			Valid: true,
		}
	}

	if title.EndYear == 0 {
		titleSQL.EndYear = sql.NullInt64{
			Int64: 0,
			Valid: false,
		}
	} else {
		titleSQL.EndYear = sql.NullInt64{
			Int64: int64(title.EndYear),
			Valid: true,
		}
	}

	if title.RuntimeMinutes == 0 {
		titleSQL.RuntimeMinutes = sql.NullInt64{
			Int64: 0,
			Valid: false,
		}
	} else {
		titleSQL.RuntimeMinutes = sql.NullInt64{
			Int64: int64(title.RuntimeMinutes),
			Valid: true,
		}
	}

	if title.Genres == "" {
		titleSQL.Genres = sql.NullString{
			String: "",
			Valid:  false,
		}
	} else {
		titleSQL.Genres = sql.NullString{
			String: title.Genres,
			Valid:  true,
		}
	}

	return titleSQL
}
