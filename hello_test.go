package hello

import "testing"

func TestSay(t *testing.T) {
	input := "Pankaj"
	expected :="Hello, Pankaj"
	output := Say("Pankaj")

	if output != expected {
		t.Errorf("wanted %s, got %s", expected, input)
	}


}