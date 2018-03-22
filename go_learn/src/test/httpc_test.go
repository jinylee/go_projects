package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHttpCall(t *testing.T) {

	res, err := http.Get("http://google.com/search?q=news")

	if err == nil {
		fmt.Println(res.StatusCode, err)
		defer res.Body.Close()

		if res.StatusCode == http.StatusOK {
			bodyBytes, err2 := ioutil.ReadAll(res.Body)
			if err2 == nil {
				fmt.Println(string(bodyBytes))
			}
		}
	} else {
		fmt.Println("Error ")
		t.Fail()
	}
}
