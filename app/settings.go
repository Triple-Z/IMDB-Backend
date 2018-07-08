package app

const (
	/********************** App Settings *****************************/
	/*
		The server serve port number.
	 */
	PortNumber = "2333"
	/*
		The `FuzzyBrutalForceSearch` option will improve the search result quality
		but extend the response time heavily.
	*/
	FuzzyBrutalForceSearch = true
	/*
		The number of titles per page should return.
	*/
	rowsTitlePerPage = 20
	/*
		The number of names per page should return.
	*/
	rowsNamePerPage = 50
	/*
		Cross-Origin Resource Sharing(CORS) support.
	*/
	AllowCORS = true
	/*
		Release mode or debug mode for Gin.
	*/
	ReleaseMode = false
	/********************** Database Settings *****************************/
	/*
		Database user name.
	*/
	DBUserName = "root"
	/*
		Database user password.
	*/
	DBUserPassword = ""
	/*
		Database remote IP address.
	*/
	DBAddress = "127.0.0.1"
	/*
		Database remote port number.
	*/
	DBPort = "3306"
	/*
		Use specific database name.
	*/
	UseDBName = "imdb"
)
