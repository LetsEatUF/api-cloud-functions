package api

import (
	"bytes"
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func Yay(w http.ResponseWriter, r *http.Request) {
	var u struct {
		Group string `json:"group"`
		Restaurant  string `json:"restaurant"`
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

	groups := client.Collection("groups").Doc(u.Group)
	doc, _ := groups.Get(ctx)
	m := doc.Data()

	var b bytes.Buffer	// this is a disgusting byte hack but what can you do its a hackathon
	fmt.Fprint(&b, m["Rests"])
	line, _ := b.ReadString(0)
	match := strings.Contains(line, u.Restaurant)
	ret := map[string]bool{"match":match}
	js,_:=json.Marshal(ret)
	w.Write(js)
}