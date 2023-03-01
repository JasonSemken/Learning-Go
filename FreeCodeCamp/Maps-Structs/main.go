package main

import (
	"fmt"
	"reflect"
)

//	Below block in a MAP
//func main() {
//	statePopulations := make(map[string]int)
//	statePopulations = map[string]int{
//		"California":   39250017,
//		"Texas":        27862596,
//		"Florida":      20612439,
//		"New York":     19745289,
//		"Pennsylvania": 12802503,
//		"Illinois":     12801539,
//		"Ohio":         11614373,
//	}
//
//	fmt.Println(statePopulations)
//}

//	Below block is a stuct
//type Doctor struct {
//	number     int
//	actorName  string
//	companions []string
//}
//
//func main() {
//	aDoctor := Doctor{
//		number:    3,
//		actorName: "Jon Pertwee",
//		companions: []string{
//			"Liz Shaw",
//			"Jo Grant",
//			"Sarah Jane Smith",
//		},
//	}
//	fmt.Println(aDoctor)
//}

//	Below is a anonomys struct
//func main() {
//	aDoctor := struct{ name string }{name: "John Pertwee"}
//	anotherDoctor := aDoctor
//	anotherDoctor.name = "Tom Baker"
//	fmt.Println(aDoctor)
//	fmt.Println(anotherDoctor)
//}

// Below is embedding

//type Animal struct {
//	Name   string
//	Origin string
//}
//
//type Bird struct {
//	Animal
//	SpeedKPH float32
//	CanFly   bool
//}
//
// Two func main options
//func main() {
//	b := Bird{}
//	b.Name = "Emu"
//	b.Origin = "Australia"
//	b.SpeedKPH = 48
//	b.CanFly = false
//	fmt.Println(b)
//}
//
//	OR
//
//func main() {
//	b := Bird{
//		Animal:   animal{Name: "Emu", Origin: "Australia"},
//		SpeedKPH: 48,
//		CanFly:   false,
//	}
//	fmt.Println(b)
//}

//	Below is Tag

type Animal struct {
	Origin string
	name   string
}

func main() {
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
