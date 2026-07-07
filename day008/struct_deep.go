package main

import "fmt"

type car struct {
	model      string
	year       int
	frontWheel wheel
	backWheel  wheel
}
type wheel struct {
	radius   int
	material string
}

type student struct {
	roll_num int
	address  struct {
		city    string
		pincode int
	}
}
type basiccar struct {
	model string
	make  string
}
type suzuki struct {
	basiccar
	topspeed int
}
type rect struct {
	width  int
	height int
}

func (r rect) area() int {
	return r.height * r.width
}

func main() {
	r := rect{
		width:  17,
		height: 31,
	}
	fmt.Println(r.area())
	var mycar car
	var yourcar car
	thiscar := car{}
	mycar.backWheel.material = "Good"
	mycar.frontWheel.radius = 3
	yourcar.frontWheel.material = "GOod"
	yourcar.backWheel.radius = 5

	fmt.Println(mycar, yourcar, thiscar)

	// anonymous struct : incase we dont want any extra copies of the struct, we dont need more than 1 instance of a struct
	anonymousStruct := struct {
		name string
		dob  int
	}{
		name: "Pradeep",
		dob:  14111999}
	fmt.Println(anonymousStruct)

	student1 := student{}
	student1.roll_num = 1
	student1.address.city = "Haldwani"
	student1.address.pincode = 263139
	student2 := student{
		roll_num: 2,
		address: struct {
			city    string
			pincode int
		}{
			city: "Bhimtal", pincode: 263132}}

	fmt.Println(student1, student2)
	var newcar1 = suzuki{}
	newcar1.make = "Maruti"
	newcar1.model = "2023"
	newcar1.topspeed = 250

	newcar2 := suzuki{
		topspeed: 255,
		basiccar: basiccar{
			make:  "Maruti",
			model: "1990",
		},
	}
	fmt.Println(newcar1, newcar2)
}
