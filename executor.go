package accounter

// Accounter
// Executor
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"runtime"
)

// newExecutor - create new batch
func newExecutor(inbox *inbox, batchSize int) *executor {
	e := &executor{
		inbox: inbox,
		//batchSize: batchSize,
	}
	return e
}

// executor is a  ...
type executor struct {
	inbox *inbox
	stop  bool
}

// start of a executor
func (e *executor) start() {
	for {
		if e.stop {
			return
		}

		counter, b := e.inbox.Take()
		if counter == 0 {
			runtime.Gosched()
			break
		}
		nb := b.(batch)
		for c, t := range nb {
			nt := t.(*Transaction)
			nt.code = c // 222
		}
	}
}
