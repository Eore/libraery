package http

import (
	"net/http"
	"strings"
)

//Method struct untuk handler di route
type Method struct {
	GET     func(http.ResponseWriter, *http.Request)
	POST    func(http.ResponseWriter, *http.Request)
	PUT     func(http.ResponseWriter, *http.Request)
	DELETE  func(http.ResponseWriter, *http.Request)
	OPTIONS func(http.ResponseWriter, *http.Request)
}

//mux adalah variable global untuk routing
var mux *http.ServeMux = http.DefaultServeMux

//Route fungsi untuk pembuatan rute api dengan method yang ada
func Route(url string, method Method) {
	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		switch strings.ToUpper(r.Method) {
		case "GET":
			method.GET(w, r)
		case "POST":
			method.POST(w, r)
		case "PUT":
			method.PUT(w, r)
		case "DELETE":
			method.DELETE(w, r)
		case "OPTIONS":
			method.OPTIONS(w, r)
		}
	})
}
