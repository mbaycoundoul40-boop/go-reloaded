package main

import "strings"

const punctuation = ".,!?:;"

// fixArticles remplace "a" par "an" devant une voyelle ou un h.
func fixArticles(words []string) []string {
	for i := 0; i < len(words)-1; i++ {
		if words[i] != "a" && words[i] != "A" {
			continue
		}
		next := strings.ToLower(words[i+1])
		if next != "" && strings.ContainsRune("aeiouh", rune(next[0])) {
			words[i] += "n"
		}
	}
	return words
}

// fixQuotes colle chaque paire de ' aux mots qu'elle entoure.
func fixQuotes(words []string) []string {
	var result []string
	open := false
	pending := false

	for _, w := range words {
		if w != "'" {
			if pending {
				w = "'" + w
				pending = false
			}
			result = append(result, w)
			continue
		}

		if !open {
			open = true
			pending = true
		} else {
			open = false
			if len(result) > 0 {
				result[len(result)-1] += "'"
			} else {
				result = append(result, "'")
			}
		}
	}

	if pending {
		result = append(result, "'")
	}
	return result
}

// fixPunctuation colle la ponctuation au mot précédent.
func fixPunctuation(words []string) []string {
	var result []string

	for _, w := range words {
		end := 0
		for end < len(w) && strings.ContainsRune(punctuation, rune(w[end])) {
			end++
		}
		lead, rest := w[:end], w[end:]

		if lead != "" && len(result) > 0 {
			result[len(result)-1] += lead
		} else if lead != "" {
			rest = w
		}

		if rest != "" {
			result = append(result, rest)
		}
	}
	return result
}
