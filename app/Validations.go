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

func ValidateNameStruct(name *NameBasics) NameBasicsSQL {
	var nameSQL NameBasicsSQL

	if name.BirthYear == 0 {
		nameSQL.BirthYear = sql.NullInt64{
			Int64: 0,
			Valid: false,
		}
	} else {
		nameSQL.BirthYear = sql.NullInt64{
			Int64: int64(name.BirthYear),
			Valid: true,
		}
	}

	if name.DeathYear == 0 {
		nameSQL.DeathYear = sql.NullInt64{
			Int64: 0,
			Valid: false,
		}
	} else {
		nameSQL.DeathYear = sql.NullInt64{
			Int64: int64(name.DeathYear),
			Valid: true,
		}
	}

	if name.PrimaryProfession == "" {
		nameSQL.PrimaryProfession = sql.NullString{
			String: "",
			Valid:  false,
		}
	} else {
		nameSQL.PrimaryProfession = sql.NullString{
			String: name.PrimaryProfession,
			Valid:  true,
		}
	}

	if name.KnownForTitles == "" {
		nameSQL.KnownForTitles = sql.NullString{
			String: "",
			Valid:  false,
		}
	} else {
		nameSQL.KnownForTitles = sql.NullString{
			String: name.KnownForTitles,
			Valid:  true,
		}
	}

	return nameSQL
}

func ValidateTitlePrincipalStruct(titlePrincipal *TitlePrincipal) TitlePrincipalSQL {
	var titlePrincipalSQL TitlePrincipalSQL

	if titlePrincipal.Category == "" {
		titlePrincipalSQL.Category = sql.NullString{
			String: "",
			Valid:  false,
		}
	} else {
		titlePrincipalSQL.Category = sql.NullString{
			String: titlePrincipal.Category,
			Valid:  true,
		}
	}

	if titlePrincipal.Job == "" {
		titlePrincipalSQL.Job = sql.NullString{
			String: "",
			Valid:  false,
		}
	} else {
		titlePrincipalSQL.Job = sql.NullString{
			String: titlePrincipal.Job,
			Valid:  true,
		}
	}

	if titlePrincipal.Characters == "" {
		titlePrincipalSQL.Characters = sql.NullString{
			String: "",
			Valid:  false,
		}
	} else {
		titlePrincipalSQL.Characters = sql.NullString{
			String: titlePrincipal.Characters,
			Valid:  true,
		}
	}

	return titlePrincipalSQL
}
