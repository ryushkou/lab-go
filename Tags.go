package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	input := read(os.Stdin)
	text, tags := split(input)
	write(os.Stdout, text)
	write(os.Stderr, tags)
}

func split(input string) (text, tags string) {
	var out = &text
	for _, c := range input {
		if c == '<' {
			out = &tags
		}
		*out += string(c)
		if c == '>' {
			out = &text
		}
	}
	return text, tags
}

func read(reader io.Reader) string {
	bReader := bufio.NewReader(os.Stdin)
	input, _ := bReader.ReadString('\n')
	return input
}

func write(writer io.Writer, output string) {
	bWriter := bufio.NewWriter(writer)
	bWriter.Write([]byte(output))
	bWriter.Flush()
}
