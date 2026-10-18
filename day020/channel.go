package main

import (
	"fmt"
	"time"
)

func filterOldChannels(emails []email) {
	isOldChannel := make(chan bool)
	go func() {
		for _, e := range emails {
			if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
				isOldChannel <- true
			}
			isOldChannel <- false
		}
	}()

	isOld := <-isOldChannel
	fmt.Println("email 1 is old:", isOld)
	isOld = <-isOldChannel
	fmt.Println("email 2 is old:", isOld)
	isOld = <-isOldChannel
	fmt.Println("email 3 is old:", isOld)

}

type email struct {
	date time.Time
}

func waitForDbs(numDBs int, dbChain chan struct{}) {
	for i := 0; i < numDBs; i++ {
		<-dbChain
	}
}
func test(numDBs int) {
	dbChan := getDatabsesChannel(numDBs)
	fmt.Printf("Waiting for %v databases...\n", numDBs)
	waitForDbs(numDBs, dbChan)
	fmt.Println("All databases are online!")
	fmt.Println("======================")
	time.Sleep(time.Millisecond * 10)
}

func main() {
	test(3)
	test(4)
	test(5)
}
func getDatabsesChannel(numDbs int) chan struct{} {
	ch := make(chan struct{})
	go func() {
		for i := 0; i < numDbs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
		}
	}()
	return ch
}
