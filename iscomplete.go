package hangmanweb

func IsComplete(wordToFind string, str string) bool {
	ruruToFind := []rune(wordToFind)
	ruru := []rune(str)
	for i := 0; i < len(str); i++ {
		if ruruToFind[i] != ruru[i] {
			return false
		}
	}
	return true
}
