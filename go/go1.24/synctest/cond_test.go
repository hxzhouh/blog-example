package main

import (
	"sync"
	"testing"
	"time"
)

type Queue struct {
	items []int
	cond  *sync.Cond
}

func NewQueue() *Queue {
	return &Queue{
		items: []int{},
		cond:  sync.NewCond(&sync.Mutex{}),
	}
}

// AddItem adds an item to the queue and notifies one waiting goroutine.
func (q *Queue) AddItem(item int) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	q.items = append(q.items, item)
	q.cond.Signal() // Notify one waiting goroutine
}

// GetItem retrieves an item from the queue, waiting if necessary.
func (q *Queue) GetItem() int {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	for len(q.items) == 0 {
		q.cond.Wait() // Wait until notified
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func TestQueueWithSyncCond(t *testing.T) {
	synctest.Run(func() {
		queue := NewQueue()
		// 启动多个消费者
		var wg sync.WaitGroup
		wg.Add(2)
		for i := 0; i < 2; i++ {
			go func(id int) {
				defer wg.Done()
				item := queue.GetItem()
				if item != id {
					t.Errorf("Consumer %d got item %d, want %d", id, item, id)
				}
			}(i)
		}

		// 启动生产者
		time.Sleep(1 * time.Second) // 等待消费者阻塞
		for i := 0; i < 2; i++ {
			queue.AddItem(i)
		}
		// 等待所有消费者完成
		synctest.Wait() // 等待所有 `bubble` 内的 goroutine 达到 durably blocked 状态
		wg.Wait()
	})
}
