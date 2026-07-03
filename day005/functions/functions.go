package main

import "fmt"

func myMessage() {
	fmt.Println("I just got executed")
}

func multipleParams(age int, designation string, isMale bool) {
	fmt.Printf("Age is %v, designation is %v & isMale =%v \n", age, designation, isMale)
}

func retunType(age int) string {
	if age >= 18 {
		return "Eligible for voting"
	} else {
		return "Not eligible for voting"
	}
}

func namedReturnValue(age int) (tenureLeft int) {
	tenureLeft = 60 - age
	if tenureLeft < 0 {
		tenureLeft = 0
	}
	return
}

func multipleReturnValues(age int) (result int, txt string) {
	result = age - 18
	if age >= 18 {
		txt = "Eligible for vote"
	} else {

		txt = "Not eligible for vote"
	}
	return
}

func rec(x int) int {
	if x == 17 {
		return x
	}
	fmt.Println(x)
	return rec(x + 2)
}

func main() {
	myMessage()
	multipleParams(30, "Senior Software Engineer", true)
	fmt.Println(retunType(18))
	fmt.Println(namedReturnValue(75))
	fmt.Println(multipleReturnValues(25))
	a, b := multipleReturnValues(55)
	fmt.Println(a, b)

	_, y := multipleReturnValues(13)
	fmt.Println(y)
	x, _ := multipleReturnValues(44)
	fmt.Println(x)

	fmt.Println("The ans is : ", rec(1))
}
