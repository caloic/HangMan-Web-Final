package src

import (
	"bufio"
	"log"
	"math/rand"
	"os"
)

// Lis le document txt
func ReadFile() []string {
	var word []string
	f, err := os.Open("data/word.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word = append(word, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return word
}

// Choisi un mot au hasard parmis la liste de mot
func ChoseRandomWord(wordlist []string) []rune {
	randomIndex := rand.Intn(len(wordlist))
	wordInArray := wordlist[randomIndex]
	return []rune(wordInArray)
}
