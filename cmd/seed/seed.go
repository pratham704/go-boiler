package main

import (
	"github.com/pratham704/golang-ddd/internal/core/database/seed"
)

func main() {

	// Start the server
	if err := seed.NewSeed(); err != nil {
		// Handle error
		panic(err)
	}
}
