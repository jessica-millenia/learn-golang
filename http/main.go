package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	bs := make([]byte, 9999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))

	fmt.Println("\n using a custom writer")
	resp2, err2 := http.Get("http://udemy.com")
	if err2 != nil {
		fmt.Println("Error:", err2)
		os.Exit(1)
	}
	lw := logWriter{}
	io.Copy(lw, resp2.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
