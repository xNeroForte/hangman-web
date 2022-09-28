package hangmanweb

import (
	"encoding/json"
	"log"
	"os"
)

type SaveFormat struct {
	Mot       string `json:"MotComplet"`
	MotVide   string `json:"MotVide"`
	Attempt   int    `json:"Attempt"`
	Ascii     bool   `json:"Ascii"`
	AsciiFile string `json:"AsciiFile"`
}

func SaveGame(word string, emptyword string, attempt int, ascii bool, asciiFile string) {
	savefile := SaveFormat{Mot: word, MotVide: emptyword, Attempt: attempt, Ascii: ascii, AsciiFile: asciiFile}
	sbytes, err := json.Marshal(savefile)
	err2 := os.WriteFile("save.txt", sbytes, 0666)
	if err != nil || err2 != nil {
		log.Fatal(err)
	}
}
