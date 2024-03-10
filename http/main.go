package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://example.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Print html #1
	//bs := make([]byte, 1255)
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))

	// Print html #2
	io.Copy(os.Stdout, resp.Body)
}
