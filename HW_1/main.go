package main

import (
	"fmt"
)

func main() {
	//1
	var (
		hello      string = "hello "
		world      string = "world"
		helloWorld string
	)

	helloWorld = hello + world
	fmt.Println(helloWorld)

	//2
	var number int = 123

	fmt.Printf("%s %d \n", helloWorld, number)

	//3
	anotherNumber := 321

	fmt.Printf("%d + %d = %d\n", number, anotherNumber, number+anotherNumber)
	fmt.Printf("%d - %d = %d\n", anotherNumber, number, anotherNumber-number)
	fmt.Printf("%d * %d = %d\n", anotherNumber, number, anotherNumber*number)
	fmt.Printf("%d / %d = %d\n", anotherNumber, anotherNumber, anotherNumber/anotherNumber)

	//4
	var (
		floatVar        float64 = 2.2
		anotherFloatVar float64 = 3.3
	)

	if floatVar > anotherFloatVar {
		fmt.Printf("%.1f greater than %.1f\n", floatVar, anotherFloatVar)
	} else if floatVar < anotherFloatVar {
		fmt.Printf("%.1f less than %.1f\n", floatVar, anotherFloatVar)
	} else {
		fmt.Printf("it`s equal")
	}

	floatVar = 3.3

	if floatVar == anotherFloatVar {
		fmt.Println("floats are equal")
	} else {
		fmt.Println("floats are not equal")
	}

	if floatVar != 3.2 {
		fmt.Printf("%.1f is not equal 3.2\n", floatVar)
	}

	if floatVar >= 3.2 {
		fmt.Println("var is greater or equal 3.2")
	}

	if floatVar <= 3.3 {
		fmt.Println("var is less or equal 3.3")
	}

	//5
	var fruit string = "apple"

	switch fruit {
	case "banana":
		fmt.Println("it`s BANANA!")
	case "apple":
		fmt.Println("it`s APPLE!")
	case "pineapple":
		fmt.Println("it`s PINAPPLE!")
	}

	//6
	switch {
	case fruit == "banana":
		fmt.Println("it`s second BANANA!")
	case fruit == "apple":
		fmt.Println("it`s second APPLE!")
		fallthrough
	case fruit == "pineapple":
		fmt.Println("it`s second PINAPPLE or APPLE!")
	}

	//7
	switch {
	case fruit == "banana":
		fmt.Println("it`s second BANANA!")
	case fruit == "pineapple":
		fmt.Println("it`s second PINAPPLE!")
	default:
		fmt.Println("Maybe it`s apple?")
	}

	//8
	arr := [5]int{2, 8, 28, 15, 22}

	arr[3]++
	arr[4]--
	fmt.Print(arr, "\n")

	//9
	anotherSlice := make([]string, 0)
	slice := []string{"a", "b", "c"}

	len := fmt.Sprint(len(slice))
	anotherSlice = append(anotherSlice, helloWorld)
	fmt.Println(anotherSlice)

	anotherSlice = append(anotherSlice, len)
	fmt.Println(anotherSlice[1])

}
