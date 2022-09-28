package hangmanweb

import (
	"os"
)

func LoadArgs() int {
	load := []rune("--startWith")
	equal := true
	var ruru []rune
	for i := 0; i < len(os.Args); i++ {
		equal = true
		ruru = []rune(os.Args[i])
		for j := 0; j < len(load) && j < len(ruru); j++ {
			if len(ruru) != len(load) {
				equal = false
			}
			if ruru[j] != load[j] {
				equal = false
			}
		}
		if equal {
			return i + 1
		}

	}

	return -1
}
