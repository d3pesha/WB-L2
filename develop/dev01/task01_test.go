package dev01

import "testing"

func TestGetCurrentTime(t *testing.T) {
	_, err := getTimeNow()
	if err != nil {
		t.Fatalf("Test error: %v", err)
	}
}
