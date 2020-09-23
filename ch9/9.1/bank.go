package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

type WithdrawInfo struct {
	amount int
	result chan bool
}
var withdraw = make(chan WithdrawInfo)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func Withdraw(amount int) bool {
	ch := make(chan bool)
	withdraw <- WithdrawInfo{amount, ch}
	return <-ch
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case wd := <-withdraw:
			tmp := balance - wd.amount
			if tmp >= 0 {
				balance = tmp
				wd.result <- true
			} else {
				wd.result <- false
			}
		case amount := <-deposits:
			balance += amount

		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
