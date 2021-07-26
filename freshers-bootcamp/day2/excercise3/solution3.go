package main

import (
	"fmt"
	"sync"
)

var balance  uint64 = 500
var mutex = &sync.Mutex{}
func withdraw(amount uint64,wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	if amount<=balance {
		balance -= amount
	}else
	{
		fmt.Printf("Cannot withdraw amount %d\n",amount)
	}
	mutex.Unlock()
}
func deposit(amount uint64,wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	balance+= amount
	mutex.Unlock()
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(5)
	go deposit(100,&wg)
	go deposit(100,&wg)
	go deposit(100,&wg)
	go deposit(100,&wg)
	go deposit(100,&wg)
	go withdraw(1000,&wg)//gives a prompt regarding low balance
	wg.Wait()
	fmt.Println(balance)
}
