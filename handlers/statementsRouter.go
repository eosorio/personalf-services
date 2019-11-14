package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

// StatementsRouter handles the statements route
func StatementsRouter(w http.ResponseWriter, r *http.Request) {

	var (
		sAccountID []string
		sCycle     []string

		accountID int64
		cycle     int64
		err       error
	)

	fmt.Printf("StatementsRouter: r.URL.Path: '%v'\n", r.URL.Path) //igm
	fmt.Printf("StatementsRouter: r.URL: '%v'\n", r.URL)           //igm
	query := r.URL.Query()
	fmt.Printf("StatementsRouter: Query: '%v'\n", query)
	print("path is /statements\n") //igm
	for key, value := range query {
		switch key {
		case "accountId":
			sAccountID = value
		case "cycle":
			sCycle = value
		}
	}

	// Validate accountId and cycle. Mandatory arguments, integers only
	if len(sAccountID) == 0 || len(sCycle) == 0 {
		postError(w, http.StatusMethodNotAllowed)
		return
	} else if sAccountID[0] == "" || sCycle[0] == "" {
		postError(w, http.StatusMethodNotAllowed)
		return
	} else if accountID, err = strconv.ParseInt(sAccountID[0], 10, 64); err != nil {
		postError(w, http.StatusMethodNotAllowed)
		return
	} else if cycle, err = strconv.ParseInt(sCycle[0], 10, 64); err != nil {
		postError(w, http.StatusMethodNotAllowed)
		return
	}

	switch r.Method {
	case http.MethodGet:
		statementsGetCycle(w, r, accountID, cycle)
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
