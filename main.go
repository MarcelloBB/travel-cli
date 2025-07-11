package main

import (
	"fmt"
	"travel-cli/cmd"
	"travel-cli/db"
)

func main() {
	_, err := db.ConnectDB()
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
	}
	cmd.Execute()
}
