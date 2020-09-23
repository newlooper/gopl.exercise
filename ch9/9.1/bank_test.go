package bank_test

import (
	"fmt"
	"testing"

	. "gopl.exercise/ch9/9.1"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		Deposit(200)
		fmt.Println(Withdraw(200))
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		Deposit(50)
		fmt.Println(Withdraw(50))
		Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := Balance(), 100; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}

func TestWithdraw(t *testing.T) {
	Deposit(100)
	balance := Balance()
	ok := Withdraw(50)
	if !ok {
		t.Errorf("ok = false, want true. balance = %d", Balance())
	}
	expected := balance - 50
	if balance := Balance(); balance != expected {
		t.Errorf("balance = %d, want %d", balance, expected)
	}
}

func TestWithdrawFails(t *testing.T) {
	before := Balance()
	ok := Withdraw(before + 1)
	after := Balance()
	if ok {
		t.Errorf("ok = true, want false. balance = %d", after)
	}
	if after != before {
		t.Errorf("balance = %d, want %d", after, before)
	}
}
