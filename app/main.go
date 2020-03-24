package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "skaffold playground app!!\n")
}

func main() {
	log.Print("skaffold playground app server ready")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
