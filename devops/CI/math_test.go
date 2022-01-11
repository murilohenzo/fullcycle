package main

import "testing"

func TestSum(t *testing.T)  {
	sum := Sum(15, 15)
	expected := 30

	if sum != expected {
		t.Errorf("Invalid sum result: Result %d. Expected %d", sum, expected)
	}

}