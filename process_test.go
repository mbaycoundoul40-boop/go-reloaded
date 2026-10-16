package main

import "testing"

func TestProcess(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"hex", "1E (hex) files were added", "30 files were added"},
		{"bin", "It has been 10 (bin) years", "It has been 2 years"},
		{"up", "Ready, set, go (up) !", "Ready, set, GO!"},
		{"low", "I should stop SHOUTING (low)", "I should stop shouting"},
		{"cap", "Welcome to the Brooklyn bridge (cap)", "Welcome to the Brooklyn Bridge"},
		{"up avec nombre", "This is so exciting (up, 2)", "This is SO EXCITING"},
		{"ponctuation", "I was sitting over there ,and then BAMM !!", "I was sitting over there, and then BAMM!!"},
		{"groupe de ponctuation", "I was thinking ... You were right", "I was thinking... You were right"},
		{"apostrophes un mot", "I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},
		{"apostrophes plusieurs mots",
			"As Elton John said: ' I am the most well-known homosexual in the world '",
			"As Elton John said: 'I am the most well-known homosexual in the world'"},
		{"a devient an", "There it was. A amazing rock!", "There it was. An amazing rock!"},
		{"exemple 2", "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			"Simply add 66 and 2 and you will see the result is 68."},
		{"exemple 3", "There is no greater agony than bearing a untold story inside you.",
			"There is no greater agony than bearing an untold story inside you."},
		{"exemple 4", "Punctuation tests are ... kinda boring ,what do you think ?",
			"Punctuation tests are... kinda boring, what do you think?"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Process(tt.input)
			if got != tt.want {
				t.Errorf("\nentrée : %q\nobtenu : %q\nattendu: %q", tt.input, got, tt.want)
			}
		})
	}
}
