/** 
 * @author Ethan White
 * @version 1.0.0
 * @date 2025-12-11
 * @fileoverview This program will take a inputted sentence and tell you if a inputted word is in the sentence
*/


// Word search program

let sentence = prompt("Enter a sentence: ")!;
let word = prompt("Enter a word to search: ")!;

// Convert to lowercase for case-insensitive search
let sentenceLower = sentence.toLowerCase();
let wordLower = word.toLowerCase();

if (sentenceLower.includes(wordLower)) {
  console.log("The word exists in the sentence.");
} else {
  console.log("The word does not exist in the sentence.");
}

console.log("\nDone.")