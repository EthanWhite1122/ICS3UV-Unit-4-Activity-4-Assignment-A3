//Author: Ethan White
//Version: 1.0.0
//Date:2025-12-11
//Fileoverview This program will take a inputted sentence and tell you if a inputted word is in the sentence

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a sentence: ")
	sentence, _ := reader.ReadString('\n')
	sentence = strings.TrimSpace(sentence)

	fmt.Print("Enter a word to search: ")
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word)

	// Convert to lowercase for case-insensitive search
	sentenceLower := strings.ToLower(sentence)
	wordLower := strings.ToLower(word)

	if strings.Contains(sentenceLower, wordLower) {
		fmt.Println("The word exists in the sentence.")
	} else {
		fmt.Println("The word does not exist in the sentence.")
	}
}

