package main

import "fmt"

type Person struct {
	name        string
	age         int
	designation string
	salary      int
}

func printAll(pers Person) {
	fmt.Println("Name : ", pers.name)
	fmt.Println("Age : ", pers.age)
	fmt.Println("Designation : ", pers.designation)
	fmt.Println("Salary : ", pers.salary)
}

func main() {
	var pers1 Person
	var pers2 Person

	pers1.name = "Shivangi"
	pers1.age = 29
	pers1.designation = "Senior QA"
	pers1.salary = 100000

	pers2.name = "Ajinkya"
	pers2.age = 30
	pers2.designation = "UP Lead"
	pers2.salary = 200000
	fmt.Println(pers1)

	printAll(pers1)
	printAll(pers2)

}
