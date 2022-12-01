package main

import (
	api "example/gopractice/api"
	"example/gopractice/client"
	"fmt"
)

func main() {
	fmt.Println("Starting server...")
	go api.RunServer()
	fmt.Println("Server started.  Starting Client...")
	client.RunClient()
}
