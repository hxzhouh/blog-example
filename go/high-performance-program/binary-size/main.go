package main

import (
	"fmt"
	"net/http"
)

// main

func main() {
	// create a http server and create a handler hello, return hello world
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World\n")
	})
	// listen to port 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

/**
➜  binary-size git:(main) ✗ go build -ldflags="-s -w"  -o server main.go && upx --brute server && ls -lh server
                       Ultimate Packer for eXecutables
                          Copyright (C) 1996 - 2023
UPX 4.1.0       Markus Oberhumer, Laszlo Molnar & John Reiser    Aug 8th 2023

        File size         Ratio      Format      Name
   --------------------   ------   -----------   -----------
   4693922 ->   1441808   30.72%   macho/arm64   server

Packed 1 file.
-rwxr-xr-x  1 hxzhouh  staff   1.4M Jul  2 14:54 server

----
➜  binary-size git:(main) ✗ go build  -o server main.go  && ls -lh server
-rwxr-xr-x  1 hxzhouh  staff   6.5M Jul  2 16:43 server
*/
