package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := os.Getenv("NAME")
		planet := os.Getenv("PLANET")

		fmt.Fprintf(w, "<h1>Hello, World!</h1>\n <p>I am %s, from %s</p>", name, planet)
	})

	http.ListenAndServe(":80", nil)
}
