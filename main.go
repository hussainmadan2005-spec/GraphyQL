package main

import (
    "log"
    "net"
    "net/http"
)

func main() {
    listener, err := net.Listen("tcp", "127.0.0.1:0")
    if err != nil {
        log.Fatal(err)
    }
    defer listener.Close()

    port := listener.Addr().(*net.TCPAddr).Port
    fs := http.FileServer(http.Dir("./static"))

    http.Handle("/", fs)

    log.Printf("Server running on http://localhost:%d", port)

    if err := http.Serve(listener, nil); err != nil {
        log.Fatal(err)
    }
}
