package src

// displayWord affiche le mot avec les lettres devinées.
func (g *Game) displayWord() string {
	display := ""
	for _, char := range g.currentWord {
		if contains(g.guessedLetters, string(char)) {
			display += string(char)
		} else {
			display += "_"
		}
		display += " "
	}
	return display
}

// contains vérifie si un élément existe dans un tableau.
func contains(arr []string, element string) bool {
	for _, e := range arr {
		if e == element {
			return true
		}
	}
	return false
}

// startGame initialise un nouveau jeu.
func (g *Game) startGame() {
	var wordlist []string
	g.guessedLetters = nil
	wordlist = ReadFile()
	g.currentWord = string(ChoseRandomWord(wordlist))
	g.attempts = 6
	g.hangmanStep = 0
	g.gameStarted = true
}

// restartGame réinitialise le jeu.
func (g *Game) restartGame() {
	g.gameStarted = false
}

// compareAndReveal compare une lettre et révèle si elle est dans le mot.
func (g *Game) compareAndReveal(letter string) {
	g.guessedLetters = append(g.guessedLetters, letter)
}
