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
	fmt.Println("You have", fuelRemaining, "litres of fuel left!")
}

// Reports fuel requirement for each planet
func calculateFuel(planet string) int {
	var fuel int

	switch planet {
	case "Mercury", "mercury":
		fuel = fuelMercury
	case "Venus", "venus":
		fuel = fuelVenus
	case "Mars", "mars":
		fuel = fuelMars
	case "Juipter", "jupiter":
		fuel = fuelJupiter
	case "Saturn", "saturn":
		fuel = fuelSaturn
	case "Uranus", "uranus":
		fuel = fuelUranus
	case "Neptune", "neptune":
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
	case "Mercury", "mercury":
		fmt.Printf("We need %d litres of fuel to get there.\n", fuelMercury)
	case "Venus", "venus":
		fmt.Printf("We need %d litres of fuel to get there.\n", fuelVenus)
	case "Earth", "earth":
		fmt.Println("We are already on Earth! Try again.")
		planetInput()
	case "Mars", "mars":
		fmt.Printf("We need %d litres of fuel to get there.\n", fuelMars)
	case "Jupiter", "jupiter":
		fmt.Printf("We need %d litres of fuel to get there.\n", fuelJupiter)
	case "Saturn", "saturn":
		fmt.Printf("We need %d litres of fuel to get there.\n", fuelSaturn)
	case "Uranus", "uranus":
		fmt.Printf("We need %d litres of fuel to get there.\n", fuelUranus)
	case "Neptune", "neptune":
		fmt.Printf("We need %d litres of fuel to get there.\n", fuelNeptune)
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
