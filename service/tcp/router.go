package tcp

import (
	"encoding/json"
	"fmt"
	condition_shop "github-cn-search/service/condition-shop"
	result_expoter "github-cn-search/service/result-expoter"
	"net/http"
)

func router()  {
	// for search
	http.HandleFunc("/search", func(writer http.ResponseWriter, request *http.Request) {
		searchEngine(writer, request)
	})

	// for menu
	http.HandleFunc("/menu", func(writer http.ResponseWriter, request *http.Request) {
		menuEngine(writer, request)
	})
}

func searchEngine(w http.ResponseWriter, r *http.Request) (e error) {
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

func menuEngine(w http.ResponseWriter, r *http.Request) (e error) {
	menu := condition_shop.Menu()

	fmt.Println("get menu result:", menu)
	bytes, e := json.Marshal(menu)
	if e != nil {
		fmt.Println("menu Engine parse json fail...err=",e)
		failResult,_ := json.Marshal(result_expoter.FailReturn("system error"))
		fmt.Fprintf(w, string(failResult))
		return nil
	}

	fmt.Fprintf(w, string(bytes))
	return nil
}