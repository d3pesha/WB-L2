package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUniqueString(t *testing.T) {
	test := []string{"12345", "12345", "abc", "abc"}
	result := UniqueString(test)
	require.Equal(t, result, []string{"12345", "abc"})
}

func TestSortByColumn(t *testing.T) {
	test := []string{"apple orange banana", "kiwi grape", "cherry pear", "melon lemon"}
	result := SortByColumn(test, 1)
	require.Equal(t, result, []string{"kiwi grape", "melon lemon", "apple orange banana", "cherry pear"})
}

func TestSortNum(t *testing.T) {
	test := []string{"8 2 6 1 21 200"}
	result := SortNum(test)
	require.Equal(t, result, []string{"1 2 6 8 21 200"})
}

func TestReversSort(t *testing.T) {
	test := []string{"12345", "123", "12345", "abc", "abc"}
	result := ReversSort(test)
	require.Equal(t, result, []string{"abc", "abc", "12345", "12345", "123"})
}
