package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAnagram(t *testing.T) {
	tests := []struct {
		input    []string
		expected map[string][]string
	}{
		{
			input:    []string{"Кот", "Ток", "Дом", "Мод", "Мот", "Том"},
			expected: map[string][]string{"дом": {"дом", "мод"}, "кот": {"кот", "ток"}, "мот": {"мот", "том"}},
		},
		{
			input:    []string{"Апельсин", "Лимон", "Спаниель", "Мелисса", "Лиса"},
			expected: map[string][]string{"апельсин": {"апельсин", "спаниель"}},
		},
	}
	for _, test := range tests {
		result := Anagram(test.input)
		require.Equal(t, test.expected, result)
	}
}
