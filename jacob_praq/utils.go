package main

import (
	"github.com/google/uuid"
)

func generate_uuid() string {
	uuid_val := uuid.New().String()

	return uuid_val
}
