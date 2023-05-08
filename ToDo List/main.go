package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//fmt.Println(loadList())
	newList()
}

// Reads directory for file named 'My_list.txt', if no file s found a new file is created
func newList() {
	d, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("Error:")
	}

	// For loop to search found file in directory for file named 'My_list.txt'
	for _, f := range d {
		if f.Name() == "My_list.txt" {
			fmt.Println("Yay! Found a file!")
		} else {
			e := []byte("")
			err := ioutil.WriteFile("My_list.txt", e, 0666)
			if err != nil {
				fmt.Println("Error: Doesn't wanna save a new file!!")
			}
		}
	}
}
