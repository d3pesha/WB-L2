package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCD(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	initialDir, _ := os.Getwd()

	cd([]string{"cd", tmpDir})
	currentDir, _ := os.Getwd()
	if currentDir != tmpDir {
		t.Errorf("Expected current directory to be %s, got %s", tmpDir, currentDir)
	}

	cd([]string{"cd", initialDir})
	currentDir, _ = os.Getwd()
	if currentDir != initialDir {
		t.Errorf("Expected current directory to be %s, got %s", initialDir, currentDir)
	}
}
