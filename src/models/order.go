package models

import (
	"encoding/json"
)

type Order struct {
	ID      int    `json:"id"`
	Symbol  string `json:"symbol"`
}
