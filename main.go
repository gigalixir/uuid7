package main

import (
	"fmt"

	"github.com/gofrs/uuid/v5"
)

func main() {
	uuid7, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	fmt.Println(uuid7.String())
}
