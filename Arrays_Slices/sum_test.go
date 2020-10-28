package main

import "testing"

// Arrays have a fixed capacity which you define when you declare the variable. We can initialize an array in two ways:
// numbers := [5]int{1, 2, 3, 4, 5}
// numbers := [...]int{1, 2, 3, 4, 5}

func TestSum(t *testing.T) {

	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
