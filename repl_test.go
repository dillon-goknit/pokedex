package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q): got %d words, want %d", c.input, len(actual), len(c.expected))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%q): word %d, is %q, want %q", c.input, i, word, expectedWord)
			}
		}
	}
}

func TestCommandResponse(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    "well hello there",
			expected: "Your command was: well",
		},
		{
			input:    "POKEMON was underrated",
			expected: "Your command was: pokemon",
		},
		{
			input:    "    charmander   ",
			expected: "Your command was: charmander",
		},
		{
			input:    "",
			expected: "",
		},
		{
			input:    " ",
			expected: "",
		},
	}
	for _, c := range cases {
		actual := commandResponse(c.input)
		if actual != c.expected {
			t.Errorf("commandResponse(%q): got %q, want %q", c.input, actual, c.expected)
		}
	}
}
