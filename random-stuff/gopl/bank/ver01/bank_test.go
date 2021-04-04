package bank_test

import (
	"fmt"
	"testing"

	bank "github.com/fpdevil/goprog/random-stuff/gopl/bank/ver01"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// wait for both the transactions
	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
