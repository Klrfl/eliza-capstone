package main

import (
	"strings"
)

// reflections maps first-person words to second-person equivalents and vice
// versa, as specified in the original ELIZA paper. Applied when a captured
// wildcard group is echoed back to the user.
var reflections = map[string]string{
	"am":     "are",
	"was":    "were",
	"i":      "you",
	"i'd":    "you'd",
	"i've":   "you've",
	"i'll":   "you'll",
	"i'm":    "you're",
	"my":     "your",
	"me":     "you",
	"myself": "yourself",
	"are":    "am",
	"were":   "was",
	"you":    "I",
	"you'd":  "I'd",
	"you've": "I've",
	"you'll": "I'll",
	"you're": "I'm",
	"your":   "my",
	"yours":  "mine",
}

// reflect applies pronoun reflection to a captured phrase so it reads
// naturally when inserted into a response template.
// e.g. "i am sad" → "you are sad"
func reflect(phrase string) string {
	words := strings.Fields(phrase)
	for i, w := range words {
		lower := strings.ToLower(w)
		if mapped, ok := reflections[lower]; ok {
			// Preserve leading capitalisation if the original word was capitalised.
			if len(w) > 0 && w[0] >= 'A' && w[0] <= 'Z' {
				words[i] = strings.ToUpper(mapped[:1]) + mapped[1:]
			} else {
				words[i] = mapped
			}
		}
	}
	return strings.Join(words, " ")
}

// normalize lower-cases the input and strips punctuation that would
// interfere with pattern matching, while preserving apostrophes in
// contractions (e.g. "can't", "i'm").
func normalize(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	prev := ' '
	for _, r := range strings.ToLower(s) {
		switch {
		case r >= 'a' && r <= 'z':
			b.WriteRune(r)
			prev = r
		case r == '\'':
			// Keep apostrophe only inside words (contractions).
			if prev != ' ' {
				b.WriteRune(r)
			}
			prev = r
		case r == ' ' || r == '\t':
			if prev != ' ' {
				b.WriteByte(' ')
			}
			prev = ' '
		default:
			// Replace other punctuation with a space separator.
			if prev != ' ' {
				b.WriteByte(' ')
			}
			prev = ' '
		}
	}
	return strings.TrimSpace(b.String())
}
