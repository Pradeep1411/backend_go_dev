package main

import "fmt"

func main() {
	// Create variables of different types.
	var Name = "Pradeep"
	Age := 28
	var Salary = 10000.50
	IsWorking := true
	fmt.Printf(" Name : %v , type : %T \n Age: %v , type : %T \n Salary : %v , type : %T \n IsWorking : %v , type : %T \n", Name, Name, Age, Age, Salary, Salary, IsWorking, IsWorking)
	Age = 29
	fmt.Printf("Age : %v, type : %T\n", Age, Age)

	// Constants
	const SEX = "MALE"
	const COMPANY_NAME = "STOCKGRO"
	fmt.Printf("SEX : %v, type : %T\n", SEX, SEX)
	fmt.Printf("COMPANY_NAME : %v, type : %T\n", COMPANY_NAME, COMPANY_NAME)

	// Arrays
	var arr1 = [3]string{"Praddep", "Recro", "Stockgro"}
	var arr2 = [...]int{1, 2, 3}
	arr3 := [...]string{"Haldwani", "Delhi", "Bengaluru"}
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(arr1[0], arr1[len(arr1)-1], len(arr1))

	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}
	arr4 := [3]int{1}
	fmt.Println(arr4)
	// specific initialization
	arr5 := [7]int{2: 55, 8, 5: -55}
	fmt.Println(arr5, len(arr5))

	max := 0
	min := 0
	for i, v := range arr5 {

		if max < v {
			max = v
		} else if min > v {
			min = v
		} else if i == len(arr5)-1 {
			fmt.Println("Max : ", max)
			fmt.Println("Min : ", min)

		}
	}
	fmt.Println(arr2)
	// reversing arr2
	var (
		i = 0
		j = len(arr2) - 1
	)
	for i < j {

		// temp := arr2[i]
		// arr2[i] = arr2[j]
		// arr2[j] = temp
		arr2[i], arr2[j] = arr2[j], arr2[i]

		i++
		j--

	}
	fmt.Println(arr2)

}
