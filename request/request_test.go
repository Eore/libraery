package request_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/Eore/libraery/request"
)

const url = "https://jsonplaceholder.typicode.com"

type response struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

var testGET map[string]response = map[string]response{
	"/posts/1": response{
		UserID: 1,
		ID:     1,
		Title:  "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
		Body:   "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto",
	},
	"/posts/2": response{
		UserID: 1,
		ID:     2,
		Title:  "qui est esse",
		Body:   "est rerum tempore vitae\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\nqui aperiam non debitis possimus qui neque nisi nulla",
	},
}

func TestGET(t *testing.T) {
	for key, value := range testGET {
		t.Run(key, func(t *testing.T) {
			res := request.GET(url+key, nil)
			var rsp response
			json.NewDecoder(res.Body).Decode(&rsp)
			if rsp != value {
				t.Fatalf("\ndata yang di dapat :\n%+v \nharusnya :\n%+v", rsp, value)
			}
		})
	}
}

func TestPOST(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		res := request.POST(url+"/posts", `{"title":"foo","body":"bar","userId":1}`, nil)
		b, _ := ioutil.ReadAll(res.Body)
		t.Fatal(string(b))
	})
}
