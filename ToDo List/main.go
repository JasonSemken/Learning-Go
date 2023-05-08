package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//fmt.Println(loadList())
	fmt.Println(newList())
}

// Reads directory for file named 'My_list.txt', if file found data loaded to byte, if no file found empty file created
func newList() []byte {
	var data []byte

	d, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("Error:")
	}
	for _, f := range d {
		if f.Name() == "My_list.txt" {
			fmt.Println("Yay! Found a file!")
			data, _ := ioutil.ReadFile("My_list.txt")
			return data
		} else {
			fmt.Println("No file Found :( Created new file!!")
			newFile()
		}
	}
	return data
}

// Creates new file
func newFile() {
	e := []byte("")
	err := ioutil.WriteFile("My_list.txt", e, 0666)
	if err != nil {
		fmt.Println("Error: Doesn't wanna save a new file!!")
	}
}
