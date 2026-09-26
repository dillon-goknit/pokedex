package main

import (
	"strings"
)

func cleanInput(text string) []string {
	lowercase := strings.ToLower(text)
	words := strings.Fields(lowercase)
	return words
}

func commandResponse(command string) string {
	words := cleanInput(command)
	if len(words) == 0 {
		return ""
	}
	first := words[0]
	return "Your command was: " + first
}
