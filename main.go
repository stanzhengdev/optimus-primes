package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func parsePrimes(start, end int) {

}

// Removes filterEmpty
func filterEmpty(v string) bool {
	return v != ""
}

func fileOpen(f, limit int) (lines []string) {
	fname := fmt.Sprintf("data/primes%d.txt", f)
	strip := 2
	count := 0
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
	var count, start, end, rangeStart int
	query := r.URL.Query()
	c := query.Get("count")
	s := query.Get("start")
	rs := query.Get("rangeStart")
	e := query.Get("end")
	if c == "" {
		count = 100
	} else {
		count, _ = strconv.Atoi(c)
	}

	if s == "" {
		start = 0
	} else {
		start, _ = strconv.Atoi(s)
	}

	if rs == "" {
		rangeStart = 0
	} else {
		rangeStart, _ = strconv.Atoi(rs)
	}

	if e == "" {
		end = 0
	} else {
		end, _ = strconv.Atoi(e)
	}
	nums := fileOpen(1, count+rangeStart)[start+rangeStart : end+rangeStart]
	resp, _ := json.Marshal(nums)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func init() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/api/v1", PrimeHandler)

	http.Handle("/", r)
}
