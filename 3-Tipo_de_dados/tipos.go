package main

import (
	"errors"
	"fmt"
)

func main() {
	var err error = errors.New("texto do erro")

	fmt.Println(err)
}
