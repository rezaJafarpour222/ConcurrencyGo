package main

import (
	"fmt"
	"sync"
)

type Income struct {
	Source string
	Amount int
}

func complex_race() {

	var wg sync.WaitGroup

	var banBalance int
	var balance sync.Mutex
	fmt.Printf("Initial account balance: $%d.00", banBalance)
	fmt.Println()
	incomes := []Income{
		{Source: "Main Job", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Part time job", Amount: 50},
		{Source: "Investment", Amount: 100},
	}
	wg.Add(len(incomes))
	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := banBalance
				temp += income.Amount
				banBalance = temp
				balance.Unlock()
				fmt.Printf("On week %d, you earn $%d from %s\n", week, income.Amount, income.Source)
			}
		}(i, income)
	}
	wg.Wait()
	fmt.Printf("final bank balance: $%d.00", banBalance)
	fmt.Println()
}
