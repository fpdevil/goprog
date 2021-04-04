package bank_test

import (
	"sync"
	"testing"

	bank "github.com/fpdevil/goprog/random-stuff/gopl/bank/ver03"
)

func TestBank(t *testing.T) {
	// Concurrently deposit Deposit [1..1000]
	var wg sync.WaitGroup
	for i := 0; i <= 1000; i++ {
		wg.Add(1)
		func(amount int) {
			bank.Deposit(i)
			wg.Done()
		}(i)
	}

	wg.Wait()

	// now test for results
	if got, want := bank.Balance(), (1000+1)*1000/2; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
