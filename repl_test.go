package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " Hello  wOrLd ",
			expected: []string{"hello", "world"},
		},
	}
	for _, cs := range cases {
		assert.Equal(t, cleanInput(cs.input), cs.expected)
	}
}
