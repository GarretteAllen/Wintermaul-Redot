package utils

import (
	"log"

	"github.com/google/uuid"
)

func GenerateUniqueID() string {
	id, err := uuid.NewUUID()
	if err != nil {
		log.Println("Error generating UUID:", err)
		return ""
	}
	return id.String()
}
