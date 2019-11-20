package handlers

import (
	"net/http"
	"git.osmon.local/personalf-services/statementEntry"
)

func statementsGetCycle(w http.ResponseWriter, _ *http.Request, accountID int64, cycle int64) {
	statement, err := statementEntry.GetAllCycle(accountID, cycle)
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	if len(statement) == 0 {
		postError(w, http.StatusNotFound)
	} else {
		postBodyResponse(w, http.StatusOK, jsonResponse{"statementEntries": statement})
	}
}
