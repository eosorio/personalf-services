package accounts

import (
	"database/sql"
	"git.osmon.local/personalf-services/databaseInfo"
	"errors"
	"fmt"
	"log"
)

// errors
var (
	// ErrRecordInvalid invalid record found
	ErrRecordInvalid = errors.New("Record is invalid")
)

// Account holds items for accounts info
type Account struct {
	ID                int64  `json:"id"`
	FinancialEntityID int64  `json:"financialEntityID"`
	Name              string `json:"name"`
	AccountTypeID     int64  `json:"accountTypeID"`
}

// GetAccounts returns info from all the accounts
func GetAccounts(accountID int64, accountType int64) ([]Account, error) {
	dbConnectInfo := databaseInfo.DBconnectInfo
	connectString := fmt.Sprintf("host=%s dbname=%s user=%s sslmode=disable", dbConnectInfo.Hostname, dbConnectInfo.Name, dbConnectInfo.User)
	query := "SELECT id, fin_entity_id, name, card_type FROM cards "
	if accountID != 0 || accountType != 0 {
		query = query + "WHERE "
	}
	if accountID != 0 {
		query = fmt.Sprintf("%s id=%d", query, accountID)
		if accountType != 0 {
			query = query + " AND "
		}
	}
	if accountType != 0 {
		query = fmt.Sprintf("%s id=%d ", query, accountType)
	}
	query = fmt.Sprintf("%s ORDER BY name ASC", query)
	db, err := sql.Open("postgres", connectString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer db.Close()
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	accountsList := make([]Account, 0)

	for rows.Next() {
		var accountItem Account
		if err := rows.Scan(&accountItem.ID, &accountItem.FinancialEntityID,
			&accountItem.Name,
			&accountItem.AccountTypeID); err != nil {
			log.Fatal(err)
			return nil, err
		}
		accountsList = append(accountsList, accountItem)
	}
	return accountsList, nil
}
