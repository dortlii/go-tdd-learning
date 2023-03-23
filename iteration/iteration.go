package iteration

// Repeat takes a character as string and a count as int and returns the character repeated.
func Repeat(character string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += character
	}

	return repeated
}
