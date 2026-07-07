package main

// type assertions
func getExpenseReport(e expenses) (string, float64) {
	val, ok := e.(emails)
	if ok {
		return val.toAddress, val.cost()
	}
	val1, ok1 := e.(sms)
	if ok1 {
		return val1.toPhoneNumber, val1.cost()
	}
	return "", 0.0
}

func (e emails) cost() float64 {
	if e.isSubscribed {
		return float64(len(e.body)) * 0.01
	} else {
		return float64(len(e.body)) * 0.05
	}

}
func (s sms) cost() float64 {
	if s.isSubscribed {
		return float64(len(s.body)) * 0.01
	} else {
		return float64(len(s.body)) * 0.05
	}

}
func (i invalid) cost() float64 {
	return 0.0
}

type expenses interface {
	cost() float64
}
type emails struct {
	isSubscribed bool
	body         string
	toAddress    string
}
type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}
type invalid struct{}

func main() {

}
