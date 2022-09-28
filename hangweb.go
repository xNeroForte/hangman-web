package hangmanweb

import (
	"fmt"
)

// INITAILISATION DES VARIABLES
var (
	file           string
	LetterState    string
	TempFile       string
	WordToFind     string = ""
	WordInComplete string = ""
	AllUsedLetters []rune
	Attempts       int     = 10
	Complete       bool    = false
	ptrUsed        *[]rune = &AllUsedLetters
)

func HangmanFileInit(nb int) {
	if nb == 1 {
		file = "words.txt"
		TempFile = file
		return
	}

	file = "words" + NbToString(nb) + ".txt"
	TempFile = file
}

func HangmanWebInit() {
	// CHARGEMENT DES VARIABLE
	// prendre le WordToFind
	Complete = false
	var temp []rune
	AllUsedLetters = temp
	nb := RandomWordNumber(file)
	WordToFind = GetWord(file, nb)
	WordToFind = ToUpper(WordToFind)
	// detecter les input et les valider
	WordInComplete = RevealLetters(WordToFind)
	fmt.Println("Good Luck, you have " + NbToString(Attempts) + " Attempts.")
	println(WordInComplete)
	fmt.Println()
}

func HangmanWebPlay(entry string) *string {

	// CE QUI SE PASSE IN GAME
	var HavetoReturn bool
	var RetrunState string
	entry, HavetoReturn, RetrunState = TreatInput(entry, WordToFind, ptrUsed)
	if HavetoReturn {
		return &RetrunState
	}
	entry = ToUpper(entry)
	var letter rune
	for _, v := range entry {
		letter = v
		break
	}

	if len(entry) == 1 {
		if IsRight(WordToFind, letter) {
			oldstr := WordInComplete
			WordInComplete = Change(entry, WordInComplete, WordToFind)
			if oldstr != WordInComplete {
				Picture("thumb.txt", 1, 15)
				UsedLetters(entry, ptrUsed)
				println(WordInComplete)
				fmt.Println()
				Complete = IsComplete(WordInComplete, WordToFind)

			} else {
				Attempts--
				if Attempts > 0 {
					fmt.Println("No more of this letter in the word, " + NbToString(Attempts) + " attempt(s) remaining\n")
					RetrunState = "BAD CHOISE"
					Picture("hangman.txt", 10-Attempts, 8)
					UsedLetters(entry, ptrUsed)
					println(WordInComplete)
					fmt.Println()
				}
			}

		} else {
			Attempts--
			if Attempts > 0 {
				fmt.Println("Not present in the word, " + NbToString(Attempts) + " attempt(s) remaining\n")
				RetrunState = "BAD CHOISE"
				Picture("hangman.txt", 10-Attempts, 8)
				UsedLetters(entry, ptrUsed)
				println(WordInComplete)
				fmt.Println()
			}
		}
	} else {
		if SameWord(entry, WordToFind) {
			WordInComplete = entry
			Picture("thumb.txt", 1, 15)
			println(RemoveSpace(WordToFind))
			fmt.Println()
			Complete = true

		} else {
			Attempts -= 2
			if Attempts > 0 {
				fmt.Println("Incorrect word, " + NbToString(Attempts) + " attempt(s) remaining\n")
				RetrunState = "BAD CHOISE"
				Picture("hangman.txt", 10-Attempts, 8)
				UsedLetters(entry, ptrUsed)
				println(WordInComplete)
			}
		}
	}

	if Complete {
		fmt.Println("Congrats !")
	} else if Attempts <= 0 {
		Picture("hangman.txt", 10, 8)
		fmt.Println("You failed to save José. He died.")
		fmt.Println("The word was : ")
		println(RemoveSpace(WordToFind))
		fmt.Println()

	}

	return &RetrunState
}

/*func OldHangmanWebPlay() {
	file := os.Args[1]
	asciiMode := false
	var asciiFile string
	var tempAsciiMode bool
	var tempAsciiFile string
	if len(os.Args) < 2 || len(os.Args) > 6 {
		fmt.Println("Error : invalid number of arguments")
		os.Exit(0)
	}
	asciiPos := 0
	for i := 0; i < len(os.Args); i++ {
		if os.Args[i] == "--letterFile" {
			asciiPos = i
		}
	}
	if asciiPos != 0 {
		if len(os.Args) > asciiPos && os.Args[asciiPos] == "--letterFile" {
			if len(os.Args) < asciiPos+2 {
				fmt.Println("No ascii template found.")
				os.Exit(0)
			} else {
				asciiFile = os.Args[asciiPos+1]
				tempAsciiFile = os.Args[asciiPos+1]
				asciiMode = true
				tempAsciiMode = true
			}
		}
	}
	autosave := false
	WordToFind := ""
	WordInComplete := ""
	var AllUsedLetters []rune
	ptrUsed := &AllUsedLetters
	Attempts := 10
	Complete := false
	loadargnb := -1
	if HaveToLoad() {
		loadargnb = LoadArgs()
		if loadargnb >= len(os.Args) {
			fmt.Println("No save file found.")
			os.Exit(0)
		}
		WordToFind, WordInComplete, Attempts, asciiMode, asciiFile = Load(os.Args[loadargnb])
		if tempAsciiMode {
			asciiMode = tempAsciiMode
			asciiFile = tempAsciiFile
		}
		if Attempts > 0 {
			autosave = true
			println("Welcome Back, you have " + NbToString(Attempts) + " Attempts remaining.")
			if asciiMode {
				AsciiArt(RemoveSpace(WordInComplete), asciiFile)
			} else {
				println(WordInComplete + "\n")
			}
		} else {
			fmt.Println("No save file found.")
			os.Exit(0)
		}
	} else {
		// prendre le WordToFind
		nb := RandomWordNumber(file)
		WordToFind = GetWord(file, nb)
		WordToFind = ToUpper(WordToFind)
		// detecter les input et les valider
		WordInComplete = RevealLetters(WordToFind)
		fmt.Println("Good Luck, you have " + NbToString(Attempts) + " Attempts.")
		if asciiMode {
			AsciiArt(RemoveSpace(WordInComplete), asciiFile)
		} else {
			println(WordInComplete)
		}
		fmt.Println()
	}
	savemode := false
	for !Complete && Attempts > 0 {
		fmt.Printf("Choose: ")
		var entry string
		fmt.Scanln(&entry)
		entry, savemode = TreatInput(entry, WordToFind, ptrUsed)
		if savemode {
			SaveGame(WordToFind, WordInComplete, Attempts, asciiMode, asciiFile)
			os.Exit(0)
		}
		entry = ToUpper(entry)
		var letter rune
		for _, v := range entry {
			letter = v
			break
		}
		if len(entry) == 1 {
			if IsRight(WordToFind, letter) {
				oldstr := WordInComplete
				WordInComplete = Change(entry, WordInComplete, WordToFind)
				if oldstr != WordInComplete {
					Picture("thumb.txt", 1, 15)
					AllUsedLetters(entry, ptrUsed)
					if asciiMode {
						AsciiArt(RemoveSpace(WordInComplete), asciiFile)
					} else {
						println(WordInComplete)
					}
					fmt.Println()
					Complete = IsComplete(WordInComplete, WordToFind)

				} else {
					Attempts--
					if Attempts > 0 {
						fmt.Println("No more of this letter in the word, " + NbToString(Attempts) + " attempt(s) remaining\n")
						Picture("hangman.txt", 10-Attempts, 8)
						AllUsedLetters(entry, ptrUsed)
						if asciiMode {
							AsciiArt(RemoveSpace(WordInComplete), asciiFile)
						} else {
							println(WordInComplete)
						}
						fmt.Println()
					}
				}
			} else {
				Attempts--
				if Attempts > 0 {
					fmt.Println("Not present in the word, " + NbToString(Attempts) + " attempt(s) remaining\n")
					Picture("hangman.txt", 10-Attempts, 8)
					AllUsedLetters(entry, ptrUsed)
					if asciiMode {
						AsciiArt(RemoveSpace(WordInComplete), asciiFile)
					} else {
						println(WordInComplete)
					}
					fmt.Println()
				}
			}

		} else {
			if SameWord(entry, WordToFind) {
				WordInComplete = entry
				Picture("thumb.txt", 1, 15)
				if asciiMode {
					AsciiArt(RemoveSpace(WordToFind), asciiFile)
				} else {
					println(RemoveSpace(WordToFind))
				}
				fmt.Println()
				Complete = true


			} else {
				Attempts -= 2
				if Attempts > 0 {
					fmt.Println("Incorrect word, " + NbToString(Attempts) + " attempt(s) remaining\n")
					Picture("hangman.txt", 10-Attempts, 8)
					AllUsedLetters(entry, ptrUsed)
					if asciiMode {
						AsciiArt(RemoveSpace(WordInComplete), asciiFile)
					} else {
						println(WordInComplete)
					}
				}
			}
		}
		if autosave {
			SaveGame(WordToFind, WordInComplete, Attempts, asciiMode, asciiFile)
		}
	}
	if Complete {
		fmt.Println("Congrats !")
		if autosave {
			SaveGame("", "", 0, false, "")
		}
	} else if Attempts <= 0 {
		Picture("hangman.txt", 10, 8)
		fmt.Println("You failed to save José. He died.")
		fmt.Println("The word was : ")
		if asciiMode {
			AsciiArt(RemoveSpace(WordToFind), asciiFile)
		} else {
			println(RemoveSpace(WordToFind))
		}
		fmt.Println()
		if autosave {
			SaveGame("", "", 0, false, "")
		}
	}
}*/
