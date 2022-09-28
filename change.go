package hangmanweb

func Change(entry string, wordToComplete string, wordToFind string) string {
	result := ""
	lettre := rune(entry[0])
	for i := 0; i < len(wordToComplete); i++ {
		if rune(wordToFind[i]) == lettre {
			result += string(wordToFind[i])
			continue
		}
		result += string(wordToComplete[i])
	}
	return result
}
