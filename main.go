package main

import (
	"net/http"

	"github.com/bookun/newspaper-reminder/handler"
)

func main() {
	http.HandleFunc("/remind", handler.Remind)
	http.ListenAndServe(":8080", nil)
}
