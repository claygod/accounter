package accounter

// Accounter
// Batcher test
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	// "sync/atomic"
	"testing"
	"time"
)

/*
func TestBatcherXwe222(t *testing.T) {
	tr := &Transaction{tid: 111}
	tr.Credit(1, "USD", 1)
	tr.Debit(2, "USD", 1)

	nb := newBatcher(nil, nil, 4)
	//deferred := make([]*Transaction, 0, 22222)

	if nb.size == 150 {
		t.Error("Error ")
	}
}
*/
func TestAccounterAaa(t *testing.T) {
	a := New(32)
	time.Sleep(time.Millisecond * 5)
	tr := a.Begin().Credit(int64(1), "USD", 1).Debit(int64(2), "USD", 1).End()
	time.Sleep(time.Millisecond * 5)
	if tr.tid == 777 {
		t.Error(tr)
	}
}
