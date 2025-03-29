package main

import (
	"encoding/xml"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"sync/atomic"
)

func syncPool(docs []string) int {
	var count int32
	g := 10
	wg := sync.WaitGroup{}
	wg.Add(g)
	ch := make(chan string, 100)
	go func() {
		for _, v := range docs {
			ch <- v
		}
		close(ch)
	}()

	for i := 0; i < g; i++ {
		go func() {
			var iFound int32
			defer func() {
				atomic.AddInt32(&count, iFound)
				wg.Done()
			}()
			for doc := range ch {
				f, err := os.OpenFile(doc, os.O_RDONLY, 0)
				if err != nil {
					return
				}
				data, err := io.ReadAll(f)
				if err != nil {
					return
				}
				_ = f.Close()
				d := documentPool.Get().(*document)
				if err = xml.Unmarshal(data, &d); err != nil {
					log.Printf("Decoding Document [Ns] : ERROR :%+v", err)
					return
				}
				for _, item := range d.Channel.Items {
					if strings.Contains(strings.ToLower(item.Title), "go") {
						iFound++
					}
				}
				d.reset()
				documentPool.Put(d)
			}
		}()
	}

	wg.Wait()
	return int(count)
}
