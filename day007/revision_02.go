package main

import "fmt"

func main() {
	// slices -empty slice
	sl1 := []int{}
	fmt.Println(sl1)
	fmt.Printf("%T\n", sl1)
	var sl2 = []int{3, 4, 5}
	fmt.Println(sl2)
	sl3 := make([]string, 5)
	fmt.Println("sl3: ", sl3)
	fmt.Printf("%v\t%v\n", len(sl3), cap(sl3))

	ar := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sl4 := ar[:]
	sl41 := ar[1:5]
	fmt.Println(sl4, sl41)

	// nil slice
	var nilSlice []int
	fmt.Println(sl1 == nil)
	fmt.Println(nilSlice == nil)

	//indexing or accessing elements of a slice
	fmt.Println(sl4[2])
	sl4[2] = -555
	fmt.Println(sl4)
	fmt.Println(sl41)
	fmt.Println(ar)
	sl4 = append(sl4, 99, 99, 99, 99, 99, 99, 99)
	sl4[2] = -999
	fmt.Println(sl4)
	fmt.Println(sl41)
	fmt.Println(ar)

	sl21 := []int{9, 9, 9}
	sl213 := append(sl21, sl2...)
	fmt.Println(sl213)

	// copy function
	copy(sl21, sl213)
	fmt.Println(sl21)

	var nilMap map[int]string
	nonNil := map[int]string{0: "a", 1: "b", 2: "c"}
	fmt.Println(nilMap == nil)
	fmt.Println(nonNil == nil)

	delete(nonNil, 1)
	val, ok := nonNil[1]
	fmt.Println(val, ok)
	type Employee struct {
		id    int
		name  string
		email string
		dept  string
	}
	var emp1 Employee
	emp1.dept = "QA"
	emp1.email = "qa@stockgro.com"
	emp1.id = 123
	emp1.name = "PD"
	fmt.Println(emp1)

}

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
