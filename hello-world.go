package api

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

// HelloHTTP is an HTTP Cloud Function with a request parameter.
func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Message string `json:"message"`
	}
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