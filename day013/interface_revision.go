package main

import "fmt"

// type assertions
// type swtich

func getExpenseReport(e expense) (string, float64) {
	// type switch
	switch v := e.(type) {
	case email:
		return v.toAddress, v.cost()
	case sms:
		return v.toPhoneNumber, v.cost()
	default:
		return "", 0.0
	}
	// type assetions
	// em, ok := e.(email)
	// if ok {
	// 	return em.toAddress, em.cost()
	// }
	// sm, ok1 := e.(sms)
	// if ok1 {
	// 	return sm.toPhoneNumber, sm.cost()
	// }
	// return "", 0.0
}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * 0.05
	}
	return float64(len(e.body)) * 0.01
}
func (i invalid) cost() float64 {
	return 0.0
}
func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * 0.1
	}
	return float64(len(s.body)) * 0.03
}

type expense interface {
	cost() float64
}
type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}
type invalid struct{}
type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

func test(e expense) {
	address, cost := getExpenseReport(e)
	//  type switch
	switch e.(type) {
	case email:
		fmt.Printf("Report: The email going to %s will cost : %.2f \n", address, cost)
	case sms:
		fmt.Printf("Report: The sms going to %s will cost : %.2f \n", address, cost)
	default:
		fmt.Println("Invalid expense")
	}
}
func main() {
	test(email{
		isSubscribed: true,
		body:         "Hello there",
		toAddress:    "pradeep@stockgor.com",
	})
	test(email{
		isSubscribed: false,
		body:         "Hello there,SS",
		toAddress:    "shivnagi@stockgor.com",
	})
	test(sms{
		isSubscribed:  true,
		body:          "This meeting could have been an email",
		toPhoneNumber: "7105440920",
	})
	test(sms{
		isSubscribed:  false,
		body:          "This meeting could have been an email",
		toPhoneNumber: "7105440930",
	})
	test(invalid{})

}
