package accounts

import (
	"errors"
	"fmt"
)

/*
* account struct
소문자로 시작한다면 private이다.
*/
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// receiver
// 명명규칙: struct의 가장 첫 문자의 소문자 Account -> a

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

var errNoMoney = errors.New("can't withdraw")

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	// nil : null
	if a.balance < amount {
		//return errors.New("Can't withdraw you ar poor")
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	//return "whatever you want"
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
