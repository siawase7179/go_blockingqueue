package blockingQueue

import (
	"container/list"
	"errors"
	"math"
	"sync"
)

var ErrorCapacity = errors.New("Error Capacity")
var ErrorFull = errors.New("Queue is Full")
var ErrorEmpty = errors.New("Queue is Empty")

type BlockingQueue struct {
	count      uint64
	lock       *sync.Mutex
	writeIndex uint64
	readIndex  uint64
	store      *list.List
	capacity   uint64
}

func (q *BlockingQueue) inc(idx uint64) uint64 {
	if idx >= math.MaxUint64 {
		panic("Overflow")
	}

	if 1+idx == q.capacity {
		return 0
	} else {
		return idx + 1
	}
}

func NewBlockingQueue(capacity uint64) (*BlockingQueue, error) {
	if capacity < 1 {
		return nil, ErrorCapacity
	}

	lock := new(sync.Mutex)
	return &BlockingQueue{
		lock:     lock,
		count:    uint64(0),
		store:    list.New(),
		capacity: capacity,
	}, nil
}

func (q *BlockingQueue) Size() uint64 {
	q.lock.Lock()
	res := q.count
	q.lock.Unlock()

	return res
}

func (q *BlockingQueue) Capacity() uint64 {
	q.lock.Lock()
	res := uint64(q.capacity - q.count)
	q.lock.Unlock()

	return res
}

func (q *BlockingQueue) Push(item interface{}) (res bool, err error) {
	if item == nil {
		panic("Null item")
	}

	q.lock.Lock()
	if q.count == q.capacity {
		res, err = false, ErrorFull
	} else {

		q.store.PushBack(item)
		q.writeIndex = q.inc(q.writeIndex)
		q.count += 1

		res, err = true, nil
	}

	q.lock.Unlock()

	return
}

func (q *BlockingQueue) Pop() (res interface{}, err error) {
	q.lock.Lock()
	if q.count == 0 {
		res, err = nil, ErrorEmpty
	} else {
		var item = q.store.Remove(q.store.Front())
		q.readIndex = q.inc(q.readIndex)
		q.count -= 1

		res, err = item, nil
	}
	q.lock.Unlock()

	return
}

func (q BlockingQueue) IsEmpty() bool {
	return q.Size() == 0
}
