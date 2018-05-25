package accounter

// Accounter
// Batch
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import "log"
import "runtime"

// newBatcher - create new batch
func newBatcher(batchSize int64, chanBatcher chan *Transaction, chanExecutor chan *batch) *batcher {
	b := &batcher{
		batchSize:    batchSize,
		chanBatcher:  chanBatcher,
		chanExecutor: chanExecutor,
	}
	return b
}

// batcher is a  ...
type batcher struct {
	chanBatcher  chan *Transaction
	chanExecutor chan *batch
	batchSize    int64
	stop         bool
}

// start of a batcher
func (b *batcher) start() {
	deferred := make([]*Transaction, 0, 222)
	log.Print("----- start BATCHER --------")
	for {
		//log.Print("----- iterat --------")
		if b.stop {
			return
		}
		if !b.iteration(deferred, int(b.batchSize)) {
			//log.Print(res)
			runtime.Gosched()
		}
	}
}

// iteration
func (b *batcher) iteration(deferred []*Transaction, batchSize int) bool {
	var nb batch
	ns := 0 // new batch size
	lock := make(map[uint64]string)

	for i := 0; i < len(deferred) && i < batchSize; i++ {
		if b.addToBatch(&nb, deferred[i], lock, ns) {
			deferred = append(deferred[:i], deferred[i+1:]...)
			ns++
		}
	}
	for ; ns < batchSize; ns++ {
		select {
		case t := <-b.chanBatcher:
			//log.Print("итератор получил транзакцию")
			if !b.addToBatch(&nb, t, lock, ns) {
				deferred = append(deferred, t)
				ns--
			}
		default:
			break
		}
	}
	b.chanExecutor <- &nb
	if ns != batchSize-1 {
		return false
	}
	return true
}

func (b *batcher) addToBatch(ba *batch, tr *Transaction, lock map[uint64]string, ns int) bool {
	lockLocal := make(map[uint64]string)

	for _, req := range tr.requests {
		_, ok := lock[req.hash]
		_, ok2 := lockLocal[req.hash]
		if ok || ok2 {
			return false
		}
		lockLocal[req.hash] = req.key
	}

	for k, v := range lockLocal {
		lock[k] = v
	}
	(*ba)[ns] = tr
	return true
}

// batch is a  ...
type batch [maxBatchSize]interface{}
