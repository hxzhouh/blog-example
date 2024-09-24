package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"runtime/trace"
	"strings"
	"sync"
	"sync/atomic"
	"time"
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

func freqPool(docs []string) int {
	var count int32
	g := runtime.GOMAXPROCS(0)
	wg := sync.WaitGroup{}
	wg.Add(g)
	ch := make(chan string, 100)
	go func() {
		for _, v := range docs {
			time.Sleep(1 * time.Millisecond)
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
				var d document
				if err := xml.Unmarshal(data, &d); err != nil {
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
	trace.Start(os.Stdout)
	defer trace.Stop()
	// download https://huizhou92.com/index.xml
	url := "https://huizhou92.com/index.xml"
	fileName := "index.xml"
	downLoadUrl(url, fileName)
	files := make([]string, 0)
	for i := 0; i < 100; i++ {
		files = append(files, fileName)
	}
	count := freqPool(files)
	log.Println(fmt.Sprintf("find key word go %d count", count))
}

func downLoadUrl(url, fileName string) {
	// URL of the XML file
	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error making the GET request:", err)
		return
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		log.Println("Error: received non-200 response code:", resp.StatusCode)
		return
	}

	if _, err = os.Stat(fileName); err == nil {
		// File exists, so delete it
		err = os.Remove(fileName)
		if err != nil {
			log.Println("Error deleting the existing file:", err)
			return
		}
	}

	// Create the file to save the XML content
	file, err := os.Create(fileName)
	if err != nil {
		log.Println("Error creating the file:", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Println("Error saving the content to the file:", err)
		return
	}
	log.Println("File downloaded successfully!")
}
