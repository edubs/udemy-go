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
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// empty 'storage' byte slice
	//bs := make([]byte, 99999)
	// read body data into byte slice
	//resp.Body.Read(bs)

	// condensed way
	//io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

// we're implementing the Writer interface by declaring a function
// that takes an argument of type byte string and returns an int and an error
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil

}
