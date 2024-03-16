package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	rc := triangle{
		height: 5.6,
		side:   6.5,
	}

	sq := square{
		sideLength: 6,
	}

	printArea(rc)
	printArea(sq)

	fileName := os.Args[1]
	f, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
