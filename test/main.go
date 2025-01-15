package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	id, _ := uuid.NewV7()
	fmt.Println("Generated UUIDv7:", id)
}
