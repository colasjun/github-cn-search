package tcp

import (
	"github-cn-search/service/handler"
	"net/http"
)

func router()  {
	// for search
	http.HandleFunc("/search", func(writer http.ResponseWriter, request *http.Request) {
		handler.SearchIndex(writer, request)
	})

	// for menu
	http.HandleFunc("/menu", func(writer http.ResponseWriter, request *http.Request) {
		handler.MenuIndex(writer, request)
	})
}