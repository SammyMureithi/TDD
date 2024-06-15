package iteration

func Repeat(character string, times int) string {
	var resultString string
	for i := 0; i < times; i++ {
		resultString += character
	}
	return resultString
}