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
		fuel = 700
	case "Venus":
		fuel = 300
	case "Mars":
		fuel = 500
	case "Juipter":
		fuel = 1200
	case "Saturn":
		fuel = 1500
	case "Uranus":
		fuel = 2000
	case "Neptune":
		fuel = 2500
	}
	return fuel
}

// Prints planet greating to terminal
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
	fmt.Println("Where would you like to go?")
	fmt.Scan(&planetChoice)

	switch planetChoice {
	case "Mercury":
		fmt.Println("We need 700 units of fuel to get there")
	case "Venus":
		calculateFuel(planetChoice)
	case "Earth":
		fmt.Println("We are already on Earth! Try again.")
		planetInput()
	case "Mars":
		calculateFuel(planetChoice)
	case "Jupiter":
		calculateFuel(planetChoice)
	case "Saturn":
		calculateFuel(planetChoice)
	case "Uranus":
		calculateFuel(planetChoice)
	case "Neptune":
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
