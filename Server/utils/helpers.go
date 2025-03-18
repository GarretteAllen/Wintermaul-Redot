package utils

import (
	"encoding/json"
	"log"
)

func ParseJSON(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		log.Println("JSON parsing error:", err)
	}
	return err
}
