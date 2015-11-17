package main

import (
	"fmt"
	"net/url"
)

func main() {
	querys := []string{"family", "happy"}
	params := url.Values{
		"etag": querys}
	fmt.Println(params.Encode())
}
