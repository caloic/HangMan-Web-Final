package src

type Game struct {
	currentWord    string
	guessedLetters []string
	attempts       int
	gameStarted    bool
	hangmanStep    int
	guessletters   string
}

var t templateData

type templateData struct {
	Theword        string
	WordDisplay    string
	GuessedLetters string
	Attempts       int
	HangmanImage   string
}
