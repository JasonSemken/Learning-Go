package main

import "fmt"

func main() {
	//	var students [3]string
	//	fmt.Printf("Students: %v\n", students)
	//	students[0] = "Lisa"
	//	students[1] = "Shrek"
	//	students[2] = "Dumbo"
	//	fmt.Printf("Students: %v\n", students)
	//	fmt.Printf("Number of Students: %v\n\n", len(students))

	//	var identityMatrix [3][3]int
	//	identityMatrix[0] = [3]int{1, 0, 0}
	//	identityMatrix[1] = [3]int{1, 1, 0}
	//	identityMatrix[2] = [3]int{1, 1, 1}
	//	fmt.Println(identityMatrix)

	//	a := []int{1, 2, 3}
	//	b := a
	//	b[0] = 5
	//	fmt.Println(a)
	//	fmt.Println(b)
	//	fmt.Printf("Length: %v\n", len(a))
	//	fmt.Printf("Capacity: %v\n", cap(a))

	//	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//	b := a[:]   // slice of all elements
	//	c := a[3:]  // slice from 4th element to end
	//	d := a[:6]  // slice first 6 elements
	//	e := a[3:6] // slice the 4th, 5th, and 6th elements
	//	a[5] = 50
	//	fmt.Println(a)
	//	fmt.Println(b)
	//	fmt.Println(c)
	//	fmt.Println(d)
	//	fmt.Println(e)

	//	a := make([]int, 3, 100)
	//	fmt.Println(a)
	//	fmt.Printf("Length: %v\n", len(a))
	//	fmt.Printf("Capacity: %v\n", cap(a))

	//	a := []int{}
	//	fmt.Println(a)
	//	fmt.Printf("Length: %v\n", len(a))
	//	fmt.Printf("Capacity: %v\n", cap(a))
	//	a = append(a, 1)
	//	fmt.Println(a)
	//	fmt.Printf("Length: %v\n", len(a))
	//	fmt.Printf("Capacity: %v\n", cap(a))
	//	a = append(a, 2, 3, 4, 5)
	//	fmt.Println(a)
	//	fmt.Printf("Length: %v\n", len(a))
	//	fmt.Printf("Capacity: %v\n", cap(a))

	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	b := append(a[:2], a[4:]...)
	fmt.Println(b)
	fmt.Println(a)

}
