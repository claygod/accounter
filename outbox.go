package accounter

// Accounter
// Outbox for transactions
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	//"errors"
	//"log"
	"runtime"
	"sync"
	//"sync/atomic"
)

// outboxTransaction is a  ...
type outboxTransaction struct {
	//counterIn  int64
	//counterOut int64
	box sync.Map
}

// newOutboxTransaction - create new outboxTransaction
func newOutboxTransaction() outboxTransaction {
	o := outboxTransaction{}
	return o
}

// Add a new transaction to the box
func (o *outboxTransaction) Add(counter int64, t *Transaction) {
	o.box.Store(counter, t)
	return
}

// Take the next transaction from the box
func (o *outboxTransaction) Take(num int64) *Transaction {
	for {
		if r, ok := o.box.Load(num); ok {
			o.box.Delete(num)
			return r.(*Transaction)
		}
		runtime.Gosched()
	}
	return nil
}
