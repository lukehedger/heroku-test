package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":" + port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Heroku!")
}
