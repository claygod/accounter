package accounter

// Accounter
// Batch
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

// newBatcher - create new batch
func newBatcher(inbox *inbox, size int) *batcher {
	b := &batcher{
		inbox: inbox,
		size:  size,
	}
	return b
}

// batcher is a  ...
type batcher struct {
	inbox *inbox
	//outbox     *inboxTransaction
	size int
	stop bool
}

// start of a batcher
func (b *batcher) start() {
	for {
		if b.stop {
			return
		}
		nb := newBatch()
		for i := 0; i < b.size; i++ {
			num, t := b.inbox.Take()
			if num == 0 {
				break
			}
			nb[num] = t
		}
		// run banch
	}
}

// newBatch - create new batch
func newBatch() batch {
	b := make(map[int64]interface{})
	return b
}

// batch is a  ...
type batch map[int64]interface{}
