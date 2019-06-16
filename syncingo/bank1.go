package main

var sema = make(chan struct{}, 1)
var balance int

func Deposit(amount int) {
	sema <- struct{}{}
	balance = balance + amount
	<-sema
}

// func Withdraw() int {
// 	return <-balances
// }
func Balance() int {

	sema <- struct{}{}
	b := balance
	<-sema
	return b

}
func Withdraw(amount int) bool {
	Deposit(-amount)
	if Balance() < 0 {
		Deposit(amount)
		return false
	}
	return true
}

// func teller() {
// 	var balance int
// 	for {
// 		select {
// 		case amount := <-deposits:
// 			balance += amount

// 		case balances <- balance:

// 		case w := <-withdraws:
// 			if balance < w {
// 				fmt.Println("error! overflow!")
// 			} else {
// 				balance -= w
// 				fmt.Println(balance)
// 			}
// 		}
// 	}

// }
