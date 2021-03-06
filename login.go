package api

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var u struct {
		Username string  `json:"username"`
	}

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		fmt.Fprint(w, "Lets Eat says: Hello, World!")
		return
	}

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "lets-eat-303301")
	if err != nil {
		fmt.Println("err")
	}
	defer client.Close()

	users := client.Collection("users")
	users.Doc(u.Username).Create(ctx, map[string]string{"created": time.Now().String()})

	fmt.Fprint(w, html.EscapeString(u.Username))
}
