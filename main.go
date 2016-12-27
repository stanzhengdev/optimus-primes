package optimusprime

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"google.golang.org/appengine"

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
// Can handle first 1,000,000 primes in sequence, count is not by value
// but by `start` and `end` index
func PrimeHandler(w http.ResponseWriter, r *http.Request) {
	// TODO Parse from multiple data files
	var count, start, end, rangeStart int
	var err error
	var resp []byte
	query := r.URL.Query()
	c := query.Get("count")
	e := query.Get("end")
	s := query.Get("start")
	rs := query.Get("rangeStart")
	var errs []string
	if c == "" {
		count = 100
	} else {
		count, err = strconv.Atoi(c)
		if err != nil {
			errs = append(errs, err.Error())
		}
	}

	if s == "" {
		start = 0
	} else {
		start, err = strconv.Atoi(s)
		if err != nil {
			errs = append(errs, err.Error())
		}
	}

	if rs == "" {
		rangeStart = 0
	} else {
		rangeStart, err = strconv.Atoi(rs)
		if err != nil {
			errs = append(errs, err.Error())
		}
	}
	check := func() bool {
		for _, i := range errs {
			if i != "" {
				fmt.Println(i)
				return false
			}
		}
		return true
	}()
	fmt.Println("check")
	if false && check {

		if e == "" {
			end = 0
		} else {
			end, _ = strconv.Atoi(e)
		}
		nums := fileOpen(1, count+rangeStart)[start+rangeStart : end+rangeStart]
		resp, _ = json.Marshal(nums)
	} else {
		resp, err = json.Marshal(errs)
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func init() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/api/v1", PrimeHandler)

	http.Handle("/", r)
}

func main() {
	appengine.Main()
}
