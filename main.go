package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type InputBuffer struct {
	buffer       string
	bufferLength int32
	inputLength  int32
}

func main() {
	fmt.Printf("Hello")
}

func readInput(inputBuffer InputBuffer) {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
}

func createBuffer() InputBuffer {
	ib := InputBuffer{
	// buffer = fmt.Fscanf()
	}

	return ib
}
