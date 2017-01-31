package sbb

import "fmt"

// PrintConnectionList prints out a list of connections
func PrintConnectionList(connections []Connection) {
	fmt.Println("---------------------------------------------------------")
	for _, conn := range connections {
		conn.Print()
		fmt.Println("---------------------------------------------------------")
	}
}

// Print prints out a single Connection
func (c *Connection) Print() {
	fmt.Printf("%-30s dep: %s     dur: %s\n", c.DisplayFrom(), c.DisplayDeparture(), c.DisplayDuration())
	fmt.Printf("%-30s arr: %s\n", c.DisplayTo(), c.DisplayArrival())
}
