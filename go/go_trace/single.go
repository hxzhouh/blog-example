package main

import (
	"encoding/xml"
	"io"
	"log"
	"os"
	"strings"
)

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
