// JD Waldron and Gian Garnica
// Program to create and output a flight listing and reserve seats

package main

type Flight struct {
	flightNum       string
	originCity      string
	destCity        string
	aircraftType    string
	maxSeats        int
	currentResSeats int
}
