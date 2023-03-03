package main

import (
	"fmt"
	"os"
)

// Create the function fuelGauge() here
func fuelGuage(fuelRemaining int) {
	fmt.Println("You have", fuelRemaining, "gallons of fuel left!")
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
	var fuel int

	// function configured for 'if else' or switch. Comment out either to test.

	//if planet == "Mercury" {
	//	fuel = 500000
	//} else if planet == "Venus" {
	//	fuel = 300000
	//} else if planet == "Mars" {
	//	fuel = 700000
	//} else {
	//	fmt.Println("That's not a planet!! Try again.")
	//	os.Exit(69)
	//}
	//return fuel

	switch planet {
	case "Mercury":
		fuel = 500000
	case "Venus":
		fuel = 300000
	case "Mars":
		fuel = 700000
	default:
		fmt.Println("That's not a planet!! Try again.")
		os.Exit(69)
	}
	return fuel
}

// Create the function greetPlanet() here
func greetPlanet(planet string) {
	fmt.Println("Welcome to planet", planet)
}

// Create the function cantFly() here
func cantFly() {
	fmt.Println("We do not have the available fuel to fly there.")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining, fuelCost int
	fuelRemaining = fuel

	fuelCost = calculateFuel(planet)

	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining = fuelRemaining - fuelCost
	} else {
		cantFly()
	}
	fuelGuage(fuelRemaining)
	return fuelRemaining

}

func main() {
	// Test your functions!

	// Create `planetChoice` and `fuel`
	var planetChoice string
	var fuelInput int

	fmt.Println("Where would you like to go? Mars, Venus, or Mercury.")
	fmt.Scan(&planetChoice)

	//fmt.Println("How much fuel do we have?")
	//fmt.Scan(&fuelInput)
	//fuelAmount := 1000000
	//planetChoice := "Mars"
	flyToPlanet(planetChoice, fuelInput)
	// And then liftoff!

}
