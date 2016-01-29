package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type dish struct {
	Name  string
	Bites int
}

type person struct {
	Name string
}

var dishes = []dish{
	{Name: "chorizo", Bites: 5},
	{Name: "chopitios", Bites: 10},
	{Name: "pimentos", Bites: 7},
	{Name: "big bowl of mayonnaise", Bites: 8},
	{Name: "patas bravas", Bites: 9},
}

var people = []person{
	{Name: "Alice"},
	{Name: "Bob"},
	{Name: "Charlie"},
	{Name: "Dave"},
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	var foodPool sync.Pool

	for _, d := range dishes {
		foodPool.Put(d)
	}

	fmt.Println("Bon appétit!")
	var wg sync.WaitGroup
	for _, p := range people {
		wg.Add(1)
		go p.dine(&wg, &foodPool)
	}
	wg.Wait()
	fmt.Println("That was delicious!")
}

func (p person) dine(wg *sync.WaitGroup, foodPool *sync.Pool) {
	for {
		if !foodRemains() {
			break
		}

		potentialDish := foodPool.Get()
		if potentialDish == nil {
			fmt.Printf("%s recieved a nil dish\n", p.Name)
			break
		}
		someDish, ok := potentialDish.(dish)
		if !ok {
			fmt.Printf("%s was unable to turn a potential dish into a real dish\n", p.Name)
			continue
		}
		p.consume(&someDish)
		if someDish.Bites <= 0 {
			fmt.Printf("%s finished %s\n", p.Name, someDish.Name)
		} else {
			foodPool.Put(someDish)
		}
	}
	wg.Done()
}

func (p person) consume(d *dish) {
	timeEating := time.Duration(rand.Intn(1500)+300) * time.Millisecond
	d.Bites--
	fmt.Printf("%s is enjoying some %s\n", p.Name, d.Name)
	time.Sleep(timeEating)
}

func foodRemains() bool {
	for _, d := range dishes {
		if d.Bites > 0 {
			return true
		}
	}
	return false
}
