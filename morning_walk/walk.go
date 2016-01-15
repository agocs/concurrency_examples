package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	println("Let's go for a walk!")

	var wg sync.WaitGroup
	var wg2 sync.WaitGroup

	wg.Add(2)
	go getReady(&wg, "Alice")
	go getReady(&wg, "Bob")
	wg.Wait()

	wg2.Add(1)
	fmt.Println("Arming alarm")
	go armAlarm(&wg2)

	wg.Add(2)
	go putOnShoes(&wg, "Alice")
	go putOnShoes(&wg, "Bob")
	wg.Wait()
	fmt.Println("Exiting and locking door.")

	wg2.Wait()
}

func getReady(wg *sync.WaitGroup, name string) {
	sleepTime := time.Duration(rand.Intn(31)+60) * time.Millisecond
	fmt.Printf("%v started getting ready\n", name)
	time.Sleep(sleepTime)
	wg.Done()
	fmt.Printf("%v got ready in %v\n", name, sleepTime)
}

func armAlarm(wg *sync.WaitGroup) {
	fmt.Println("Alarm is counting down...")
	sleepTime := time.Duration(60) * time.Millisecond
	time.Sleep(sleepTime)
	fmt.Println("Alarm is armed.")
	wg.Done()
}

func putOnShoes(wg *sync.WaitGroup, name string) {
	fmt.Printf("%v is putting on shoes.\n", name)
	sleepTime := time.Duration(rand.Intn(11)+35) * time.Millisecond
	time.Sleep(sleepTime)
	wg.Done()
	fmt.Printf("%v spent %v putting on shoes\n", name, sleepTime)
}
