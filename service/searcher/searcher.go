package searcher

import (
	"fmt"
	"github-cn-search/service/config"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func Search(searchMsg string) (result string,e error){
	SearchAddr := config.CONF.Get("GITHUB_HOST") + url.QueryEscape(strings.TrimSpace(searchMsg))
	fmt.Println("Search get addr=", SearchAddr)

	resp, err := http.Get(SearchAddr)
	if err != nil {
		fmt.Println("Search error...err=",err)
		return "Search Get net error",err
	}

	defer resp.Body.Close()

	bytes, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		fmt.Println("Search error...err=",err)
		return "Search Read Body error",err
	}

	return string(bytes),nil
}
