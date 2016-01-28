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
	Lock  sync.Mutex
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

	fmt.Println("Bon appétit!")
	var wg sync.WaitGroup
	for _, p := range people {
		wg.Add(1)
		go p.dine(wg, dishes)
	}
	wg.Wait()
	fmt.Println("That was delicious!")
}

func (p person) dine(wg sync.WaitGroup, d []dish) {
	dishIndex := rand.Intn(5)
	for {
		if !foodRemains(d) {
			break
		}
		if d[dishIndex].Bites > 0 {
			d[dishIndex].Lock.Lock()
			p.consume(&d[dishIndex])
			d[dishIndex].Lock.Unlock()
		}
		dishIndex++
		dishIndex = dishIndex % len(dishes)
	}
	wg.Done()
}

func (p person) consume(d *dish) {
	timeEating := time.Duration(rand.Intn(150)+30) * time.Millisecond
	d.Bites--
	fmt.Printf("%s is enjoying some %s\n", p.Name, d.Name)
	time.Sleep(timeEating)
}

func foodRemains(dishes []dish) bool {
	for _, d := range dishes {
		if d.Bites > 0 {
			return true
		}
	}
	return false
}
