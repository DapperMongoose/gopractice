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

func incrementCounter() error {
	currentCount, err := api.ReadDB()
	if err != nil {
		return err
	}
	currentCount++
	err = api.WriteDB(currentCount)
	return err

}

func decrementCounter() error {
	currentCount, err := api.ReadDB()
	if err != nil {
		return err
	}
	currentCount--
	err = api.WriteDB(currentCount)
	return err
}
