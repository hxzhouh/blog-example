package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
)

//func NewHttpService() {
//	go func() {
//		log.Println(http.ListenAndServe("localhost:6060", nil))
//	}()
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		w.Write([]byte("Hello, world!"))
//	})
//	err := http.ListenAndServe(":8081", nil)
//	if err != nil {
//		log.Println("http.ListenAndServe err:", err)
//	}
//}
//
//func main() {
//	go NewHttpService()
//	time.Sleep(1 * time.Second)
//	for { // 模拟大量请求
//		go makeRequest()
//		time.Sleep(50 * time.Millisecond)
//	}
//}

func makeRequest() {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8081", nil)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	_, err = ioutil.ReadAll(res.Body)
	// defer res.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
}
