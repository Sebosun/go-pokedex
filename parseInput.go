package main

import "strings"

func parseInput(input string) []string {
	input = strings.TrimSpace(input)
	return strings.Split(input, " ")
}
