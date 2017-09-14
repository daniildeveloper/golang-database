package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// InputBuffer description
type InputBuffer struct {
	buffer       string
	bufferLength int
	inputLength  int
}

func main() {
	input := createBuffer()

	for true {
		printPromt()

		readInput(&input)

		if strings.Compare(input.buffer, ".exit") == 0 {
			os.Exit(1)
		}
	}
}

func readInput(inputBuffer *InputBuffer) {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	// text = strings.TrimSpace()
	text = strings.TrimRight(text, "\r\n")
	inputBuffer.buffer = text

	if len(text) <= 0 {
		panic("Error reading input")
	}

	inputBuffer.bufferLength = len(text)
	inputBuffer.buffer = text
}

// create new buffer
func createBuffer() InputBuffer {
	ib := InputBuffer{
		buffer:       "",
		bufferLength: 0,
		inputLength:  0}

	return ib
}

func printPromt() {
	fmt.Println("db > ")
}
func printStringInfo(str string) {
	fmt.Printf("String: %s\n", str)
	fmt.Printf("Length: %d\n", len(str))
}
