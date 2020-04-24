package handler

import (
	"fmt"
	condition_shop "github-cn-search/service/condition-shop"
	"net/http"
)

func SearchIndex(w http.ResponseWriter, r *http.Request) (e error) {
	r.ParseForm()
	fmt.Printf("searchEngine recieve msg success\n",r.Form)

	var msg string
	for k,v := range r.Form {
		if k == "s" {
			msg = v[0]
		}
	}

	condition_shop.Condition(msg)

	fmt.Fprintf(w, "ok")

	return nil
}