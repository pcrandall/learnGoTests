package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat(5, "a")
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// Benchmarking code has to go in the <filename>_test.go file
// TO RUN BENCHMARK "GO TEST -BENCH=."
// by default benchmarks are ran sequentually
// Form  func BenchmarkXxx(*testing.B)

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(5, "a")
	}
}

// EXAMPLES
// functions that reside in the <filename>_test.go file
// take no arguements(niladic) and begind with the word Example instead of Test
// niladic means a function with no arguements
// when running test, compares output to the output comment in the Example function

func ExampleRepeat() {
	fmt.Println(Repeat(5, "a"))
	// Output: aaaaa
}
