package databaseInfo

import (
	"log"
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
	} else {
		// Try reading password from file if environment variable is missing
		passwordFilePath := "/run/secrets/db_password"
		passwordData, err := os.ReadFile(passwordFilePath)
		if err == nil {
			dbConnectInfo.Pass = string(passwordData)
		} else {
			log.Printf("Warning: Could not read password file %s: %v", passwordFilePath, err)
		}
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
