package sbb

import (
	"strings"
	"time"
)

const timeFormat = "2006-01-02T15:04:05-0700"

// RequiredFields from the transport APi
var RequiredFields = []string{
	"connections/from/station/name",
	"connections/from/departure",
	"connections/from/platform",
	"connections/to/station/name",
	"connections/to/arrival",
	"connections/duration",
}

// Connection represents a possible journey between two locations
type Connection struct {
	From     Checkpoint
	To       Checkpoint
	Duration string
}

// Checkpoint represents an arrival or a departure point (in time and space) of
// a connection
type Checkpoint struct {
	Station   Location
	Departure string
	Arrival   string
	Platform  string
}

// Location represents a Checkpoint on a Connection
type Location struct {
	Name string
}

// DisplayFrom returns the formatted starting Location
func (c *Connection) DisplayFrom() string {
	return c.From.Station.Name
}

// DisplayTo returns the formatted starting Location
func (c *Connection) DisplayTo() string {
	return c.To.Station.Name
}

// DisplayDeparture returns the formatted starting Location
func (c *Connection) DisplayDeparture() string {
	departure, err := time.Parse(timeFormat, c.From.Departure)
	if err != nil {
		panic(err)
	}
	return departure.Format("15:04")
}

// DisplayArrival returns the formatted starting Location
func (c *Connection) DisplayArrival() string {
	arrival, err := time.Parse(timeFormat, c.To.Arrival)
	if err != nil {
		panic(err)
	}
	return arrival.Format("15:04")
}

// DisplayDuration returns the formatted starting Location
func (c *Connection) DisplayDuration() string {
	return strings.TrimSuffix(strings.TrimPrefix(c.Duration, "00d"), ":00")
}
