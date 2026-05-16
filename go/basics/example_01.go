// Package challenge6 contains the solution for Challenge 6.
package challenge6

import (
	"strings"
	"regexp"
)

// CountWordFrequency takes a string containing multiple words and returns
// a map where each key is a word and the value is the number of times that
// word appears in the string. The comparison is case-insensitive.
//
// Words are defined as sequences of letters and digits.
// All words are converted to lowercase before counting.
// All punctuation, spaces, and other non-alphanumeric characters are ignored.
//
// For example:
// Input: "The quick brown fox jumps over the lazy dog."
// Output: map[string]int{"the": 2, "quick": 1, "brown": 1, "fox": 1, "jumps": 1, "over": 1, "lazy": 1, "dog": 1}
func CountWordFrequency(text string) map[string]int {
	// Your implementation here
	
	//1. Convertimos el texto a minúscula
	text = strings.ToLower(text)
	
	//2. Eliminamos los apostrofes para que las palabras como "let's" se unan
	text = strings.ReplaceAll(text, "'", "")
	text = strings.ReplaceAll(text, "`", "")
	
	//3. Filtramos los caracteres de puntuación usando regex
 	re := regexp.MustCompile(`\w+`)
 	
 	//4. Extraemos las palabras y guardamos en el slice words
 	words := re.FindAllString(text, -1)
 	
 	//Creamos el mapa 
 	counts := make(map[string]int)
 	
 	for _, word := range words{
 	    counts[word]++
 	}
	
	return counts
} 