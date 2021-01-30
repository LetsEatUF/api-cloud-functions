package api

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

// HelloWorld is an HTTP Cloud Function with a request parameter.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Message string `json:"message"`
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "Lets Eat says: Hello, World!")
		return
	}
	if d.Message == "" {
		fmt.Fprint(w, "Lets Eat says: Hello, World!")
		return
	}
	fmt.Fprintf(w, "Lets Eat says: %s!", html.EscapeString(d.Message))

}