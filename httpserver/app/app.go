package main

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Создание конфигурации Redis
	redisConf := &redis.Options{
		Addr:     "localhost:8080",
		Password: "1234",
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true, // Установите false, если у вас есть подлинный сертификат
		},
	}

	// Создание клиента Redis
	redisClient := redis.NewClient(redisConf)

	// Проверка соединения
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to Redis")

	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)

	err = http.ListenAndServe(":3333", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
