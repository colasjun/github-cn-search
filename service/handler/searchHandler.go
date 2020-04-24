package handler

import (
	"encoding/json"
	"fmt"
	"github-cn-search/service/common"
	condition_shop "github-cn-search/service/condition-shop"
	result_expoter "github-cn-search/service/result-expoter"
	"net/http"
)

func SearchIndex(w http.ResponseWriter, r *http.Request) (e error) {
	r.ParseForm()
	fmt.Printf("SearchIndex recieve msg success\n",r.Form)

	var msg string
	for k,v := range r.Form {
		if k == "s" {
			msg = v[0]
		}
	}

	s,c := condition_shop.Condition(msg)
	if c == common.Code.FAIL {
		fmt.Println("SearchIndex use condition deal msg fail...err=",c)
		failResult,_ := json.Marshal(result_expoter.FailReturn(s))
		fmt.Fprintf(w, string(failResult))
		return nil
	}

	return nil
}