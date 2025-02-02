package main

import (
	"fmt"
)

func Recipient(r string) string {
	return r
}

func main() {
	fmt.Printf("Hello %s!\n", Recipient("Paul"))
}