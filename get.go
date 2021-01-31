package api

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetGroup(w http.ResponseWriter, r *http.Request) {
	var u struct {
		Username string `json:"username"`
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

	groups := client.Collection("groups").Where("Members", "array-contains", u.Username)
	doc, err := groups.Documents(ctx).Next()
	data := make(map[string]string)
	if err != nil {
		data["group"] = "None"
		data["found"] = "false"
	} else {
		data["group"] = doc.Ref.ID
		data["found"] = "true"
	}

	js, _:=json.Marshal(data)
	w.Write(js)
}