package main

import (
	"errors"
	"fmt"
)

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	userMap := make(map[string]user)
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("Invalid sizes")
	}
	for i := 0; i < len(names); i++ {
		name := names[i]
		phoneNumber := phoneNumbers[i]
		userMap[name] = user{
			name:        name,
			phoneNumber: phoneNumber,
		}
	}
	return userMap, nil
}

type user struct {
	name        string
	phoneNumber int
}

func test(names []string, phoneNumber []int) {
	fmt.Println("Creating Map...")
	users, errorMsg := getUserMap(names, phoneNumber)
	if errorMsg != nil {
		fmt.Println("Error Message : ", errorMsg)
		fmt.Println("~~~~~~~~~~~~~~~~~~~")
		return
	}
	for _, name := range names {
		fmt.Printf("Key:%v, value:\n", name)
		fmt.Println("-name : ", users[name].name)
		fmt.Println("-number : ", users[name].phoneNumber)
	}
	fmt.Println("~~~~~~~~~~~~~~~~~~~")
}
func main() {
	test(
		[]string{"John", "Bob", "Jill"},
		[]int{7895684593, 8077742518, 7896531246},
	)
	test(
		[]string{"John", "Bob"},
		[]int{7895684593, 8077742518, 7896531246},
	)
	test(
		[]string{"Gorge", "Sally", "Jill", "Sue"},
		[]int{5555684593, 805555558, 7855551246, 9999911111},
	)

}
