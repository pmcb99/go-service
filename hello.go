package main

import (
	"fmt"
)

func GetRecipient(r string) string {
	return r
}

func HelloRecipient(r string) string {
	if r != "" {
		return "Hello, " + r
	}
	return "Hello stranger..."
}

func main() {
	fmt.Println(HelloRecipient("Paul"))
}
