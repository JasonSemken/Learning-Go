package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type data []string

func main() {
	//fmt.Println(loadList())
	//fmt.Println(loadFile())
	homePage()
}

// Home page for ToDo list
func homePage() {
	var uI string
	fmt.Printf("\nWelcome to the homepage!!\nWhat do you wanna do?\n\n")
	fmt.Println("0 - Start a new File   1 - Load file")
	fmt.Printf("2 - Save File          3 - Exit program\n\n")
	fmt.Scan(&uI)
	switch uI {
	case "0":
		fmt.Println("You selected 0")
		newFile()
	case "1":
		fmt.Println("Loading file")
		printData(loadFile())
	case "2":
		fmt.Println("You selected 2")
	case "3":
		fmt.Println("Closing program")
		os.Exit(1)
	default:
		fmt.Println("Incorrect input")
		homePage()
	}
}

// Reads directory for file named 'My_list.txt', if file found data loaded to byte, if no file found empty file created
func loadFile() []string {
	var s []string
	d, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("Error:")
	}
	for _, f := range d {
		if f.Name() == "My_list.txt" {
			fmt.Println("Yay! Found a file!")
			d, _ := ioutil.ReadFile("My_list.txt")
			s := strings.Split(string(d), ",")
			return s
		} else {
			fmt.Println("No file Found :( Created new file!!")
			newFile()
		}
	}
	return s
}

// Creates new file
func newFile() {
	e := []byte("")
	err := ioutil.WriteFile("My_list.txt", e, 0666)
	if err != nil {
		fmt.Println("Error: Doesn't wanna save a new file!!")
	}
}

func printData(d []string) {
	for i, s := range d {
		fmt.Println(i, s)
	}
}
