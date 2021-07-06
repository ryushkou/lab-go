package main

import (
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	text, tags := split(input)

	textWriter := bufio.NewWriter(os.Stdout)
	textWriter.Write([]byte(text))
	textWriter.Flush()

	tagsWriter := bufio.NewWriter(os.Stderr)
	tagsWriter.Write([]byte(tags))
	tagsWriter.Flush()
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
