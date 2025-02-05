package databaseInfo

import (
	"os"
)

// GetConfigFromEnv Get environment variables for connecting to the database
func GetConfigFromEnv() DBconnectData {
	var dbConnectInfo DBconnectData
	value, isSet := os.LookupEnv("DB_USER")
	if isSet {
		dbConnectInfo.User = value
	}
	value, isSet = os.LookupEnv("DB_PASSWORD")
	if isSet {
		dbConnectInfo.Pass = value
	}
	value, isSet = os.LookupEnv("DB_NAME")
	if isSet {
		dbConnectInfo.Name = value
	}
	value, isSet = os.LookupEnv("DB_HOST")
	if isSet {
		dbConnectInfo.Hostname = value
	}
	return dbConnectInfo
}
