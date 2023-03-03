package main

import (
	"fmt"
)

var planetChoice string
var fuelMercury int = 700
var fuelVenus int = 300
var fuelMars int = 500
var fuelJupiter int = 1200
var fuelSaturn int = 1500
var fuelUranus int = 2000
var fuelNeptune int = 2500

// Printer remaining fuel to terminal
func fuelGuage(fuelRemaining int) {
	fmt.Println("You have", fuelRemaining, "gallons of fuel left!")
}

// Reports fuel requirement for each planet
func calculateFuel(planet string) int {
	var fuel int

	switch planet {
	case "Mercury":
		fuel = fuelMercury
	case "Venus":
		fuel = fuelVenus
	case "Mars":
		fuel = fuelMars
	case "Juipter":
		fuel = fuelJupiter
	case "Saturn":
		fuel = fuelSaturn
	case "Uranus":
		fuel = fuelUranus
	case "Neptune":
		fuel = fuelNeptune
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
func flyToPlanet(planet string, fuel int) {
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
}

// Collects user planet input and compares to list. If not on list the function will restart.
func planetInput() {
	fmt.Println("Where would you like to go?")
	fmt.Scan(&planetChoice)

	switch planetChoice {
	case "Mercury":
		fmt.Printf("We need %d of fuel to get there.\n", fuelMercury)
	case "Venus":
		fmt.Printf("We need %d of fuel to get there.\n", fuelVenus)
	case "Earth":
		fmt.Println("We are already on Earth! Try again.")
		planetInput()
	case "Mars":
		fmt.Printf("We need %d of fuel to get there.\n", fuelMars)
	case "Jupiter":
		fmt.Printf("We need %d of fuel to get there.\n", fuelJupiter)
	case "Saturn":
		fmt.Printf("We need %d of fuel to get there.\n", fuelSaturn)
	case "Uranus":
		fmt.Printf("We need %d of fuel to get there.\n", fuelUranus)
	case "Neptune":
		fmt.Printf("We need %d of fuel to get there.\n", fuelNeptune)
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
