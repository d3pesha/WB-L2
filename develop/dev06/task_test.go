package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPrintString(t *testing.T) {
	f := flags{}
	tests := []struct {
		name           string
		input          string
		expectedOutput string
		flags          flags
	}{
		{
			name:           "Print selected field",
			input:          "apple orange banana",
			expectedOutput: "orange\n",
			flags: flags{
				field:     2,
				delimiter: " ",
				separated: false,
			},
		},
		{
			name:           "Print separated field",
			input:          "apple orange banana",
			expectedOutput: "orange",
			flags: flags{
				field:     2,
				delimiter: " ",
				separated: true,
			},
		},
		{
			name:           "Print non-existent field",
			input:          "apple orange banana",
			expectedOutput: "\n",
			flags: flags{
				field:     5,
				delimiter: " ",
				separated: false,
			},
		},
		{
			name:           "Print entire string when separated is true",
			input:          "apple orange banana",
			expectedOutput: "apple orange banana",
			flags: flags{
				field:     2,
				delimiter: " ",
				separated: true,
			},
		},
	}

	for _, test := range tests {
		inp := f.printString(test.input)
		require.Equal(t, test.expectedOutput, inp)
	}
}
