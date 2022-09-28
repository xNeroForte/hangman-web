package hangmanweb

import (
	"fmt"
	"io/ioutil"
	"os"
)

func AsciiArt(word string, filename string) {
	var tabChar [][]string
	var initRow []string
	tabChar = append(tabChar, initRow)
	tabChar[0] = append(tabChar[0], "")
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("No ascii template found.")
		os.Exit(0)
	}
	index := 0
	row := 0
	for _, b := range file {
		if b != 10 && b != 13 {
			tabChar[index][row] += string(b)
		}
		if b == 10 {
			tabChar[index] = append(tabChar[index], "")
			row++
			if row > 8 {
				var char []string
				tabChar = append(tabChar, char)
				row = 0
				index++
				tabChar[index] = append(tabChar[index], "")
			}
		}
	}
	for i := 0; i <= row; i++ {
		for _, v := range word {
			if v == '_' {
				fmt.Print(tabChar[63][i])
			} else if v == '-' {
				fmt.Print(tabChar[13][i])
			}
			for r := 65; r <= 90; r++ {
				if v == rune(r) {
					fmt.Print(tabChar[r-32][i])
				}
			}
		}
		fmt.Print("\n")
	}
}
