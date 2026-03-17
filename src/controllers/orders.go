package controllers

import (
	"aurora-trading-platform/src/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func OrdersHandler(w http.ResponseWriter, r *http.Request) {
	orders := []models.Order{{ID: 1, Symbol: "AAPL"}, {ID: 2, Symbol: "GOOG"}}
	json.NewEncoder(w).Encode(orders)
}
