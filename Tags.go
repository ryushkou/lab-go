package main

import (
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	tagsWriter := bufio.NewWriter(os.Stderr)
	textWriter := bufio.NewWriter(os.Stdout)
	writer := &textWriter
	for _, c := range input {
		if c == '<' {
			writer = &tagsWriter
		}
		(*writer).WriteString(string(c))
		if c == '>' {
			writer = &textWriter
		}
	}
	tagsWriter.Flush()
	textWriter.Flush()
}
