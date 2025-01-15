package main

import (
	"log"
	"sync"
	"time"
)

type Driver struct {
	uid      string
	expireAt int64
}

type TimeWheel struct {
	drivers     map[string]*Driver // all drivers
	slots       [][]*Driver        // time wheel solts
	position    int                // current position
	mu          sync.Mutex         //
	ticker      *time.Ticker       //
	stop        chan struct{}      //
	interval    int
	slotCount   int
	timeoutSecs int
}

func NewTimeWheel(interval, slotCount, timeoutSecs int) *TimeWheel {
	tw := &TimeWheel{
		interval:    interval,
		slotCount:   slotCount,
		timeoutSecs: timeoutSecs,
		drivers:     make(map[string]*Driver),
		slots:       make([][]*Driver, slotCount),
		position:    0,
		ticker:      time.NewTicker(time.Duration(interval) * time.Second),
		stop:        make(chan struct{}),
	}

	for i := 0; i < slotCount; i++ {
		tw.slots[i] = make([]*Driver, 0)
	}
	go tw.run()
	return tw
}

func (tw *TimeWheel) Add(uid string) {
	tw.mu.Lock()
	defer tw.mu.Unlock()

	expireAt := time.Now().Unix() + int64(tw.timeoutSecs)
	driver := &Driver{
		uid:      uid,
		expireAt: expireAt,
	}
	tw.drivers[uid] = driver
	slot := tw.GetSlot(-1)
	tw.slots[slot] = append(tw.slots[slot], driver)
	log.Printf("time:%d,Driver %s added to slot %d\n", time.Now().Unix(), uid, slot)
}

func (tw *TimeWheel) GetSlot(index int) int {
	return (tw.position + index + tw.slotCount) % tw.slotCount
}

func (tw *TimeWheel) run() {
	for {
		select {
		case <-tw.ticker.C:
			log.Printf("tick, position:%d\n", tw.position)
			tw.mu.Lock()
			expired := tw.slots[tw.position]
			tw.slots[tw.position] = make([]*Driver, 0)
			tw.position = (tw.position + 1) % tw.slotCount
			tw.mu.Unlock()
			for _, driver := range expired {
				if time.Now().Unix() >= driver.expireAt {
					tw.mu.Lock()
					delete(tw.drivers, driver.uid)
					tw.mu.Unlock()
					log.Printf("Driver %s timeout\n", driver.uid)
				}
			}
		case <-tw.stop:
			return
		}
	}
}

func (tw *TimeWheel) Stop() {
	close(tw.stop)
	tw.ticker.Stop()
}
