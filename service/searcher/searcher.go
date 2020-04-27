package searcher

import (
	"fmt"
	"github-cn-search/service/common"
	"github-cn-search/service/config"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)


type SearchData struct {
	SearchItems []SearchItem `json:"searchItems"`
	PageData common.PageDataStruct `json:"pageData"`
}

type SearchItem struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Stars int `json:"stars"`
	Labels []string `json:"labels"`
	Language string `json:"language"`
}

func Search(searchMsg string) (result SearchData,e error){
	SearchAddr := config.CONF.Get("GITHUB_HOST") + url.QueryEscape(strings.TrimSpace(searchMsg))
	fmt.Println("Search get addr=", SearchAddr)

	resp, err := http.Get(SearchAddr)
	if err != nil {
		fmt.Println("Search error...err=",err)
		return result,err
	}

	defer resp.Body.Close()

	document, e := goquery.NewDocumentFromReader(resp.Body)
	if e != nil {
		fmt.Println("Search error...err=",err)
		return result,err
	}

	// get search sata
	var searchData SearchData
	var searchItem SearchItem

	// find total
	document.Find(".codesearch-results .pb-3").Each(func(i int, s *goquery.Selection) {
		content := strings.TrimSpace(s.Find("h3").Text())
		totalIndex := strings.Index(content, " ")
		total := strings.ReplaceAll(content[0:totalIndex], ",", "")
		totalNum,_ := strconv.Atoi(total)

		searchData.PageData.Total = totalNum
		searchData.PageData.PageSize = 10

		CurrentPage,_ := strconv.Atoi(strings.TrimSpace(s.Find(".paginate-container .pagination .current").Text()))
		searchData.PageData.CurrentPage = CurrentPage

		searchData.PageData.TotalPage = 100

		fmt.Println("search page Data:", searchData)
	})

	// find item data
	document.Find(".repo-list .repo-list-item").Each(func(i int, s *goquery.Selection) {
		searchItem.Name = strings.TrimSpace(s.Find(".f4 a").Text())
		searchItem.Description = strings.TrimSpace(s.Find("p").Text())

		stars := strings.TrimSpace(s.Find(".mt-n1 .text-small .mr-3 a").Text())
		if strings.Index(stars, "k") == -1 {
			searchItem.Stars,_ = strconv.Atoi(stars)
		} else {
			newStars,_ := strconv.ParseFloat(strings.Replace(stars,"k","", -1), 64)
			searchItem.Stars =  int(newStars * 1000)
		}

		searchItem.Language = strings.TrimSpace(s.Find(".mt-n1 .text-small .mr-3 span[itemprop]").Text())

		searchItem.Labels = nil
		s.Find(".mt-n1 .topic-tag").Each(func(i int, s *goquery.Selection) {
			searchItem.Labels = append(searchItem.Labels, strings.TrimSpace(s.Text()))
		})

		searchData.SearchItems = append(searchData.SearchItems, searchItem)
	})

	return searchData,nil
}
