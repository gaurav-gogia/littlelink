package model

// Constants for the server config
const (
	LISTENPORT = "8080"
)

// Constants for server routes
const (
	INDEX        = "/"
	CREATELITTLE = "/setSmall"
	GETBIG       = "/getLong"
)

// DBSTRING constant is used to define location of database
const DBSTRING = "./data"

// BUCKETSTRING constant is used to define the bucket for k:v db
const BUCKETSTRING = "linkdata"

// LOGSTRING constant is used to define location of application logs
const LOGSTRING = "./data/log"

// LOGLINE is the constant used to define the long line for log files for formatting
const LOGLINE = "------------------------------------------------------------"

// Severity levels for log
const (
	INFO = iota
	LOW
	MEDIUM
	HIGH
)
