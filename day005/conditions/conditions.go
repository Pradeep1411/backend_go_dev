package main

import "fmt"

func main() {
	// if else if else
	x := -5
	y := 15
	if x > y {
		fmt.Printf("x: %v is greator than y:%d \n", x, y)
	} else if (x == -20) || (x < 0) {
		fmt.Println("WOW its a match")
	} else {
		fmt.Printf("x: %v is less than y:%d \n", x, y)
	}

	// switch
	day1 := "Sunday"
	switch day1 {
	case ("Sunday"):
		fmt.Println("Holiday as today is Sunday")
	case "Mondays":
		fmt.Println("Get back to work , Monday again")
	default:
		fmt.Println("Enjoy you're dead!")

	}
	// multi case switch statement
	day := 5

	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}

}
