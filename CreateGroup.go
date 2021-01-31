package api

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	var u struct {
		Username       string `json:"user"`
		UsernameFriend string `json:"friend"`
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

	groups := client.Collection("groups")
	groups.Doc(u.Username).Create(ctx, struct {
		Members []string
		Rests []string
	}{
		Members: []string{u.Username, u.UsernameFriend},
		Rests: []string{},
	})

	fmt.Fprint(w, html.EscapeString(u.Username))
}