// JD Waldron and Gian Garnica
// Program to create and output a flight listing and reserve seats

package main

import (
	"fmt"
)

var (
	flights [4]Flight
)

type Flight struct {
	flightNum       string
	originCity      string
	destCity        string
	aircraftType    string
	maxSeats        int
	currentResSeats int
}

func (f Flight) getSeatsAvailable() int {
	return f.maxSeats - f.currentResSeats
}

func (f Flight) String() string {
	return fmt.Sprintf("%-10s %-10s %-10s %-3d seats available", f.flightNum, f.originCity,
		f.destCity, f.getSeatsAvailable())
}

func (f Flight) reserve(seats int) bool {
	if seats <= f.getSeatsAvailable() {
		f.currentResSeats += seats
		return true
	} else {
		return false
	}
}

func (f *Flight) calcMaximumSeats(airType string) {
	if airType == "737-800" {
		f.maxSeats = 160
	} else if airType == "A321" {
		f.maxSeats = 196
	} else if airType == "CRJ-700" {
		f.maxSeats = 63
	} else {
		f.maxSeats = 0
	}

}

func main() {
	addFlights()
	fmt.Println("Welcome to Redbird Airlines!")
	fmt.Println("Choose one of the following:")
	fmt.Println("  L - list available flights")
	fmt.Println("  R - reserve seats")
	fmt.Println("  Q - quit")
	fmt.Println("Enter your choice: ")
	var userChoice string
	fmt.Scanln(userChoice)
	fmt.Println()

	// using break statements since Go doesn't have a "while" keyword
	for {
		if userChoice == "Q" {
			break
		} else if userChoice == "L" {
			listFlights()
			fmt.Println("Choose one of the following:")
			fmt.Println("  L - list available flights")
			fmt.Println("  R - reserve seats")
			fmt.Println("  Q - quit")
			fmt.Println("Enter your choice: ")
			fmt.Scanln(userChoice)
			fmt.Println()
		} else if userChoice == "R" {
			reserveSeats()
		}
	}
}

func addFlights() {
	flights[0] = Flight{flightNum: "BT7274", originCity: "Dallas", destCity: "Normal", aircraftType: "CRJ-700"}
	flights[1] = Flight{flightNum: "VS8156", originCity: "Fremont", destCity: "Portland", aircraftType: "A321"}
	flights[2] = Flight{flightNum: "FD5574", originCity: "Juneau", destCity: "Key West", aircraftType: "737-800"}
	flights[3] = Flight{flightNum: "GZ9601", originCity: "Sacremento", destCity: "Atlanta", aircraftType: "Cessna"}
}

func listFlights() {
	for _, element := range flights {
		if element.getSeatsAvailable() > 0 {
			fmt.Println(element)
		}
	}
}

func reserveSeats() {
	fmt.Println("On which flight?")
	var userInput string
	fmt.Scanln(userInput)
	switch userInput {
	case "BT7274":

	}
}
