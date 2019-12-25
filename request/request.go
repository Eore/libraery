package request

import (
	"bytes"
	"log"
	"net/http"
	"time"
)

//request fungsi untuk pemanggilan http request
func request(method, url, data string, header map[string]string) http.Response {
	req, err := http.NewRequest(method, url, bytes.NewBufferString(data))
	if err != nil {
		log.Fatalln("inisiasi request gagal")
	}
	for key, value := range header {
		req.Header.Set(key, value)
	}
	client := http.Client{Timeout: time.Second * 10}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln("pemanggilan request gagal")
	}
	return *res
}

//GET fungsi untuk pemanggilan method GET ke url tertentu
func GET(url string, header map[string]string) http.Response {
	return request("GET", url, "", header)
}

//POST fungsi untuk pemanggilan method POST ke url tertentu dengan data tertentu
func POST(url, data string, header map[string]string) http.Response {
	return request("POST", url, data, header)
}

//PUT fungsi untuk pemanggilan method PUT ke url tertentu dengan data tertentu
func PUT(url, data string, header map[string]string) http.Response {
	return request("PUT", url, data, header)
}

//DELETE fungsi untuk pemanggilan method DELETE ke url tertentu dengan data tertentu
func DELETE(url string, header map[string]string) http.Response {
	return request("DELETE", url, "", header)
}
