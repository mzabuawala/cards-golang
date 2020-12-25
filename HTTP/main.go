package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type bodyWritter struct{}

func main() {
	resp, err := http.Get("https://www.duckduckgo.com")
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}
	// Use in-built method to read the body data
	bw := bodyWritter{}
	io.Copy(bw, resp.Body)

	// Manually read the data
	// bs := make([]byte, 7000)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
}

// Emulate the Writer interface by our self
func (bw bodyWritter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Wrote that many bytes: ", len(bs))
	return len(bs), nil
}
