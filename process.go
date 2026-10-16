package main

import "strings"

// Process applique toutes les modifications au texte, ligne par ligne.
func Process(text string) string {
	text = strings.ReplaceAll(text, "\r\n", "\n")
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		lines[i] = processLine(line)
	}
	return strings.Join(lines, "\n")
}

// processLine applique les étapes dans l'ordre sur une seule ligne.
func processLine(line string) string {
	line = normalizeMarkers(line)
	words := strings.Fields(line)

	words = applyMarkers(words)
	words = fixArticles(words)
	words = fixQuotes(words)
	words = fixPunctuation(words)

	return strings.Join(words, " ")
}
