package api

import (
	"context"
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"cloud.google.com/go/firestore"
)

type User struct {
	username    string  `firestore:"username"`
	email float64 `firestore:"email"`
	created string `firestore:"created"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var u struct {
		username string
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
	user := users.Doc("swXMXzuLhBQZjYTGdxH2")

	docsnap, err := user.Get(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	dataMap := docsnap.Data()
	fmt.Println(dataMap)

	fmt.Fprint(w, html.EscapeString(u.username))
}
