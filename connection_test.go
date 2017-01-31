package sbb

import "testing"

var conn = &Connection{
	From: Checkpoint{
		Station:   Location{Name: "Asgard"},
		Departure: "2017-01-31T19:02:00+0100",
	},
	To: Checkpoint{
		Station: Location{Name: "Hel"},
		Arrival: "2017-01-31T21:11:00+0100",
	},
	Duration: "00d02:09:00",
}

func TestDisplayFrom(t *testing.T) {
	assert(t, "DisplayFrom()", "Asgard", conn.DisplayFrom())
}

func TestDisplayTo(t *testing.T) {
	assert(t, "DisplayTo()", "Hel", conn.DisplayTo())
}

func TestDisplayDeparture(t *testing.T) {
	assert(t, "DisplayDeparture()", "19:02", conn.DisplayDeparture())
}

func TestDisplayArrival(t *testing.T) {
	assert(t, "DisplayArrival()", "21:11", conn.DisplayArrival())
}

func TestDisplayDuration(t *testing.T) {
	assert(t, "DisplayDuration()", "02:09", conn.DisplayDuration())
}

func assert(t *testing.T, name, expected, actual string) {
	if actual != expected {
		t.Errorf("Expected %s to be '%s', got '%s'", name, expected, actual)
	}
}
