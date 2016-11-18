package main

import "net/http"
import "fmt"

func main() {
	client := &http.Client{}

	// resp, err := client.Get("http://example.com")
	// ...

	req, _ := http.NewRequest("GET", "http://google.com", nil)
	// ...
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, _ := client.Do(req)
	fmt.Println(resp)
	// ...
}
