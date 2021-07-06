package main

import (
	"testing"
)

func TestSplit(t *testing.T) {
	test_inputs := []string{
		"aaa",
		"<tag>bla bla bla</tag>",
		"a<t></t>b",
		"<c><b>",
		"<c<b>",
		"<>",
		"  <c>  <b>  "}
	expected_texts := []string{
		"aaa",
		"bla bla bla",
		"ab",
		"",
		"",
		"",
		"      "}
	expected_tags := []string{
		"",
		"<tag></tag>",
		"<t></t>",
		"<c><b>",
		"<c<b>",
		"<>",
		"<c><b>"}

	for i, input := range test_inputs {
		text, tags := split(input)
		if expected_texts[i] != text {
			t.Errorf("[Test %d] is failed for TEXT (expected=%s, actual=%s)", i, expected_texts[i], text)
		}
		if expected_tags[i] != tags {
			t.Errorf("[Test %d] is failed for TAGS (expected=%s, actual=%s)", i, expected_tags[i], tags)
		}
	}
}
