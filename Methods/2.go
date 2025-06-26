package main

import (
	"fmt"
)

type BankAccount struct {
	owner   string
	balance float64
}

func NewBankAccount(owner string, initialBalance float64) *BankAccount {
	return &BankAccount{
		owner:   owner,
		balance: initialBalance,
	}
}

func (acc *BankAccount) Deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("Неправильная сумма, она должна быть больше 0")
	} else {
		acc.balance += amount
		fmt.Printf("Зачислено %.2f руб.\n", amount)
	}
}

func (acc *BankAccount) Withdraw(amount float64) {
	if amount <= 0 {
		fmt.Println("Неправильная сумма, она должна быть больше 0")
	}
	if amount > acc.balance {
		fmt.Println("Недостаточно средств для снятия")
	} else {
		acc.balance -= amount
		fmt.Printf("Снято %.2f руб.\n", amount)
	}

}

func (acc BankAccount) Balance() float64 {
	return acc.balance
}

func main() {
	account := NewBankAccount("Волчков Владислав", 100000)
	fmt.Printf("Баланс счета %s: %.2f руб.\n", account.owner, account.Balance())

	account.Deposit(24000)
	account.Deposit(-1000)
	account.Withdraw(58000)
	account.Withdraw(200000)

	fmt.Printf("Баланс счета %s: %.2f руб.\n", account.owner, account.Balance())
}
