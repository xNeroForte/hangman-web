package main

import (
	"fmt"
	"hangmanweb"
	"net/http"
	"text/template"
)

type Page struct {
	Title          *string
	Sub            string
	LettreStateWeb *string
	MyScore        string
}

var (
	Score         int  = 0
	InGame        bool = false
	WinScreen     bool = false
	WelcomeScreen bool = true
	TowerOfGodMod bool = false
	tmpl          *template.Template
	Username      string
)

func main() {
	tmpl, _ = template.ParseGlob("./Web/templates/*.html")
	// Set routing rules
	http.HandleFunc("/", FunctionManager)

	fileserver := http.FileServer(http.Dir("./Web/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	//Use the default DefaultServeMux.
	http.ListenAndServe("localhost:8080", nil)

}

func FunctionManager(w http.ResponseWriter, r *http.Request) {
	if WelcomeScreen {
		fmt.Println("GO TO WELCOME")
		Welcome(w, r)
	}
	if !InGame && !WelcomeScreen && Username != "" {
		Home(w, r)
	}
	if InGame && !WinScreen {
		GameMode(w, r)
	}
	if WinScreen {
		EndMode(w, r)
	}
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	UsernameRaw := r.Form["Username"]
	if len(UsernameRaw) > 0 {
		Username = UsernameRaw[0]
		fmt.Println("User is " + Username)
		WelcomeScreen = false
	}

	/*AllMyMot := r.Form["PLAY"]
	if len(AllMyMot) > 0 {
		MyButton := AllMyMot[0]
		fmt.Println(MyButton)

		if MyButton == "PLAY" {

		}

	}*/

	if WelcomeScreen {
		tmpl.ExecuteTemplate(w, "welcome", tmpl)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {

	MainVarReset()
	r.ParseForm()
	AllMyMot := r.Form["subject"]
	//var MyMot string
	if len(AllMyMot) > 0 {
		MyButton := AllMyMot[0]
		fmt.Println(MyButton)

		if MyButton == "EASY" {
			hangmanweb.HangmanFileInit(1)
		} else if MyButton == "MEDIUM" {
			hangmanweb.HangmanFileInit(2)
		} else if MyButton == "HARD" {
			hangmanweb.HangmanFileInit(3)
		} else if MyButton == "TOWEROFGOD" {
			hangmanweb.HangmanFileInit(4)
			TowerOfGodMod = true
		}
		fmt.Println(hangmanweb.TempFile)
		hangmanweb.HangmanWebInit()
		InGame = true
	}
	if !InGame {
		fmt.Println("INGAME")
		tmpl.ExecuteTemplate(w, "home", "")
	}

}

func GameMode(w http.ResponseWriter, r *http.Request) {

	/*data := Page{
		Title: "Wesh Wesh",
		Sub:   "Canne à peche",
	}
	tmpl.ExecuteTemplate(w, "index", data)*/

	MyGame := Page{
		Title:          &hangmanweb.WordInComplete,
		Sub:            " ",
		LettreStateWeb: &hangmanweb.LetterState,
		MyScore:        Username + "'s Score is " + hangmanweb.NbToString(Score),
	}
	r.ParseForm()
	gobackbutton := r.Form["GOBACK"]
	if len(gobackbutton) > 0 {
		GoBack := gobackbutton[0]
		if GoBack == "GOBACK" {
			fmt.Println("GOBACK CLICKED")
			InGame = false
			Home(w, r)
		}
	}

	AllMyMot := r.Form["MyLetter"]
	if len(AllMyMot) > 0 {
		MyMot := AllMyMot[0]
		if len(MyMot) > 0 {
			MyGame.LettreStateWeb = hangmanweb.HangmanWebPlay(MyMot)
		}
		fmt.Println(hangmanweb.LetterState)
	}

	//WinString := "YOU WIN!"
	if (hangmanweb.Complete && !TowerOfGodMod) || hangmanweb.Attempts <= 0 {
		println("TOWER COMPLETED")
		WinScreen = true
		//MyGame.Title = &WinString
	}

	if TowerOfGodMod && hangmanweb.Complete && hangmanweb.Attempts > 0 {
		println("NEXT LEVEL")
		Score += (hangmanweb.Attempts * 2)
		if hangmanweb.Attempts < 15 {
			hangmanweb.Attempts++
		}

		hangmanweb.HangmanWebInit()
	}
	if TowerOfGodMod {
		MyGame.MyScore = Username + "'s Score is " + hangmanweb.NbToString(Score)
	} else {
		MyGame.MyScore = "Score is avaible only in Tower Of God"
	}
	MyGame.Sub = "T'as encore " + hangmanweb.MyAttemptToString(&hangmanweb.Attempts) + " vies"
	if InGame && !WinScreen {
		tmpl.ExecuteTemplate(w, "index", MyGame)
		//tmpl.ExecuteTemplate(w, "lettreused", "")
	}
	//http.ServeFile(w, r, "./Web/templates/index.html")
	fmt.Println("LE MOT EST " + hangmanweb.WordToFind)
}

func EndMode(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	AllMyMot := r.Form["PLAYAGAIN"]

	if len(AllMyMot) > 0 {
		MyButton := AllMyMot[0]
		if MyButton == "PLAYAGAIN" {
			fmt.Println("PlayAgain Clicked")
			MainVarReset()
		} else {
			fmt.Println("BACK TO HOOOOME")
			fmt.Println("BACKTOHOME Clicked")
			WelcomeScreen = true
			MainVarReset()
		}
	}
	WinOrLose := Page{
		Title:          &hangmanweb.WordToFind,
		Sub:            "",
		LettreStateWeb: &hangmanweb.LetterState,
	}
	if hangmanweb.Complete {
		WinOrLose.Sub = " You WIN! The Word was " + hangmanweb.WordToFind
	} else {
		WinOrLose.Sub = " You LOSE! The Word was " + hangmanweb.WordToFind
	}
	if WinScreen {
		tmpl.ExecuteTemplate(w, "endscreen", WinOrLose)
		fmt.Println("Winscreen")
	} else if !WelcomeScreen {
		fmt.Println("PLAY AGAIN")
		Home(w, r)
	} else {
		fmt.Println("Back")
		Welcome(w, r)
	}

}

func MainVarReset() {
	InGame = false
	Score = 0
	hangmanweb.Attempts = 10
	WinScreen = false
	TowerOfGodMod = false
	fmt.Println("RESET VARIABLE")
}
