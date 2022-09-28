package hangmanweb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func Load(database string) (string, string, int, bool, string) {
	file, err := ioutil.ReadFile(database)
	var savefile SaveFormat
	err2 := json.Unmarshal(file, &savefile)
	if err != nil || err2 != nil {
		fmt.Println("No save file found.")
		os.Exit(0)
	}
	return savefile.Mot, savefile.MotVide, savefile.Attempt, savefile.Ascii, savefile.AsciiFile
}
