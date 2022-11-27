package main

import (
	db "example/gopractice/api"
	"fmt"
)

func main() {
	// messy placeholder code to validate the back end package during development, please ignore
	currentCount, err := db.ReadDB()
	if err != nil {
		fmt.Print(err)
	}
	// print current count
	fmt.Printf("Current count: %d\n", currentCount)

	//increment by one
	err = incrementCounter()
	if err != nil {
		fmt.Println(err)
	}
	currentCount, err = db.ReadDB()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("Current count: %d\n", currentCount)

	err = incrementCounter()
	if err != nil {
		fmt.Println(err)
	}
	currentCount, err = db.ReadDB()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("Current count: %d\n", currentCount)

	err = decrementCounter()
	if err != nil {
		fmt.Println(err)
	}
	currentCount, err = db.ReadDB()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("Current count: %d\n", currentCount)

	// reset to 0
	err = db.ResetDB()
	if err != nil {
		fmt.Println(err)
	}
	currentCount, err = db.ReadDB()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("Current count: %d\n", currentCount)
}

func incrementCounter() error {
	currentCount, err := db.ReadDB()
	if err != nil {
		return err
	}
	currentCount++
	err = db.WriteDB(currentCount)
	return err

}

func decrementCounter() error {
	currentCount, err := db.ReadDB()
	if err != nil {
		return err
	}
	currentCount--
	err = db.WriteDB(currentCount)
	return err
}
