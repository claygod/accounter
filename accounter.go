package accounter

// Indicator
// API
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

//"errors"
//import "log"
import "time"
import "runtime"
import "sync/atomic"

const trialLimitConst int = 2147483647
const trialStop int = 64
const closed int64 = -2147483647
const maxBatchSize int64 = 32
const accStopped int64 = 4611686018427387904

var trialLimit int = trialLimitConst
var batchSize int64 = 32
var typicalTransactSize int = 4

// New - create new Accounter
func New(batchSize int64) *Accounter {
	chanBatcher := make(chan *Transaction, 4000)
	chanExecutor := make(chan *batch, 1000)
	a := &Accounter{
		batcher:      newBatcher(batchSize, chanBatcher, chanExecutor),
		executor:     newExecutor(chanExecutor),
		chanBatcher:  chanBatcher,
		chanExecutor: chanExecutor,
	}
	//log.Print("----- New 1000 --------")
	go a.batcher.start()
	time.Sleep(time.Millisecond * 700)
	//log.Print("----- New 2000 --------")
	go a.executor.start()
	time.Sleep(time.Millisecond * 700)
	//log.Print("----- New 3000 --------")
	return a
}

// Accounter is a  ...
type Accounter struct {
	batcher      *batcher
	executor     *executor
	chanBatcher  chan *Transaction
	chanExecutor chan *batch
}

func (a *Accounter) Begin() *Transaction {
	return newTransaction(a)
}

func (a *Accounter) SetBatchSize(size int) *Transaction {
	return nil
}

func (a *Accounter) doTransaction(t *Transaction) *Transaction {
	a.chanBatcher <- t
	runtime.Gosched()
	for i := 0; i < 1000000; i++ { //
		if atomic.LoadUint32(&t.processed) != 0 {
			return t
		}
	}
	t.tid = 777
	return t
}
