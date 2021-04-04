package main

import (
	"sync"
	"testing"
)

func TestWithdraw(t *testing.T) {
	Deposit(9999)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(amount int) {
			Withdraw(amount)
			wg.Done()
		}(i)
	}

	wg.Wait()

	if got, want := Balance(), 5049; got != want {
		t.Errorf("got %#v, want: %#v", got, want)
	}
}
