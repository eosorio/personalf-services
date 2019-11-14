package handlers

import (
	"net/http"
	"strconv"
)

// AccountsRouter handles the statements route
func AccountsRouter(w http.ResponseWriter, r *http.Request) {

	var (
		sAccountID   []string
		sAccountType []string

		accountID   int64
		accountType int64
		err         error
	)

	// Get parameters passed in the URL
	query := r.URL.Query()
	for key, value := range query {
		switch key {
		case "accountId":
			sAccountID = value
		case "accountType":
			sAccountType = value
		}
	}

	// Only one account ID is accepted
	if len(sAccountID) == 1 {
		if accountID, err = strconv.ParseInt(sAccountID[0], 10, 64); err != nil {
			postError(w, http.StatusMethodNotAllowed)
			return
		}
	}

	// Only one account Type ID is accepted
	if len(sAccountType) == 1 {
		if accountType, err = strconv.ParseInt(sAccountType[0], 10, 64); err != nil {
			postError(w, http.StatusMethodNotAllowed)
			return
		}
	}

	switch r.Method {
	case http.MethodGet:
		accountsGetInfo(w, r, accountID, accountType)
		return
	case http.MethodPost:
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
		return
	}

	//path = strings.TrimPrefix(path, "/statments/")
	// check here if path is on the statemnts ñist  if !(CycleStatmentExist(path) postError(w, http.StatusNotFound) rturn

	switch r.Method {
	case http.MethodPut:
	case http.MethodPatch:
	case http.MethodDelete:
		postError(w, http.StatusMethodNotAllowed)
		return
	}
}
