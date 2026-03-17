package main

import (
	"aurora-trading-platform/src/controllers"
	"aurora-trading-platform/src/models"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Aurora Trading Platform started")
	http.HandleFunc("/", controllers.IndexHandler)
	http.HandleFunc("/orders", controllers.OrdersHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
