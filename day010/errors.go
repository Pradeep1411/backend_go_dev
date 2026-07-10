package main

import "fmt"

//  error handling
// error interface is builtin

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
	costForCus, error := sendSMS(msgToCustomer)
	if error != nil {
		return 0.0, error
	}
	costForSpo, error := sendSMS(msgToSpouse)
	if error != nil {
		return 0.0, error
	}
	return costForCus + costForSpo, nil
}
func sendSMS(message string) (float64, error) {
	const maxTextLen = 25
	const costPerChar = 0.0002
	if len(message) > maxTextLen {
		return 0.0, fmt.Errorf("Canrt send text over %v character ", maxTextLen)
	}
	return costPerChar * float64(len(message)), nil

}

func testt(msgToCustomer, msgToSpouse string) {
	fmt.Println(" Message for customer", msgToCustomer)
	fmt.Println(" Message for spouse", msgToSpouse)
	totalCost, err := sendSMSToCouple(msgToCustomer, msgToSpouse)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Total cost : $%.4f\n", totalCost)

}

// formatting string using %v and .2f, .3f as per decimal places we need

// defining our own error type
type userError struct {
	name string
}

func (e userError) Error() string {
	return fmt.Sprintf("%v has problem with their account", e.name)

}

/*
	func sendSMS(msg,userName string)error{
		if !canSendToUser(userName){
			return userError{name: userName}
		}
		// rest logic
	}
*/
func main() {
	testt("My name is pradeep", "My name is govinda")
	testt("My name is pradeep,My name is pradeep", "My name is govinda")

}
