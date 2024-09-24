package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	port    string
	workDir string
)

func init() {
	flag.StringVar(&port, "port", "8080", "Port to run the HTTP server on")
	flag.StringVar(&workDir, "workDir", ".", "Working directory")
	flag.Parse()
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/files", filesHandler)
	http.HandleFunc("/download", downloadHandler)

	fmt.Printf("Starting server on port %s with workDir %s\n", port, workDir)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		os.Exit(1)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func filesHandler(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir(workDir)
	if err != nil {
		http.Error(w, "Failed to read directory", http.StatusInternalServerError)
		return
	}

	for _, file := range files {
		fmt.Fprintf(w, "%s\n", file.Name())
	}
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("file")
	newFileName := r.URL.Query().Get("rename")
	if fileName == "" {
		http.Error(w, "File name is required", http.StatusBadRequest)
		return
	}

	if newFileName == "" {
		newFileName = fileName
	}

	filePath := workDir + "/" + fileName
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Disposition", "attachment; filename="+newFileName)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", getFileSize(filePath)))

	if _, err := io.Copy(w, file); err != nil {
		http.Error(w, "Failed to send file", http.StatusInternalServerError)
	}
}

func getFileSize(filePath string) int64 {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0
	}
	return fileInfo.Size()
}
