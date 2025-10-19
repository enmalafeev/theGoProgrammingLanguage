package main

import (
	"goModule/ch1/lissajous"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	cycles, _ := strconv.Atoi(params.Get("count"))
	lissajous.Lissajous(w, cycles)
}
