package accounter

// Accounter
// Executor
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"log"
	"runtime"
	"sync"
	"sync/atomic"
)

// newExecutor - create new batch
func newExecutor(chanExecutor chan *batch) *executor {
	e := &executor{
		chanExecutor: chanExecutor,
	}
	return e
}

// executor is a  ...
type executor struct {
	chanExecutor chan *batch
	stop         bool
}

// start of a executor
func (e *executor) start() {
	log.Print("----- start EXECUTOR --------")
	for {
		if e.stop {
			return
		}
		if !e.step() {
			runtime.Gosched()
		}
	}
}

// step
func (e *executor) step() bool {
	select {
	case b := <-e.chanExecutor:
		if b[0] != nil {
			t := b[0].(*Transaction)
			t.processed = 1
			t.tid = 555
			return true
		}
	default:
		return false
	}
	return false

	/*
		var wg sync.WaitGroup
		nb := b.(batch)
		for _, t := range nb {
			if t != nil {
				nt := t.(*Transaction)
				// nt.tid = int64(c) // 222
				wg.Add(1)
				// execute transaction
				// . . .
				go e.exe(nt, &wg)
			}
		}
		wg.Wait()
		return true
	*/
}

func (e *executor) exe(t *Transaction, wg *sync.WaitGroup) {
	// . . .
	atomic.StoreUint32(&t.processed, 1)
	wg.Done()
}
