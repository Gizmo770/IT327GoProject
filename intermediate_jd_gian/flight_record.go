// JD Waldron and Gian Garnica
// Program to create and output a flight listing and reserve seats
// Uses random values for user input since Go Playground doesn't support user input

package main

import (
	"fmt"
	"math/rand"
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

func (f *Flight) getSeatsAvailable() int {
	return f.maxSeats - f.currentResSeats
}

func (f Flight) String() string {
	return fmt.Sprintf("%-10s %-10s %-10s %-3d seats available", f.flightNum, f.originCity,
		f.destCity, f.getSeatsAvailable())
}

func (f *Flight) reserve(seats int) bool {
	if seats <= f.getSeatsAvailable() {
		f.currentResSeats += seats
		return true
	} else {
		return false
	}
}

func (f *Flight) calcMaximumSeats() {
	if f.aircraftType == "737-800" {
		f.maxSeats = 160
	} else if f.aircraftType == "A321" {
		f.maxSeats = 196
	} else if f.aircraftType == "CRJ-700" {
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
	fmt.Print("Enter your choice: ")
	userChoice := "L"
	fmt.Println("Selection:", userChoice)
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
			userChoice = "R"
			fmt.Println("Selection:", userChoice)
			fmt.Println()
		} else if userChoice == "R" {
			reserveSeats()
			fmt.Println("Choose one of the following:")
			fmt.Println("  L - list available flights")
			fmt.Println("  R - reserve seats")
			fmt.Println("  Q - quit")
			userChoice = "Q"
			fmt.Println("Selection:", userChoice)
			fmt.Println()
		} else {
			fmt.Println("User input not recognized. Make sure you input capital letters.")
		}
	}
	fmt.Println("Goodbye!")
}

func addFlights() {
	flights[0] = Flight{flightNum: "BT7274", originCity: "Dallas", destCity: "Normal", aircraftType: "CRJ-700"}
	flights[1] = Flight{flightNum: "VS8156", originCity: "Fremont", destCity: "Portland", aircraftType: "A321"}
	flights[2] = Flight{flightNum: "FD5574", originCity: "Juneau", destCity: "Key West", aircraftType: "737-800"}
	flights[3] = Flight{flightNum: "GZ9601", originCity: "Sacremento", destCity: "Atlanta", aircraftType: "Cessna"}

	for i := 0; i < 4; i++ {
		flights[i].calcMaximumSeats()
	}
}

func listFlights() {
	for _, flight := range flights {
		if flight.getSeatsAvailable() > 0 {
			fmt.Println(flight)
		}
	}
	fmt.Println()
}

func reserveSeats() {
	userInput := rand.Intn(3) + 1
	switch userInput {
	case 1:
		fmt.Println("Reserving on BT7274")
		userSeats := rand.Intn(70) + 1
		fmt.Println("Attempting to reserve", userSeats, "seats")
		if flights[0].reserve(userSeats) {
			fmt.Println("Reservation successful.")
			fmt.Println()
			fmt.Println("Flights after reservation: ")
			listFlights()
		} else {
			fmt.Println("Sorry, not enough seats.")
			fmt.Println()
		}
	case 2:
		fmt.Println("Reserving on VS8156")
		userSeats := rand.Intn(200) + 1
		fmt.Println("Attempting to reserve", userSeats, "seats")
		if flights[1].reserve(userSeats) {
			fmt.Println("Reservation successful.")
			fmt.Println()
			fmt.Println("Flights after reservation: ")
			listFlights()
		} else {
			fmt.Println("Sorry, not enough seats.")
			fmt.Println()
		}
	case 3:
		fmt.Println("Reserving on FD5574")
		userSeats := rand.Intn(175) + 1
		fmt.Println("Attempting to reserve", userSeats, "seats")
		if flights[1].reserve(userSeats) {
			fmt.Println("Reservation successful.")
			fmt.Println()
			fmt.Println("Flights after reservation: ")
			listFlights()
		} else {
			fmt.Println("Sorry, not enough seats.")
			fmt.Println()
		}
	}
}
