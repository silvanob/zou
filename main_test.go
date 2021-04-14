package main

import (
	"fmt"
	"testing"
)

func TestSave(t *testing.T) {
	fmt.Println("5")
	if 5 == 6 {
		t.Errorf("hi")
	}
}
