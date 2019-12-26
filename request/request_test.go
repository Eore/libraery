package request_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/Eore/libraery/request"
)

const baseURL = "https://reqres.in/api"

type response struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type testGET struct {
	Status   int
	URL      string
	Response string
}

var listTestGET []testGET = []testGET{
	{200, baseURL + "/users/1", `{"data":{"id":1,"email":"george.bluth@reqres.in","first_name":"George","last_name":"Bluth","avatar":"https://s3.amazonaws.com/uifaces/faces/twitter/calebogden/128.jpg"}}`},
	{200, baseURL + "/users/2", `{"data":{"id":2,"email":"janet.weaver@reqres.in","first_name":"Janet","last_name":"Weaver","avatar":"https://s3.amazonaws.com/uifaces/faces/twitter/josephstein/128.jpg"}}`},
	{200, baseURL + "/users/3", `{"data":{"id":3,"email":"emma.wong@reqres.in","first_name":"Emma","last_name":"Wong","avatar":"https://s3.amazonaws.com/uifaces/faces/twitter/olegpogodaev/128.jpg"}}`},
	{200, baseURL + "/users/4", `{"data":{"id":4,"email":"eve.holt@reqres.in","first_name":"Eve","last_name":"Holt","avatar":"https://s3.amazonaws.com/uifaces/faces/twitter/marcoramires/128.jpg"}}`},
	{200, baseURL + "/users/5", `{"data":{"id":5,"email":"charles.morris@reqres.in","first_name":"Charles","last_name":"Morris","avatar":"https://s3.amazonaws.com/uifaces/faces/twitter/stephenmoon/128.jpg"}}`},
	{404, baseURL + "/users/23", `{}`},
	{200, baseURL + "/unknown/1", `{"data":{"id":1,"name":"cerulean","year":2000,"color":"#98B2D1","pantone_value":"15-4020"}}`},
}

func TestGET(t *testing.T) {
	for i, value := range listTestGET {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			res := request.GET(value.URL, nil)
			b, _ := ioutil.ReadAll(res.Body)
			if res.StatusCode != value.Status || string(b) != value.Response {
				t.Fatalf("\nexpected :\n\tstatus -> %d\n\tresponse -> %s\ngot :\n\tstatus -> %d\n\tresponse -> %s", value.Status, value.Response, res.StatusCode, string(b))
			}
		})
	}
}

type testPOST struct {
	Status int
	URL    string
	Body   string
}

var listTestPOST []testPOST = []testPOST{
	{201, baseURL + "/users", `{"miaw", "lala"}`},
	{201, baseURL + "/users", `{"miao", "lili"}`},
	{201, baseURL + "/users", `{"meow", "lulu"}`},
	{201, baseURL + "/users", `{"miau", "lele"}`},
	{201, baseURL + "/users", `{"meou", "lolo"}`},
}

func TestPOST(t *testing.T) {
	for i, value := range listTestPOST {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			res := request.POST(value.URL, value.Body, nil)
			if res.StatusCode != value.Status {
				t.Fatalf("\nexpected :\n\tstatus -> %d\ngot :\n\tstatus -> %d", value.Status, res.StatusCode)
			}
		})
	}
}

type testPUT struct {
	Status int
	URL    string
	Body   string
}

var listTestPUT []testPUT = []testPUT{
	{200, baseURL + "/users/1", `{"miaw", "lala"}`},
	{200, baseURL + "/users/2", `{"miao", "lili"}`},
	{200, baseURL + "/users/3", `{"meow", "lulu"}`},
	{200, baseURL + "/users/4", `{"miau", "lele"}`},
	{200, baseURL + "/users/5", `{"meou", "lolo"}`},
}

func TestPUT(t *testing.T) {
	for i, value := range listTestPUT {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			res := request.PUT(value.URL, value.Body, nil)
			if res.StatusCode != value.Status {
				t.Fatalf("\nexpected :\n\tstatus -> %d\ngot :\n\tstatus -> %d", value.Status, res.StatusCode)
			}
		})
	}
}

type testDELETE struct {
	Status int
	URL    string
}

var listTestDELETE []testDELETE = []testDELETE{
	{204, baseURL + "/users/1"},
	{204, baseURL + "/users/2"},
	{204, baseURL + "/users/3"},
	{204, baseURL + "/users/4"},
	{204, baseURL + "/users/5"},
}

func TestDELETE(t *testing.T) {
	for i, value := range listTestDELETE {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			res := request.DELETE(value.URL, nil)
			if res.StatusCode != value.Status {
				t.Fatalf("\nexpected :\n\tstatus -> %d\ngot :\n\tstatus -> %d", value.Status, res.StatusCode)
			}
		})
	}
}
