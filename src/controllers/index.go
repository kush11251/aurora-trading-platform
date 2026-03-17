package controllers

import (
	"aurora-trading-platform/src/models"
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Aurora Trading Platform")
}
