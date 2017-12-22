package accounter

// Indicator
// API
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	//"errors"
	//"log"
	"runtime"
	"sync/atomic"
)

const trialLimitConst int = 2147483647
const trialStop int = 64
const closed int64 = -2147483647

var trialLimit int = trialLimitConst
var batchSize int = 4

// Indicator is a counter that registers the number of requests for work with
// any resource. Each resource receives permission and at the end of the work
// is obliged to notify of this. With the stop command, the library will not
// allow the new process to work with the protected resource. The library is
// designed for multi-threaded use and uses non-blocking algorithms.
type Indicator struct {
	counter int64
}

// New - create new Indicator
func New() Indicator {
	return Indicator{}
}

// Catch - get permission to access the resource
func (i *Indicator) Catch() bool {
	if atomic.LoadInt64(&i.counter) < 0 {
		return false
	}
	if atomic.AddInt64(&i.counter, 1) > 0 {
		return true
	}
	atomic.AddInt64(&i.counter, -1)
	return false
}

// Throw - resource is no longer used
func (i *Indicator) Throw() {
	atomic.AddInt64(&i.counter, -1)
}

// Start - run resource usage
func (i *Indicator) Start() bool {
	var currentCounter int64
	for u := trialLimit; u > trialStop; u-- {
		currentCounter = atomic.LoadInt64(&i.counter)
		if currentCounter >= 0 {
			return true
		}
		// the variable `currentCounter` is expected to be `permitError`
		if atomic.CompareAndSwapInt64(&i.counter, closed, 0) {
			return true
		}
		runtime.Gosched()
	}
	return false
}

// Stop - to stop using the resource.
// Returns 0 if successful. This means that all the processes
// that took permission to work with the resource are completed.
// Otherwise, it returns the number of processes that have received
// permission to work with resources and did not release it.
func (i *Indicator) Stop() int64 {
	var currentCounter int64

	for u := trialLimit; u > trialStop; u-- {
		currentCounter = atomic.LoadInt64(&i.counter)
		switch {
		case currentCounter == 0:
			if atomic.CompareAndSwapInt64(&i.counter, 0, closed) {
				return 0
			}
		case currentCounter > 0:
			atomic.CompareAndSwapInt64(&i.counter, currentCounter, currentCounter+closed)
		case currentCounter == closed:
			return 0
		}
		runtime.Gosched()
	}
	currentCounter = atomic.LoadInt64(&i.counter)
	if currentCounter < 0 && currentCounter > closed {
		atomic.AddInt64(&i.counter, -closed)
	}
	return currentCounter - closed
}

// StopUnsafe - forced stop and reset the counter.
func (i *Indicator) StopUnsafe() {
	atomic.StoreInt64(&i.counter, closed)
	return
}

// Status - shows the current status of the counter
func (i *Indicator) Status() int64 {
	return atomic.LoadInt64(&i.counter)
}
