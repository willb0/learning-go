package iteration

func Repeat(character string, len int) string {
	var repeated string
	for i := 0; i < len; i++ {
		repeated += character
	}
	return repeated
}
