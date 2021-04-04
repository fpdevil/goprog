package main

var (
	deposits = make(chan int)  // send amount to deposit
	balances = make(chan int)  // receive balance
	withdraw = make(chan int)  // send amount to withdraw
	result   = make(chan bool) // send boolean result for withdraw
)

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func Withdraw(amount int) bool {
	withdraw <- amount
	return <-result
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case amount := <-withdraw:
			if amount <= balance {
				balance -= amount
				result <- true
			} else {
				result <- false
			}
		}
	}
}

func init() {
	go teller()
}
