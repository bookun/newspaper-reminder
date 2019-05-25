package cmd

import (
	"net/http"

	newspaperreminder "github.com/bookun/newspaper-reminder"
)

func main() {
	http.HandleFunc("/remind", newspaperreminder.Remind)
	http.ListenAndServe(":8080", nil)
}
