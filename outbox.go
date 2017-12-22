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

// outbox is a  ...
type outbox struct {
	//counterIn  int64
	//counterOut int64
	box sync.Map
}

// newOutbox - create new outbox
func newOutbox() *outbox {
	o := &outbox{}
	return o
}

// Add a new transaction to the box
func (o *outbox) Add(counter int64, t *Transaction) {
	o.box.Store(counter, t)
	return
}

// Take the next transaction from the box
func (o *outbox) Take(num int64) *Transaction {
	for {
		if r, ok := o.box.Load(num); ok {
			o.box.Delete(num)
			return r.(*Transaction)
		}
		runtime.Gosched()
	}
}
