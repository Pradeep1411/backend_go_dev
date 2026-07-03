package main

import "fmt"

func main() {
	//  syntax
	// for st1;st2;st3{ code ...}
	for i := 0; i < 5; i += 2 {
		fmt.Println(i)
	}

	// continue statement : skip the execution for if found that condition true
	for i := 4; i < 11; i += 3 {
		if i == 7 {
			continue
		}
		fmt.Println(i)
	}
	// break statement : breaks the execution/loop if found that condition true
	for j := -1; j > -5; j-- {
		if j == -3 {
			break
		}
		fmt.Println(j)
	}
	// Nested loops
	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"banana", "apple", "orange"}
	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}
	}
	// range : it returns both index & value
	//  for index,value := range array|slice|map{}
	for idx, val := range fruits {
		fmt.Printf("%v \t %v \n", idx, val)
	}

	// we can omit either the value or index using underscore _
	for id := range fruits {
		fmt.Println(id)
	}

}
