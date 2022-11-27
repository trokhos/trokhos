package main

import (
	"fmt"
	"log"

	"github.com/trokhos/trokhos/internal/server"
)

func main() {
	fmt.Println("Welcome to Trokhos... Wheels that can take you anywhere")
	log.Fatal(server.Run())
}
