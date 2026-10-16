package main

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// markerInText repère un marqueur dans le texte brut, même avec des espaces : "(up, 2)".
var markerInText = regexp.MustCompile(`\(\s*(hex|bin|up|low|cap)\s*(?:,\s*(-?\d+)\s*)?\)`)

// markerWord repère un marqueur déjà normalisé : "(up)" ou "(up,2)".
var markerWord = regexp.MustCompile(`^\((hex|bin|up|low|cap)(?:,(-?\d+))?\)$`)

// normalizeMarkers écrit chaque marqueur sans espace intérieur et l'isole.
func normalizeMarkers(line string) string {
	return markerInText.ReplaceAllStringFunc(line, func(m string) string {
		parts := markerInText.FindStringSubmatch(m)
		if parts[2] == "" {
			return " (" + parts[1] + ") "
		}
		return " (" + parts[1] + "," + parts[2] + ") "
	})
}

// applyMarkers applique chaque marqueur aux mots qui le précèdent, puis le supprime.
func applyMarkers(words []string) []string {
	var result []string

	for _, w := range words {
		parts := markerWord.FindStringSubmatch(w)
		if parts == nil {
			result = append(result, w)
			continue
		}

		kind := parts[1]
		count := 1
		if parts[2] != "" {
			count, _ = strconv.Atoi(parts[2])
		}

		for i := len(result) - 1; i >= 0 && count > 0; i-- {
			if !hasLetterOrDigit(result[i]) {
				continue
			}
			result[i] = transform(result[i], kind)
			count--
		}
	}
	return result
}

// transform applique une modification à un mot.
func transform(word, kind string) string {
	switch kind {
	case "hex":
		return convertBase(word, 16)
	case "bin":
		return convertBase(word, 2)
	case "up":
		return strings.ToUpper(word)
	case "low":
		return strings.ToLower(word)
	case "cap":
		return capitalize(word)
	}
	return word
}

// convertBase convertit un nombre en base 16 ou 2 en décimal.
func convertBase(word string, base int) string {
	n, err := strconv.ParseInt(word, base, 64)
	if err != nil {
		return word
	}
	return strconv.FormatInt(n, 10)
}

// capitalize met la première lettre en majuscule et le reste en minuscules.
func capitalize(word string) string {
	runes := []rune(strings.ToLower(word))
	for i, r := range runes {
		if unicode.IsLetter(r) {
			runes[i] = unicode.ToUpper(r)
			break
		}
	}
	return string(runes)
}

// hasLetterOrDigit indique si le mot contient au moins une lettre ou un chiffre.
func hasLetterOrDigit(word string) bool {
	for _, r := range word {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return true
		}
	}
	return false
}
