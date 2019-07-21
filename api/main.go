// cmd : REDIS_HOST=localhost REDIS_PORT=6379 go run main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	util "github.com/verlandz/docker/api/util"
	redis "gopkg.in/redis.v5"
)

var (
	port        = ":8082"
	redisKey    = "user"
	redisClient *redis.Client
)

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:       util.StringConcat(os.Getenv("REDIS_HOST"), ":", os.Getenv("REDIS_PORT")),
		Password:   "",
		DB:         0,
		MaxRetries: 5,
	})
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "[API] Service is running")
}

func data(w http.ResponseWriter, r *http.Request) {
	arr, err := redisClient.HGetAll(redisKey).Result()
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Failed to connect to Redis")
		return
	}

	users := []user{}
	for id, name := range arr {
		temp := user{
			ID:   id,
			Name: name,
		}
		users = append(users, temp)
	}

	b, _ := json.Marshal(users)
	fmt.Fprintf(w, string(b))
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/data", data)

	fmt.Println("REDIS_HOST :", os.Getenv("REDIS_HOST"))
	fmt.Println("REDIS_PORT :", os.Getenv("REDIS_PORT"))

	fmt.Println("Running in", runtime.Version())
	fmt.Println("Listen and serve", port)
	err := http.ListenAndServe(port, nil)
	fmt.Println(err)
}
