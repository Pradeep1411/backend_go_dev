package main

import "fmt"

func main() {
	//  Access elements of a slice
	prices := []int{1, 2, 3}
	fmt.Println(prices)
	fmt.Println(prices[0])
	fmt.Println(prices[1])
	fmt.Println(prices[2])

	// Changing elements
	prices[2] = -3
	fmt.Println(prices[2])

	//Append elements to a slice, slice_name=append(slice_name,element1,element2,...)

	fmt.Printf("prices : %v \n", prices)
	fmt.Printf("length : %d \n", len(prices))
	fmt.Printf("capacity : %d \n", cap(prices))

	prices = append(prices, 20, 21, 22, 5)
	fmt.Printf("prices : %v \n", prices)
	fmt.Printf("length : %d \n", len(prices))
	fmt.Printf("capacity : %d \n", cap(prices))

	// Append one slice to another slice, slice3=append(slice1, slice2...)
	// Note: The '...' after slice2 is necessary when appending the elements of one slice to another.

	myslice1 := []int{1, 2, 3}
	myslice2 := []int{4, 5, 6}
	myslice3 := append(myslice1, myslice2...)
	fmt.Printf("myslice3 =%v \n", myslice3)
	fmt.Printf("length=%d \n", len(myslice3))
	fmt.Printf("capacity=%d \n", cap(myslice3))

	// Changing the length of a slice

	arr1 := [...]int{9, 10, 11, 12, 13, 14}
	slice1 := arr1[1:5]

	fmt.Printf("slice1: %v \n", slice1)
	fmt.Printf("length: %d \n", len(slice1))
	fmt.Printf("capacity: %d \n", cap(slice1))

	// changing length by re-slicing
	slice1 = arr1[1:3]
	fmt.Printf("slice1: %v \n", slice1)
	fmt.Printf("length: %d \n", len(slice1))
	fmt.Printf("capacity: %d \n", cap(slice1))
	// changing length by appending items
	slice1 = append(slice1, 20, 21, 22, 23)
	fmt.Printf("slice1: %v \n", slice1)
	fmt.Printf("length: %d \n", len(slice1))
	fmt.Printf("capacity: %d \n", cap(slice1))

	// Memory efficiency : copy function
	// cop
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Printf("numbers: %v \n", numbers)
	fmt.Printf("length: %d \n", len(numbers))
	fmt.Printf("capacity: %d \n", cap(numbers))

	// create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)
	fmt.Printf("numbersCopy: %v \n", numbersCopy)
	fmt.Printf("length : %d \n", len(numbersCopy))
	fmt.Printf("capacity : %d \n", cap(numbersCopy))

}
