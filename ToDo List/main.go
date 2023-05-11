package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var listData []string

func main() {
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
	case "0": // Create new file
		fmt.Println("You selected 0")
		newFile()
	case "1": // Load file
		fmt.Println("Loading file")
		listData = loadFile()
		editPage()
	case "2": // Save file
		fmt.Println("You selected 2")
		saveFile(listData)
		os.Exit(1)
	case "3": // Exit program
		fmt.Println("Closing program")
		os.Exit(1)
	default:
		fmt.Println("Incorrect input")
		homePage()
	}
}

// Page for viewing and editing the TODO list
func editPage() {
	var t string
	fmt.Printf("\n\nEditing page!!\n")
	fmt.Println("0 - Add to list         1 - Remove from list")
	fmt.Printf("2 - Print list          3 - Return to Home Page\n\n")
	fmt.Scan(&t)
	switch t {
	case "0": // Add item to list
		var i string
		fmt.Println("What would you like to add?")
		fmt.Scan(&i)
		listData = addToList(i)
		editPage()
	case "1": // remove item from list - doesn't work
		var r string
		fmt.Println("You selected 1")
		printData(listData)
		fmt.Println("Enter key to remove list item")
		fmt.Scan(&r)
	case "2": // Print current list items
		fmt.Println("You selected 2")
		printData(listData)
		editPage()
	case "3": // Return to home page
		fmt.Println("You selected 3")
		homePage()
	default:
		fmt.Println("Incorrect input")
		editPage()
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

// Calls String to byte slice function and saves to file
func saveFile(s []string) {
	i := sStringToByte(s)
	err := ioutil.WriteFile("My_list.txt", i, 0666)
	if err != nil {
		fmt.Println("Error: Doesn't wanna save a new file!!")
	} else {
		fmt.Print("Saved File!")
	}
}

// converts string slice to byte slice
func sStringToByte(l []string) []byte {
	var s string
	var k []byte
	s = strings.Join(l, ",")
	k = []byte(s)
	return k
}

// Prints data in the TODO list, 1 item per line
func printData(d []string) {
	fmt.Printf("\n\nThis is what you have to do!!\n")
	for i, s := range d {
		fmt.Println(i, s)
	}
}

// Adds user input to the listData variable then returns updated listData
func addToList(line string) []string {
	listData = append(listData, line)
	return listData
}
