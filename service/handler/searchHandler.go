package handler

import (
	"encoding/json"
	"fmt"
	"github-cn-search/service/common"
	"github-cn-search/service/condition"
	"github-cn-search/service/result"
	"github-cn-search/service/searcher"
	"net/http"
)

type ReturnSearchData struct {
	Code int `json:"code""`
	Data searcher.SearchData
}

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
	searchResult,e := searcher.Search(s)
	if e != nil {
		fmt.Println("SearchIndex searcher fail...err=",e)
		failResult,_ := json.Marshal(result.FailReturn("Search Get net error"))
		fmt.Fprintf(w, string(failResult))
		return nil
	}

	var returnData ReturnSearchData
	returnData.Code = common.Code.OK
	returnData.Data = searchResult

	successResult,_ := json.Marshal(returnData)
	fmt.Println("search success..return data ===>",searchResult)
	fmt.Fprintf(w, string(successResult))
	return nil
}