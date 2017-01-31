package main

import (
	"os"

	"github.com/mhutter/sbb"
)

func main() {
	query := sbb.ParseQuery(os.Args)
	res := sbb.FetchConnections(query)

	sbb.PrintConnectionList(res.Connections)
}
