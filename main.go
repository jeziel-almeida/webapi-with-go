package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	fmt.Print("Enter the limit number: ")
	var limitNumber int
	fmt.Scan(&limitNumber)
	fmt.Println("The sorted number is: ", getRandomNumber(limitNumber))

}

func getRandomNumber(limitNumber int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(limitNumber)
}