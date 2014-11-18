package main

import (
	"testing"
)

func TestChoosingBy(t *testing.T) {
	chooseFrom := []string {
		"hoge",
		"fuga",
		"piyo",
	}
	result := choosingBy(chooseFrom)
	if result != "hoge" && result != "fuga" && result != "piyo" {
		t.Errorf("got %s want \"hoge\" or \"fuga\" or \"piyo\"", result)
	}
}
