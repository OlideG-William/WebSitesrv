package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name     string
	Standard int `json:"Standard"`
}

func main() {
	timeTrack := time.Now()

	fmt.Println("Hello me this test file")

	fmt.Println("Compile program: ", time.Since(timeTrack))
}
