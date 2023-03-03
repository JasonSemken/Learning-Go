package main

import (
	"fmt"
)

var planetChoice string

// Printers remaining fuel to terminal
func fuelGuage(fuelRemaining int) {
	fmt.Println("You have", fuelRemaining, "gallons of fuel left!")
}

// Reports fuel requirement for each planet
func calculateFuel(planet string) int {
	var fuel int

	switch planet {
	case "Mercury":
		fuel = 500000
	case "Venus":
		fuel = 300000
	case "Mars":
		fuel = 700000
	}
	return fuel
}

// Prints Planet greating to terminal
func greetPlanet() {
	fmt.Println("Welcome to planet", planetChoice)
}

// Prints fuel warning to terminal
func cantFly() {
	fmt.Println("We do not have the available fuel to fly there.")
}

// main fucntion for calculating if you can fly to a planet
func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining, fuelCost int
	fuelRemaining = fuel

	fuelCost = calculateFuel(planet)

	if fuelRemaining >= fuelCost {
		greetPlanet()
		fuelRemaining = fuelRemaining - fuelCost
	} else {
		cantFly()
	}
	fuelGuage(fuelRemaining)
	return fuelRemaining

}

// Collects user planet input and compares to list. If not on list the function will restart.
func planetInput() {
	fmt.Println("Where would you like to go? Mars, Venus, or Mercury.")
	fmt.Scan(&planetChoice)

	switch planetChoice {
	case "Mercury":
		calculateFuel(planetChoice)
	case "Venus":
		calculateFuel(planetChoice)
	case "Mars":
		calculateFuel(planetChoice)
	default:
		fmt.Println("That's not a planet!! Try again.")
		planetInput()

	}

}

func main() {

	var fuelInput int

	planetInput()

	fmt.Println("How much fuel do we have?")
	fmt.Scan(&fuelInput)

	flyToPlanet(planetChoice, fuelInput)

}
