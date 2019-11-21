package statementEntry

import (
	"fmt"
	"git.osmon.local/personalf-services/handlers"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/statements", handlers.StatementsRouter)
	http.HandleFunc("/", handlers.RootHandler)
	err := http.ListenAndServe("0.0.0.0:11111", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
