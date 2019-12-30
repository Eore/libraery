package http

import (
	"fmt"
	"net/http"
)

var handler http.Handler = mux

//Middleware fungsi untuk memasukkan fungsi - fungsi sebagai middleware
func Middleware(middleware []func(http.Handler) http.Handler) {
	for _, val := range middleware {
		handler = val(handler)
	}
}

//StartServer fungsi untuk memulai server
func StartServer(port int) {
	fmt.Printf("Server http berjalan di port %d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
}
