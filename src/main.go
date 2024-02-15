package main

import (
	"fmt"
	"net/http"
	e "src/game"
)

func main() {
	var g e.Game

	fs := http.FileServer(http.Dir("./view/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Gestion des requêtes pour le jeu du pendu.
	http.HandleFunc("/", g.HangmanHandler)
	fmt.Println("Server http://localhost:5050")
	http.ListenAndServe(":5050", nil)
}
