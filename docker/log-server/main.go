package main

import (
        "net/http"
)

func uploadFileHandler(responseWriter http.ResponseWriter, request *http.Request) {
        responseWriter.Write([]byte("Ebarle upload path"))
}

func main() {
        http.HandleFunc("/upload", uploadFileHandler)
        http.ListenAndServe("0.0.0.0:8080", nil)
}
