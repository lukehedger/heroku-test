package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
)

func main() {
	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatal(err)
	}

	client := redis.NewClient(opt)

	err = client.Set("key", "value", 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("key", val)

	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Heroku!")
}
