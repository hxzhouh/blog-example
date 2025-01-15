package main

import (
	"testing"
	"time"
)

func TestAddDriver(t *testing.T) {
	tw := NewTimeWheel(1, 10, 10)
	defer tw.Stop()

	uid := "driver1"
	tw.Add(uid)

	if _, exists := tw.drivers[uid]; !exists {
		t.Errorf("Driver %s not found in drivers map", uid)
	}

	found := false
	for _, slot := range tw.slots {
		for _, driver := range slot {
			if driver.uid == uid {
				found = true
				break
			}
		}
	}
	if !found {
		t.Errorf("Driver %s not found in slots", uid)
	}
}
func TestTimeWheel(t *testing.T) {
	tw := NewTimeWheel(1, 10, 10)
	defer tw.Stop()

	uid := "driver1"
	tw.Add(uid)

	uid2 := "driver2"
	tw.Add(uid2)
	time.Sleep(5 * time.Second)
	uid3 := "driver3"
	tw.Add(uid3)
	time.Sleep(time.Second * 12)
}
