package main

import (
	"fmt"
	"sync"
)

type Creature interface {
	speak(s string) string
}

// the following 2 structs implicitly implement the Creature interface
type Person struct {
	Name string
}

func (p Person) speak (s string) string {
	return fmt.Sprintf("I'm a person called %s and I'm saying %s!", p.Name, s)
}


type Cat struct {
	Name string
}

func (c Cat) speak (s string) string {
	return fmt.Sprintf("I'm a cat called %s and i'm thinking %s!", c.Name, s)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	msgs := make(chan string, 2)	// set the buffer size to 2 - we expect to receive 2 strings

	p := Person{Name: "Ross"}
	go goSpeak(&wg, &p, msgs)
	
	c := Cat{Name: "toots"}
	go goSpeak(&wg, &c, msgs)

	wg.Wait()	// wait for the goroutines to finish before we try to receive from the channel

	close(msgs)	// close the channel so that when we 'range' over it, we don't deadlock

	for msg := range msgs {
		fmt.Printf("msg received: %s\n", msg)
	}
}

func goSpeak(wg *sync.WaitGroup, c Creature, ch chan string) {
	defer wg.Done()
	ch <- c.speak("BALLS")
}