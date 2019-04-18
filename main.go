package main

import "fmt"
import "strconv"

func main() {
	fmt.Println("hello")

	var x string = "hello"
	fmt.Println(x)

	var y = " world" // infers type
	fmt.Println(x + y)

	const z = 10
	fmt.Println(z)

	myArray := [5]int{1, 2, 3, 4, 5} // shorthand syntax, type inferred

	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}

	var ( // multiple variable assignment
		a = 1
		b = 2
		c = 3
	)
	fmt.Println(a, b, c)

	if a == 1 {
		fmt.Println("a = " + strconv.Itoa(a))
	}

	for _, i := range myArray {
		switch i {
		case 1:
			fmt.Println("i = 1")
		case 2:
			fmt.Println("i = 2")
		case 3:
			fmt.Println("i = 3")
		default:
			fmt.Println("i was some other number")
		}
	}

	mySlice := make([]int /*slice length*/, 5 /*capacity*/, 10)
	fmt.Println("mySlice", mySlice)

	mySlice = myArray[0:2]
	fmt.Println("mySlice", mySlice)

	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)

	myMap := make(map[string]int)
	myMap["key1"] = 10
	myMap["key2"] = 11
	myMap["key3"] = 12
	fmt.Println(myMap)
	fmt.Println(myMap["key1"])
	fmt.Println(len(myMap))

	value, ok := myMap["key4"]
	fmt.Println(value, ok)

	if value, ok = myMap["key1"]; ok {
		fmt.Println(value, ok)
	}

	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
	}

	if element, ok := elements["H"]; ok {
		fmt.Println(element, ok)
	}

	elements2 := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
	}

	fmt.Println(elements2)
}
