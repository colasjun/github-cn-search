package handler

import (
	"encoding/json"
	"fmt"
	"github-cn-search/service/common"
	"github-cn-search/service/condition"
	"github-cn-search/service/result"
	"net/http"
)

func SearchIndex(w http.ResponseWriter, r *http.Request) (e error) {
	decoder := json.NewDecoder(r.Body)
	var params map[string]string
	decoder.Decode(&params)

	// param validate
	fmt.Println("SearchIndex receive msg success\n",params)
	if len(params) < 1 {
		failResult,_ := json.Marshal(result.FailReturn("params error"))
		fmt.Fprintf(w, string(failResult))
		return nil
	}

	// condition change
	s,c := condition.Condition(params)
	if c == common.Code.FAIL {
		fmt.Println("SearchIndex use condition deal msg fail...err=",c)
		failResult,_ := json.Marshal(result.FailReturn(s))
		fmt.Fprintf(w, string(failResult))
		return nil
	}

	// search result
	fmt.Fprintf(w, s)
	return nil
}