package accounter

// Accounter
// Inbox for transactions, batches
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	//"errors"
	//"log"
	//"runtime"
	"sync"
	"sync/atomic"
)

// inbox is a  ...
type inbox struct {
	counterIn  int64
	counterOut int64
	box        sync.Map
}

// newInbox - create new inbox
func newInbox() inbox {
	i := inbox{counterIn: 1, counterOut: 1}
	i.box.Store(0, &Transaction{})
	return i
}

// Add a new object to the box
func (i *inbox) Add(t interface{}) int64 {
	counter := atomic.AddInt64(&i.counterIn, 1)
	i.box.Store(counter, t)
	return counter
}

// Take the next object from the box
func (i *inbox) Take() (int64, interface{}) {
	counterIn := atomic.LoadInt64(&i.counterIn)
	counterOut := i.counterOut
	if counterIn == counterOut {
		return 0, nil
	}
	t, _ := i.box.Load(counterOut)
	i.box.Delete(counterOut)
	i.counterOut++
	return counterOut, t

}
