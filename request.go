package accounter

// Accounter
// Transaction request
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

//"errors"
//"log"
//"fmt"

type Transaction struct {
	processed uint32
	result    bool
	tid       int64
	//down      []*Request
	//up        []*Request
	//reqs      map[uint64]*Request
	requests  []*Request
	accounter *Accounter
}

// newBatcher - create new batch
func newTransaction(a *Accounter) *Transaction {
	t := &Transaction{
		//down:      make([]*Request, 0, typicalTransactSize),
		//up:        make([]*Request, 0, typicalTransactSize),
		//reqs:      make(map[uint64]*Request),
		requests:  make([]*Request, 0, typicalTransactSize),
		accounter: a,
	}
	return t
}

const multiplier uint64 = 1024

func (t *Transaction) Debit(customer int64, account string, count int64) *Transaction {
	//t.up = append(t.up, &Request{id: customer, key: account, amount: count, hash: t.hash(customer, account)})
	t.requests = append(t.requests, &Request{id: customer, key: account, amount: count, hash: t.hash(customer, account)})
	return t
}

func (t *Transaction) Credit(customer int64, account string, count int64) *Transaction {
	//t.down = append(t.down, &Request{id: customer, key: account, amount: count, hash: t.hash(customer, account)})
	t.requests = append(t.requests, &Request{id: customer, key: account, amount: -(count), hash: t.hash(customer, account)})
	return t
}

func (t *Transaction) End() *Transaction {
	return t.accounter.doTransaction(t)
}

func (t *Transaction) hash(id int64, str string) uint64 {
	out := uint64(id)
	ln := uint64(len(str) - 1)
	if ln > 1023 || ln < 0 {
		return 0
	}
	for i := uint64(0); i < ln; i++ {
		out += uint64(str[i]) + multiplier*i
	}
	return out
}

type Request struct {
	id     int64
	key    string
	amount int64
	hash   uint64
}
