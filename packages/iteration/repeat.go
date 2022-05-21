package iteration

// Repeat takes a character as input and returns 
// a string of that character repeated 5 times
func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += character
	}
	return repeated
}