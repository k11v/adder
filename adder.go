package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func add(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "{\"status\": \"error\", \"message\": \"Method not allowed\"}\n")
		return
	}

	if err := req.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"status\": \"error\", \"message\": \"Invalid input\"}\n")
		return
	}

	aStr := req.Form.Get("a")
	bStr := req.Form.Get("b")

	a, err1 := strconv.Atoi(aStr)
	b, err2 := strconv.Atoi(bStr)

	if err1 != nil || err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"status\": \"error\", \"message\": \"Invalid input\"}\n")
		return
	}

	sum := a + b

	fmt.Fprintf(w, "{\"sum\": %d}\n", sum)
}

func health(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "{\"status\": \"error\", \"message\": \"Method not allowed\"}\n")
		return
	}

	fmt.Fprintf(w, "{\"status\": \"ok\"}\n")
}

func main() {
	http.HandleFunc("/add", add)
	http.HandleFunc("/health", health)

	http.ListenAndServe(":8000", nil)
}
