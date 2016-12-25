package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

func parsePrimes(start, end int) {

}

func filterEmpty(v string) bool {
	return v != ""
}

func fileOpen(f, limit int) (lines []string) {
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
		arr := Filter(strings.Split(strings.TrimSpace(line), " "), filterEmpty)
		lines = append(lines, arr...)
		count++
		if count > limit {
			break
		}
	}
	return
}

// PrimeHandler parses request and serves back a range
func PrimeHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte(fileOpen(1, 100)))
	resp, _ := json.Marshal(fileOpen(1, 100))
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func init() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/api/v1", PrimeHandler)

	http.Handle("/", r)
}
