package accounter

// Accounter
// Account
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

type Account struct {
	counter int64
	balance int64
}

// newAccount - create new account.
func newAccount(amount int64) *Account {
	k := &Account{balance: amount}
	return k
}

func (a *Account) addition(amount int64) int64 {
	nb := a.balance + amount
	if nb >= 0 && amount > -accStopped && amount < accStopped {
		a.balance = nb
		return nb
	}
	return nb + accStopped
}

/*
func (a *Account) credit(amount int64) int64 {
	nb := a.balance - amount
	if nb >= 0 {
		a.balance = nb
	}
	return nb + accStopped
}

func (a *Account) debit(amount int64) int64 {
	if a.balance >= 0 {
		a.balance += amount
	}
	return a.balance
}
*/
func (a *Account) total() int64 {
	if a.balance < 0 {
		return a.balance + accStopped
	}
	return a.balance
}

func (a *Account) start() {
	if a.balance < 0 {
		a.balance *= -1
		a.balance += accStopped
	}
}

func (a *Account) stop() {
	if a.balance >= 0 {
		a.balance *= -1
		a.balance -= accStopped
	}
}
