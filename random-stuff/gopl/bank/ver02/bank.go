package bank

var (
	sema    = make(chan struct{}, 1) // a binary counting semaphore guarding balance
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{} // acquire a token
	balance += amount
	<-sema // release the token
}

func Balance() int {
	sema <- struct{}{} //acquire a token
	b := balance
	<-sema // release token
	return b
}
