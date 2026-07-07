package main

import (
	"fmt"
	"math"
)

// interfaces : collection of method signatures

type shape interface {
	area() float64
	perimeter() float64
}

func printShape(s shape) {
	fmt.Println(s.area(), s.perimeter())
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.height * r.width
}
func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// employee - interface

type emplyee interface {
	getName() string
	getSalary() int
}

func test(e emplyee) {
	fmt.Println(e.getName(), e.getSalary())
	fmt.Println("<------------------->")
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (ct contractor) getName() string {
	return ct.name
}
func (ct contractor) getSalary() int {
	return ct.hourlyPay * ct.hoursPerYear
}

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getName() string {
	return ft.name
}
func (ft fullTime) getSalary() int {
	return ft.salary
}

// assignment
type email struct {
	isSubscribed bool
	body         string
}

func (e email) cost() float64 {
	if e.isSubscribed {
		return float64(len(e.body)) * 0.01
	} else {
		return float64(len(e.body)) * 0.05
	}

}
func (e email) print() {
	fmt.Println(e.body)
}

type expense interface {
	cost() float64
}
type printer interface {
	print()
}

func print(p printer) {
	p.print()
}
func testM(e expense, p printer) {
	fmt.Printf("Printing with cost: $%.2f ...\n", e.cost())
	p.print()
}

func mains() {

	e := email{
		isSubscribed: true,
		body:         "hello, there, IM IN MOJO campus pg",
	}
	testM(e, e)
	e = email{
		isSubscribed: false,
		body:         "I want my money back",
	}
	testM(e, e)
	var rec = rect{
		width:  2,
		height: 3,
	}
	circ := circle{
		radius: 5.7,
	}
	printShape(rec)
	printShape(circ)

	test(fullTime{
		name:   "Pradeep",
		salary: 150000,
	})
	test(contractor{
		name:         "Kunal",
		hourlyPay:    700,
		hoursPerYear: 300,
	})
}
