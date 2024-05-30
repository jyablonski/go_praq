package iteration

func Repeat(character string, repeats int) string {
	var repeated string
	for i := 0; i < repeats; i++ {
		repeated = repeated + character
	}
	return repeated
}
