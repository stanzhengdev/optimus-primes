package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func parsePrimes(start, end int) {

}

func fileOpen(f, limit int) (lines string) {
	fname := fmt.Sprintf("data/primes%d.txt", f)
	count := 0
	strip := 2
	file, err := os.Open(fname) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Trim the first few lines
		if count < strip {
			count++
			continue
		}
		line := scanner.Text()
		fmt.Println(line)
		lines += fmt.Sprintf("%d: line %s", count-2, line)
		count++
		if count > limit {
			break
		}
	}
	return
}

// PrimeHandler parses request and serves back a range
func PrimeHandler(w http.ResponseWriter, r *http.Request) {
	// files, _ := ioutil.ReadDir("data/")
	// for _, f := range files {
	// 	w.Write([]byte(f.Name()))
	// }
	w.Write([]byte(fileOpen(1, 100)))
}

func init() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/api/v1", PrimeHandler)

	http.Handle("/", r)
}
