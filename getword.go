package hangmanweb

import (
	"io/ioutil"
	"log"
)

func GetWord(dictionnary string, nb int) string {
	file, err := ioutil.ReadFile(dictionnary)
	if err != nil {
		log.Fatal(err)
	}
	ruru := RemoveAccent(file)
	counter := 1
	start := 0
	str := ""
	for i := 0; i < len(ruru); i++ {
		if ruru[i] == '\n' {
			counter++
		}
		if counter == nb {
			for j := i + start; j < len(ruru); j++ {
				if ruru[j] == '\n' {
					break
				}
				str += string(ruru[j])
				if j < len(ruru)-1 {
					if ruru[j+1] != '\n' {
						str += " "
					}
				}
			}
			break
		}
		start = 1
	}
	return str
}
