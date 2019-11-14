package databaseInfo

// dbConnectData - Info for connecting to the database (host, user name, etc)
type DBconnectData struct {
	User     string
	Pass     string
	Hostname string
	IPhost   string
	Name     string
}

var (
	DBconnectInfo = DBconnectData{User: "eosorio", Hostname: "localhost", Name: "personal_finance_test"}
)
