package main

import (
	"os"
	"testing"
)

func TestSaveToAscii(t *testing.T) {
	if saveToAscii("./dove.png") == "failed" {
		t.Errorf("Failed Conversion")
	}
}

func TestAsciiDir(t *testing.T) {
	_, err := os.Stat("./ascii")
	if err != nil {
		t.Errorf("Directory 'ascii' DNE")
	}
}

func TestUploadsDir(t *testing.T) {
	_, err := os.Stat("./uploads")
	if err != nil {
		t.Errorf("Directory 'uploads' DNE")
	}
}
