package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/env", func(w http.ResponseWriter, r *http.Request) {
		for _, env := range os.Environ() {
			pair := strings.SplitN(env, "=", 2)
			fmt.Fprintf(w, "%s=%s\n", pair[0], pair[1])
		}
		fmt.Fprintf(w, "Hi")
	})

	log.Println("starting http server")
	log.Fatal(http.ListenAndServe(":8118", nil))
}
