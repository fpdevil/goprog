# Exercise 9.1

Add a `function Withdraw(amount int) bool` to _thegopl.io/ch9/bank1program_. The result should indicate whether the transaction **succeeded** or **failed** due to _insufficient funds_. The message sent to the monitor `gorputine` must contain both the amount to withdraw and a new channel over which the monitor `goroutine` can send the `boolean` result back to Withdraw.
