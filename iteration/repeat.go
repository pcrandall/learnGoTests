package main

const repeatCount = 5

func Repeat(reps int, character string) string {
	var repeated string
	var repitions int

	if reps > 0 {
		repitions = reps
	} else {
		repitions = repeatCount
	}

	for i := 0; i < repitions; i++ {
		repeated += character
	}
	return repeated
}
