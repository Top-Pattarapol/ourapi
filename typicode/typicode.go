package typicode

import (
	"encoding/json"
	"net/http"
)

type Decoder interface {
	Decode(result interface{}) error
}

type typicode struct {
	URL string
}

func (tc *typicode) Decode(result interface{}) error {
	resp, err := http.Get(tc.URL)
	if err != nil {
		return err
	}
	return json.NewDecoder(resp.Body).Decode(&result)
}

func Get(path string) *typicode {
	return &typicode{
		"https://jsonplaceholder.typicode.com" + path,
	}
}