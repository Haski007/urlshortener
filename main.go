package main

import (
	"math/rand"
	"fmt"
	"net/http"
	"time"

	"./config"
	"./database"
	"./handlers"
	"./logger"
)

func main() {
	initRand()

	database.InitDB()

	http.HandleFunc("/", logger.PreLogs(handlers.Main))

	fmt.Println("Listerning on port :8080")
	http.ListenAndServe(config.Port, nil)
}

// Inits seed for rand() func
func initRand() {
	rand.Seed(time.Now().UnixNano())
}