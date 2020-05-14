package api

import (
	"encoding/json"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{Route{
	"SayHello",
	"GET",
	"/hello",
	func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		dict := map[string]interface{}{
			"data":     "hello world!",
			"per_page": 100,
			"page":     1,
		}
		data, _ := json.Marshal(dict)
		_, _ = writer.Write(data)
	},
}}
