package main

import "fmt"

func main() {
	// var a= map[keytype]valueType{key1:value1,key2:value2...}
	//  b:=map [keytype]valueType{key1:value1,key2:value2...}
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Haldwani": 5, "Nainital": 1, "Bhimtal": 2, "Sattal": 3}
	fmt.Printf("a \t %v \n", a)
	fmt.Printf("b \t %v \n", b)

	// create map using main() func

	//  var a=make(map[keytype]valuetype)
	// b:=make(map[keytype]valuetype)
	var c = make(map[string]string) // creating an empty map
	fmt.Println("c: ", c)
	c["name"] = "Pradeep"
	c["age"] = "27"
	fmt.Println("c's age: ", c["age"])

	var d map[int]string
	fmt.Println("d :", d)
	// deleting an element from a map

	// delete(b, "Haldwani")
	// fmt.Println("b :", b)

	// Checking specific elements in a map
	val1, ok1 := b["Haldwani"]
	fmt.Println(val1, ok1)
	e := c
	fmt.Println(e, c)
	e["newKey"] = "Only for E"
	fmt.Println(e, c)

	// iterate over maps
	for k, v := range c {
		fmt.Println("Key : ", k, " value: ", v)
	}

	// iterating over a map in specefic order
	mapx := map[string]int{"one": 1, "Two": 2, "Three": 3, "Four": 4}
	var order []string
	order = append(order, "One", "Two", "Three", "Four")
	for _, element := range order {
		fmt.Printf("%v : %v \n", element, mapx[element])
	}

}
