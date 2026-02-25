package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

var pizzaMade, pizzasFailed, total int
var NumberOfPizzas = 9

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}
type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}
func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= pizzaNumber {
		delay := rand.Intn(5) + 1
		fmt.Printf("Recieved order number %d!\n", pizzaNumber)
		rnd := rand.Intn(12) + 1
		message := ""
		success := false
		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzaMade++
		}
		total++
		fmt.Printf("Making pizza %d. It will take %d seconds...\n", pizzaNumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)
		if rnd <= 2 {
			message = fmt.Sprint("*** We ran out of ingredients for pizza %d!", pizzaNumber)
		} else if rnd <= 4 {
			message = fmt.Sprint("*** The Cook quit while making pizza %d!", pizzaNumber)
		} else {
			success = true
			message = fmt.Sprint("Pizza order %d is ready", pizzaNumber)
		}
		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     message,
			success:     success,
		}
		return &p
	}
	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}

}
func pizzeria(pizzaMaker *Producer) {
	var i = 0
	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			case pizzaMaker.data <- *currentPizza:
			case quitChan := <-pizzaMaker.quit:
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
	}
}

func pub_sub() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	color.Cyan("The Pizzeria is open for business!")
	color.Cyan("-------------------------")
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}
	go pizzeria(pizzaJob)
	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order %d is out for delivery!", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The customer is really mad!")
			}
		} else {
			color.Cyan("Done making pizzas...")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("*** Error closing channel!", err)
			}
		}
	}

	color.Cyan("---------------------")
	color.Cyan("Done for the day.")
	color.Cyan("we made %d pizzas, but failed to make %d, with %d attepmts in total", pizzaMade, pizzasFailed, total)
	switch {
	case pizzasFailed >= 9:
		color.Red("It was a bad day...")
	case pizzasFailed >= 6:
		color.Red("It was not a good  day...")
	case pizzasFailed >= 4:
		color.Yellow("It was not an okay  day...")
	case pizzasFailed >= 2:
		color.Red("It was good  day...")
	}
}
