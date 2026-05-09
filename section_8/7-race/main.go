package main

import (
	"fmt"
	"sync"
	"time"
)

type BankAccount struct {
	balance int
	mutex   sync.Mutex
}

func (b *BankAccount) Deposit(amount int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.balance += amount
	fmt.Println("Depositing ", amount)
}

func (b *BankAccount) Withdraw(amount int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if amount > b.balance {
		fmt.Println("We cann't withdraw this amount ", amount)
		return
	}
	b.balance -= amount
	fmt.Println("Withdrawing ", amount)
}

func (b *BankAccount) Balance() int {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.balance
}

func main() {
	var wg sync.WaitGroup
	var account = &BankAccount{
		balance: 100,
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(amount int) {
			defer wg.Done()
			time.Sleep(time.Duration(amount) * time.Millisecond)
			account.Deposit(amount)
		}(i + 1)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(amount int) {
			defer wg.Done()
			time.Sleep(time.Duration(amount) * time.Millisecond)
			account.Withdraw(amount * 10)
		}(i + 1)
	}

	wg.Wait()
	fmt.Println(account.Balance())

	//counter := 0
	//var wg sync.WaitGroup
	//var mu sync.Mutex
	//
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		mu.Lock()
	//		counter++
	//		mu.Unlock()
	//	}()
	//}
	//
	//wg.Wait()
	//fmt.Println(counter)
}
