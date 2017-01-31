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
	check(err)
	body, err := ioutil.ReadAll(rawResponse.Body)
	check(err)
	json.Unmarshal(body, &res)

	return res
}
