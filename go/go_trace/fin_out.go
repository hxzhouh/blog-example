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

func concurrent(docs []string) int {
	var count int32
	wg := &sync.WaitGroup{}
	for _, v := range docs {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()
			var iFound int32
			f, err := os.OpenFile(v, os.O_RDONLY, 0)
			if err != nil {
				return
			}
			defer func() {
				atomic.AddInt32(&count, iFound)
				_ = f.Close()
			}()
			data, err := io.ReadAll(f)
			if err != nil {
				return
			}
			var d document
			if err = xml.Unmarshal(data, &d); err != nil {
				log.Printf("Decoding Document [Ns] : ERROR :%+v", err)
				return
			}
			for _, item := range d.Channel.Items {
				if strings.Contains(strings.ToLower(item.Title), "go") {
					iFound++
				}
			}
		}()
	}
	wg.Wait()
	return int(count)
}
