package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := os.Getenv("NAME")
		planet := os.Getenv("PLANET")

		fmt.Fprintf(w, "<h1>Hello, World!</h1>\n <p>I am %s, from %s</p>", name, planet)
	})

	http.HandleFunc("/secret", Secret)
	http.HandleFunc("/healthz", Healthz)

	http.ListenAndServe(":80", nil)
}

func Secret(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	fmt.Fprintf(w, "<h1>Secret</h1>\n <p>User: %s, Password: %s</p>", user, password)
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	duration := time.Since(startedAt)
	if duration.Seconds() < 10 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
	} else {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}
}
