package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting K3s Glaucus!")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprintf(w, "K3S on Glaucus! With CI/CD! Fancy! "+os.Getenv("SECRET"))
	})
	log.Fatal(http.ListenAndServe(":80", nil))
}
