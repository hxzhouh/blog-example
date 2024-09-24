package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"strings"
	"sync"
	"sync/atomic"
)

type document struct {
	Channel Channel `xml:"channel"`
}
type Channel struct {
	Channel       string `xml:"channel"`
	Items         []Item `xml:"item"`
	Title         string `xml:"title"`
	Link          string `xml:"link"`
	Description   string `xml:"description"`
	Generator     string `xml:"generator"`
	Language      string `xml:"language"`
	Copyright     string `xml:"copyright"`
	LastBuildDate string `xml:"lastBuildDate"`
}
type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	Author      string `xml:"author"`
	Guid        string `xml:"guid"`
	Description string `xml:"description"`
}

func freq(docs []string) int {
	var count int
	for _, doc := range docs {
		f, err := os.OpenFile(doc, os.O_RDONLY, 0)
		if err != nil {
			return 0
		}
		defer func() {
			_ = f.Close()
		}()
		data, err := io.ReadAll(f)
		if err != nil {
			return 0
		}
		var d document
		if err = xml.Unmarshal(data, &d); err != nil {
			log.Printf("Decoding Document [Ns] : ERROR :%+v", err)
			return 0
		}
		for _, item := range d.Channel.Items {
			if strings.Contains(strings.ToLower(item.Title), "go") {
				count++
			}
		}
	}
	return count
}

func concurrent(docs []string) int {
	var count int32
	g := runtime.GOMAXPROCS(0)
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
				defer func() {
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
			}
		}()
	}

	wg.Wait()
	return int(count)
}

var documentPool = sync.Pool{
	New: func() interface{} {
		return new(document)
	},
}

func (d *document) reset() {
	*d = document{}
}

func syncPool(docs []string) int {
	var count int32
	g := runtime.GOMAXPROCS(0)
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
			//var d *document
			for doc := range ch {
				f, err := os.OpenFile(doc, os.O_RDONLY, 0)
				if err != nil {
					return
				}
				defer func() {
					_ = f.Close()
				}()
				data, err := io.ReadAll(f)
				if err != nil {
					return
				}
				d := documentPool.Get().(*document)
				defer func() {
					d.reset()
					documentPool.Put(d)
				}()
				if err = xml.Unmarshal(data, &d); err != nil {
					log.Printf("Decoding Document [Ns] : ERROR :%+v", err)
					return
				}
				for _, item := range d.Channel.Items {

					if strings.Contains(strings.ToLower(item.Title), "go") {
						iFound++
					}
				}
			}
		}()
	}

	wg.Wait()
	return int(count)
}

// go run main go_trace/escape_test.go 2 > trace.out
// go tool trace trace.out
func main() {
	_ = trace.Start(os.Stdout)
	defer trace.Stop()
	files := make([]string, 0)
	for i := 0; i < 100; i++ {
		files = append(files, "index.xml")
	}
	//count := concurrent(files)
	count := syncPool(files)
	log.Println(fmt.Sprintf("find key word go %d count", count))
}
