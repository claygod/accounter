package accounter

// Accounter
// Batcher test
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	// "sync/atomic"
	"testing"
	"time"
)

func TestAccounterAaa(t *testing.T) {
	a := New(32)
	time.Sleep(time.Millisecond * 5)
	tr := a.Begin().Credit(int64(1), "USD", 1).Debit(int64(2), "USD", 1).End()
	time.Sleep(time.Millisecond * 5)
	if tr.tid == 777 {
		t.Error(tr)
	}
}
