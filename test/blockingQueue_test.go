package test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	blockingQueue "github.com/siawase7179/go_blockingqueue"
)

func TestQueueFull(t *testing.T) {
	var capacity uint64
	capacity = 10

	queue, err := blockingQueue.NewBlockingQueue(capacity)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("capacity : %d", queue.Capacity())

	for i := uint64(0); i < capacity+1; i++ {
		_, err := queue.Push(i)
		if err != nil {
			t.Log(err.Error())
		} else {

		}
	}
}

func TestQueueEmpty(t *testing.T) {
	var capacity uint64
	capacity = 5

	queue, err := blockingQueue.NewBlockingQueue(capacity)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("capacity : %d", queue.Capacity())

	for i := uint64(0); i < capacity; i++ {
		_, err := queue.Pop()
		if err != nil {
			t.Log(err.Error())
		} else {

		}
	}
}

func TestQueue(t *testing.T) {
	var capacity uint64
	capacity = 100
	queue, err := blockingQueue.NewBlockingQueue(capacity)
	if err != nil {
		t.Fatal(err.Error())
	}

	go func() {
		t.Logf("capacity:%d\n", queue.Capacity())

		for i := uint64(0); i < capacity; i++ {
			_, err := queue.Push(i)
			if err != nil {

			}
		}
	}()

	var groupCount int
	groupCount = 5
	wg := new(sync.WaitGroup)
	wg.Add(groupCount)

	for i := 0; i < groupCount; i++ {
		go func(n int) {
			for queue.IsEmpty() != true {
				res, err := queue.Pop()
				if err == nil {
					t.Logf("pop[%d] : %d\n", n+1, res)
					time.Sleep(1 * time.Second)
				} else {

				}
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("done...")

}
