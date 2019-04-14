package simple

import (
	"fmt"
	"net/http"
	"testing"

	"gopkg.in/h2non/baloo.v3"
)

var test = baloo.New("http://httpbin.org")

func assert(res *http.Response, req *http.Request) error {
	fmt.Println(req.Header)
	return nil
}

func init() {
	// Register assertion function at global level
	baloo.AddAssertFunc("test", assert)
}

func TestSimpleGet(t *testing.T) {
	const schema = `{"$ref": "file:///C:/Users/User/go/src/github.com/Tempeny/api_test/httpbin.org/get.json#/get"}`
	test.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3760.0 Safari/537.36 Edg/75.0.126.").
		Get("/get").
		Expect(t).
		Status(200).
		Assert("test").
		Header("Server", "nginx").
		Header("Content-Type", "application/json").
		HeaderPresent("Date").
		Type("json").
		JSONSchema(schema).
		Done()
}
