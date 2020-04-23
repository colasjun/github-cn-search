package engine

import "net/http"

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
