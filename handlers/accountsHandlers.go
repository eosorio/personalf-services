package handlers

import (
	"accounts"
	"net/http"
)

func accountsGetInfo(w http.ResponseWriter, _ *http.Request, accountID int64, accountType int64) {
	accountsInfo, err := accounts.GetAccounts(accountID, accountType)
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	if len(accountsInfo) == 0 {
		postError(w, http.StatusNotFound)
	} else {
		postBodyResponse(w, http.StatusOK, jsonResponse{"accounts": accountsInfo})
	}
}
