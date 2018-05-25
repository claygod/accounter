package accounter

// Accounter
// Batcher bench
// Copyright © 2017 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	//"sync"
	"testing"
)

/*
func BenchmarkBatchSyncMapStore(b *testing.B) {
	b.StopTimer()
	var inbox sync.Map
	tr := &Transaction{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		inbox.Store(i, tr)
	}
}
func BenchmarkBatchMapStore(b *testing.B) {
	b.StopTimer()
	inbox := make(map[int]interface{})
	tr := &Transaction{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		inbox[i] = tr
	}
}
*/
/*
func BenchmarkBatchIteration(b *testing.B) {
	b.StopTimer()
	in := newQueue()
	out := newQueue()
	for i := int64(0); i < 1000; i++ {
		t := &Transaction{tid: i}
		t.Credit(i, "USD", 1)
		t.Debit(i+1, "USD", 1)
		in.PushTail(t)
	}

	t := &Transaction{tid: 11}
	t.Credit(11, "USD", 1)
	t.Debit(12, "USD", 1)
	in.PushTail(t)

	nb := newBatcher(in, out, batchSize)
	deferred := make([]*Transaction, 0, 20)
	ne := newExecutor(out)

	b.StartTimer()
	for i := 0; i < b.N; i = i + 2 {
		for u := int64(0); u < batchSize; u++ {
			in.PushTail(t)
		}
		nb.iteration(deferred, int(batchSize))
		ne.step()
	}
}
*/
func BenchmarkAccounterAa(b *testing.B) {
	b.StopTimer()
	a := New(32)

	b.StartTimer()
	for i := 0; i < b.N; i = i + 2 {
		a.Begin().Credit(int64(i), "USD", 1).Debit(int64(i+1), "USD", 1).End()
	}
}
