package sbb

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Response represents the root object of a OpenTransport API response
type Response struct {
	Connections []Connection
}

// FetchConnections queries the OpenTransport API
func FetchConnections(q *Query) (res *Response) {
	rawResponse, err := http.Get(q.URL())
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(body, &res)

	return res
}
