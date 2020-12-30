// cmd: REDIS_HOST=127.0.0.1 REDIS_PORT=6379 go run main.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"

	util "github.com/verlandz/docker/api/util"
	"gopkg.in/redis.v5"
)

const (
	APP_PORT  = ":8080"
	REDIS_KEY = "docker-test"
)

var redisClient *redis.Client

type ApiResp struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:       util.StringConcat(os.Getenv("REDIS_HOST"), ":", os.Getenv("REDIS_PORT")),
		Password:   "",
		DB:         0,
		MaxRetries: 5,
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	var errResp string
	result, err := redisClient.Get(REDIS_KEY).Result()
	if err != nil {
		errResp = err.Error()
	}

	resp := ApiResp{
		Error: errResp,
		Data:  result,
	}
	bResp, _ := json.Marshal(resp)
	fmt.Fprintf(w, string(bResp))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println(
		"REDIS_HOST:", os.Getenv("REDIS_HOST"),
		"\nREDIS_PORT:", os.Getenv("REDIS_PORT"),
		"\nRunning in", runtime.Version(),
		"\nListen and serve", APP_PORT,
	)
	fmt.Println(http.ListenAndServe(APP_PORT, nil))
}
