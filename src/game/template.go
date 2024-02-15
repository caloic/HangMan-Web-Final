package src

import (
	"html/template"
	"net/http"
	"strings"
)

// hangmanHandler gère les requêtes HTTP pour le jeu du pendu.
func (g *Game) HangmanHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method == http.MethodGet {
		// Gestion des requêtes GET
		if r.URL.Path == "/" && !g.gameStarted {
			renderTemplate(w, "view/html/menu.html", nil)
			return
		}

		// Ajout de la gestion de la première porte
		if r.URL.Path == "/1_path_to_hangman" && !g.gameStarted {
			renderTemplate(w, "view/html/Path_to_hangman/Path/1_path_to_hangman.html", nil)
			return
		}

		if r.URL.Path == "/2_lunch_path_to_hangman" && !g.gameStarted {
			renderTemplate(w, "view/html/Path_to_hangman/Path/2_lunch_path_to_hangman.html", nil)
			return
		}

		if r.URL.Path == "/2_bar_path_to_hangman" && !g.gameStarted {
			renderTemplate(w, "view/html/Path_to_hangman/Path/2_bar_path_to_hangman.html", nil)
			return
		}

		if r.URL.Path == "/3_solodrink_path_to_hangman" && !g.gameStarted {
			renderTemplate(w, "view/html/Path_to_hangman/Path/3_solodrink_path_to_hangman.html", nil)
			return
		}

		if r.URL.Path == "/3_drink_path_to_hangman" && !g.gameStarted {
			renderTemplate(w, "view/html/Path_to_hangman/Path/3_drink_path_to_hangman.html", nil)
			return
		}

		if r.URL.Path == "/4_robbery_path_to_hangman" && !g.gameStarted {
			renderTemplate(w, "view/html/Path_to_hangman/Path/4_robbery_path_to_hangman.html", nil)
			return
		}

		if r.URL.Path == "/door_lost_path_to_hangman" && !g.gameStarted {
			renderTemplate(w, "view/html/Path_to_hangman/Lost_back/door_lost_path_to_hangman.html", nil)
			return
		}

		if r.URL.Path == "/5_backdoor_path_to_hangman" && !g.gameStarted {
			renderTemplate(w, "view/html/Path_to_hangman/Path/5_backdoor_path_to_hangman.html", nil)
			return
		}

		if r.URL.Path == "/simpleDoorLost_blocked" && !g.gameStarted {
			renderTemplate(w, "view/html/Path_to_hangman/Lost_back/simpleDoorLost_blocked.html", nil)
			return
		}

		if r.URL.Path == "/startGame_last_path" && !g.gameStarted {
			renderTemplate(w, "view/html/Path_to_hangman/startGame_last_path.html", nil)
			return
		}

		if r.URL.Path == "/3_break_path_to_hangman" && !g.gameStarted {
			renderTemplate(w, "view/html/Path_to_hangman/Path/3_break_path_to_hangman.html", nil)
			return
		}

		if r.URL.Path == "/mustang_path" {
			renderTemplate(w, "view/html/Path_after_hangman/mustang_path.html", nil)
			return
		}

		if r.URL.Path == "/crashedmustang" {
			renderTemplate(w, "view/html/Path_after_hangman/Game_lost/crashedmustang.html", nil)
			return
		}

		if r.URL.Path == "/realwinmustang" {
			renderTemplate(w, "view/html/Path_after_hangman/Game_win/realwinmustang.html", nil)
			return
		}

		if r.URL.Path == "/motorbike_path" {
			renderTemplate(w, "view/html/Path_after_hangman/motorbike_path.html", nil)
			return
		}

		if r.URL.Path == "/crashedMotorbike" {
			renderTemplate(w, "view/html/Path_after_hangman/Game_lost/crashedMotorbike.html", nil)
			return
		}

		if r.URL.Path == "/2realwinmotorbike" {
			renderTemplate(w, "view/html/Path_after_hangman/Game_win/2realwinmotorbike.html", nil)
			return
		}

		if r.URL.Path == "/subway_path" {
			renderTemplate(w, "view/html/Path_after_hangman/subway_path.html", nil)
			return
		}

		if r.URL.Path == "/won" {
			renderTemplate(w, "view/html/won.html", nil)
			return
		}

		if r.URL.Path == "/2_north_subway_path" {
			renderTemplate(w, "view/html/Path_after_hangman/2_north_subway_path.html", nil)
			return
		}

		if r.URL.Path == "/2_south_subway_path" {
			renderTemplate(w, "view/html/Path_after_hangman/2_south_subway_path.html", nil)
			return
		}

		if r.URL.Path == "/3_hide_subway_path" {
			renderTemplate(w, "view/html/Path_after_hangman/3_hide_subway_path.html", nil)
			return
		}

		if r.URL.Path == "/winsubway4" {
			renderTemplate(w, "view/html/Path_after_hangman/Game_win/winsubway4.html", nil)
			return
		}

		if r.URL.Path == "/winsubway5" {
			renderTemplate(w, "view/html/Path_after_hangman/Game_win/winsubway5.html", nil)
			return
		}

		if r.URL.Path == "/winsubway6" {
			renderTemplate(w, "view/html/Path_after_hangman/Game_win/winsubway6.html", nil)
			return
		}

		if r.URL.Path == "/surrounded_subwayLost" {
			renderTemplate(w, "view/html/Path_after_hangman/Game_lost/surrounded_subwayLost.html", nil)
			return
		}

		if r.URL.Path == "/1_steal_lost_path" {
			renderTemplate(w, "view/html/Path_after_loose/1_steal_lost_path.html", nil)
			return
		}

		if r.URL.Path == "/1_war_lost_path" {
			renderTemplate(w, "view/html/Path_after_loose/1_war_lost_path.html", nil)
			return
		}

		if r.URL.Path == "/2_hit_lost_path" {
			renderTemplate(w, "view/html/Path_after_loose/2_hit_lost_path.html", nil)
			return
		}

		if r.URL.Path == "/2_stop_lost_path" {
			renderTemplate(w, "view/html/Path_after_loose/2_stop_lost_path.html", nil)
			return
		}

		if r.URL.Path == "/2_shothun_lost_path" {
			renderTemplate(w, "view/html/Path_after_loose/2_shothun_lost_path.html", nil)
			return
		}

		if r.URL.Path == "/2_katana_lost_path" {
			renderTemplate(w, "view/html/Path_after_loose/2_katana_lost_path.html", nil)
			return
		}

		if r.URL.Path == "/lostWin_path" {
			renderTemplate(w, "view/html/Path_after_hangman/Game_win/lostWin_path.html", nil)
			return
		}

		if r.URL.Path == "/restart" {
			g.restartGame()
			g.startGame()
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	r.ParseForm()
	guess := r.FormValue("guess")
	selectedLetter := r.FormValue("selectedLetter")

	// Récupération des lettres de A à Z depuis le formulaire.
	A := r.Form.Get("A")
	B := r.Form.Get("B")
	C := r.Form.Get("C")
	D := r.Form.Get("D")
	E := r.Form.Get("E")
	F := r.Form.Get("F")
	G := r.Form.Get("G")
	H := r.Form.Get("H")
	I := r.Form.Get("I")
	J := r.Form.Get("J")
	K := r.Form.Get("K")
	L := r.Form.Get("L")
	M := r.Form.Get("M")
	N := r.Form.Get("N")
	O := r.Form.Get("O")
	P := r.Form.Get("P")
	Q := r.Form.Get("Q")
	R := r.Form.Get("R")
	S := r.Form.Get("S")
	T := r.Form.Get("T")
	U := r.Form.Get("U")
	V := r.Form.Get("V")
	W := r.Form.Get("W")
	X := r.Form.Get("X")
	Y := r.Form.Get("Y")
	Z := r.Form.Get("Z")

	// Concaténation des lettres pour obtenir les lettres déjà devinées.
	g.guessletters = A + B + C + D + E + F + G + H + I + J + K + L + M + N + O + P + Q + R + S + T + U + V + W + X + Y + Z

	if r.Method == http.MethodGet {
		// Gestion des requêtes GET
		if r.URL.Path == "/" && !g.gameStarted {
			renderTemplate(w, "view/html/menu.html", nil)
			return
		}

		if r.URL.Path == "/restart" {
			g.restartGame()
			g.startGame()
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		if !g.gameStarted {
			g.startGame()
		}

		// Affichage de la page de défaite.
		if g.attempts == 0 {
			renderTemplate(w, "view/html/lost.html", templateData{
				Theword:        g.currentWord,
				WordDisplay:    g.displayWord(),
				GuessedLetters: strings.Join(g.guessedLetters, ", "),
				Attempts:       g.attempts,
				//HangmanImage: fmt.Sprintf("/assets/css/Tree/arbre%d.png", g.hangmanStep), previous way to change image every fault
			})
			return
		}
		// Affichage de la page de victoire.
		if strings.Replace(g.displayWord(), " ", "", -1) == g.currentWord {
			renderTemplate(w, "view/html/won.html", nil)
			return
		}
		// Affichage de la page principale du jeu.
		renderTemplate(w, "view/html/index.html", templateData{
			WordDisplay:    g.displayWord(),
			GuessedLetters: strings.Join(g.guessedLetters, ", "),
			Attempts:       g.attempts,
			//HangmanImage: fmt.Sprintf("/assets/css/Tree/arbre%d.png", g.hangmanStep), previous way to change image every fault
		})

		// Gestion des requêtes POST
	} else if r.Method == http.MethodPost {
		if g.attempts > 0 {
			var letterToCheck string

			if selectedLetter != "" {
				letterToCheck = selectedLetter
			} else if guess != "" && !contains(g.guessedLetters, guess) {
				letterToCheck = guess
			}

			// Comparaison et révélation de la lettre.
			if letterToCheck != "" {
				g.compareAndReveal(letterToCheck)

				if !strings.Contains(g.currentWord, letterToCheck) {
					// Réduction du nombre de tentatives et avancement de l'étape du pendu.
					g.attempts--
					g.hangmanStep++
				}
			}
		}

		// Redirection vers la page de victoire.
		if strings.Replace(g.displayWord(), " ", "", -1) == g.currentWord {
			http.Redirect(w, r, "/restart-won", http.StatusSeeOther)
			return
		}
		// Redirection vers la page de défaite.
		if g.attempts == 0 {
			http.Redirect(w, r, "/restart-lost", http.StatusSeeOther)
			return
		}

		// Affichage de la page principale du jeu.
		renderTemplate(w, "view/html/index.html", templateData{
			WordDisplay:    g.displayWord(),
			GuessedLetters: strings.Join(g.guessedLetters, ", "),
			Attempts:       g.attempts,
			//HangmanImage: fmt.Sprintf("/assets/css/Tree/arbre%d.png", g.hangmanStep), previous way to change image every fault
		})
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// renderTemplate rend le template HTML avec les données fournies.
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}
