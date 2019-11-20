package statementEntry

import (
	//"context"
	"database/sql"
	"git.osmon.local/personalf-services/databaseInfo"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// errors

var (
	// ErrRecordInvalid found invalid record
	ErrRecordInvalid = errors.New("Record is invalid")
)

// StatementEntry holds items for (bank) statements
type StatementEntry struct {
	ID          int     `json:"id"`
	Date        string  `json:"date"`
	Amount      float32 `json:"amount"`
	Balance     float32 `json:"balance"`
	CurrencyID  int     `json:"currencyID"`
	Description string  `json:"description"`
	AuthNumber  int     `json:"authNumber"`
}

// GetAllCycle returns all records for a specific cycle and account
func GetAllCycle(accountID int64, cycle int64) ([]StatementEntry, error) {
	//	var (
	//		ctx context.Context
	//	)

	dbConnectInfo := databaseInfo.DBconnectInfo
	connectString := fmt.Sprintf("host=%s dbname=%s user=%s sslmode=disable", dbConnectInfo.Hostname, dbConnectInfo.Name, dbConnectInfo.User)
	query := fmt.Sprintf("SELECT id, s_date, description, amount, currency_id, auth_number FROM bank_statements WHERE account_id=%d AND s_cycle=%d", accountID, cycle)
	db, err := sql.Open("postgres", connectString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer db.Close()
	//fmt.Printf("Executing query: '%s'\n", query)  //igm
	//rows, err := db.QueryContext(ctx, query)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//fmt.Printf("Finished query\n")   //igm
	defer rows.Close()
	statementRows := make([]StatementEntry, 0)

	for rows.Next() {
		var statementRecord StatementEntry
		if err := rows.Scan(&statementRecord.ID, &statementRecord.Date, &statementRecord.Description,
			&statementRecord.Amount,
			&statementRecord.CurrencyID, &statementRecord.AuthNumber); err != nil {
			log.Fatal(err)
			return nil, err
		}
		statementRows = append(statementRows, statementRecord)
	}
	return statementRows, nil
}

// ValidateEntry validates entry is valid
func ValidateEntry(entry StatementEntry) error {
	if (entry.Date == "") || (entry.Description == "") {
		return ErrRecordInvalid
	}
	return nil
}
