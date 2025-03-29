package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"sync"
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

var documentPool = sync.Pool{
	New: func() interface{} {
		return new(document)
	},
}

func (d *document) reset() {
	*d = document{}
}

// go run main go_trace/escape_test.go 2 > trace.out
// go tool trace trace.out
func main() {
	_ = trace.Start(os.Stdout)
	defer trace.Stop()
	files := make([]string, 0)
	for i := 0; i < 200; i++ {
		files = append(files, "index.xml")
	}
	//count := freq(files)
	//count := concurrent(files)
	count := syncPool(files)
	log.Println(fmt.Sprintf("find key word go %d count", count))
}
