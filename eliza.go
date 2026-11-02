package main

import (
	"fmt"
	"regexp"
	"strings"
)

// ELIZA is the stateful conversation engine.
type ELIZA struct {
	rules    []Rule          // script rules, keyed by keyword
	ruleMap  map[string]*Rule // fast lookup by keyword
}

// New builds an ELIZA instance from the bundled script.
// Duplicate keyword entries are merged so later Decomps are appended.
func New() *ELIZA {
	e := &ELIZA{
		ruleMap: make(map[string]*Rule),
	}
	for i := range script {
		r := &script[i]
		if existing, ok := e.ruleMap[r.Keyword]; ok {
			existing.Decomps = append(existing.Decomps, r.Decomps...)
		} else {
			e.ruleMap[r.Keyword] = r
			e.rules = append(e.rules, *r)
		}
	}
	return e
}

// Respond generates ELIZA's reply to the given user input.
func (e *ELIZA) Respond(input string) string {
	text := preTransform(normalize(input))

	// Check for quit signal.
	if text == "quit" || text == "bye" || text == "goodbye" || text == "exit" {
		return "Goodbye. It was nice talking to you."
	}

	// Find the highest-priority rule whose keyword appears in the input.
	var best *Rule
	for i := range script {
		r := &script[i]
		if r.Keyword == "xnone" {
			continue
		}
		if containsWord(text, r.Keyword) {
			if best == nil || r.Priority > best.Priority {
				best = r
			}
		}
	}

	if best == nil {
		best = e.ruleMap["xnone"]
	}

	return e.applyRule(best, text, 0)
}

// applyRule tries each decomposition pattern of the rule in order and
// returns the assembled response. depth guards against infinite redirect loops.
func (e *ELIZA) applyRule(r *Rule, text string, depth int) string {
	if depth > 10 {
		return "I see."
	}
	for i := range r.Decomps {
		d := &r.Decomps[i]
		captures, ok := matchDecomp(d.Pattern, text)
		if !ok {
			continue
		}

		resp := d.Responses[d.index%len(d.Responses)]
		d.index++

		// Handle redirect responses (=keyword).
		if strings.HasPrefix(resp, "=") {
			target := resp[1:]
			if tr, ok := e.ruleMap[target]; ok {
				return e.applyRule(tr, text, depth+1)
			}
			// Redirect target not found; fall through to next decomp.
			continue
		}

		return assemble(resp, captures)
	}

	// No decomp matched — fall back to xnone.
	if r.Keyword != "xnone" {
		return e.applyRule(e.ruleMap["xnone"], text, depth+1)
	}
	return "I am not sure I understand you fully."
}

// containsWord reports whether word appears as a whole word inside s.
func containsWord(s, word string) bool {
	// Multi-word keywords (e.g. "can you") need a simple substring check
	// bounded by spaces or start/end.
	re := regexp.MustCompile(`(?i)(^|\s)` + regexp.QuoteMeta(word) + `(\s|$)`)
	return re.MatchString(s)
}

// matchDecomp matches the ELIZA decomposition pattern against text.
// The pattern uses `*` to denote "zero or more words" (a capture group).
// Returns the captured substrings (reflected) and true on success.
//
// Pattern example:  "* i am *"
// Text:             "i think i am very sad"
// captures[0] → "" (before "i am"), captures[1] → "very sad"
func matchDecomp(pattern, text string) ([]string, bool) {
	if pattern == "*" {
		return []string{text}, true
	}

	// Build a regex from the ELIZA pattern.
	// Each `*` becomes a non-greedy capture group that matches any words.
	parts := strings.Split(pattern, "*")
	var sb strings.Builder
	sb.WriteString(`^\s*`)
	for i, p := range parts {
		sb.WriteString(regexp.QuoteMeta(strings.TrimSpace(p)))
		if i < len(parts)-1 {
			sb.WriteString(`\s*(.*?)\s*`)
		}
	}
	sb.WriteString(`\s*$`)

	re, err := regexp.Compile(`(?i)` + sb.String())
	if err != nil {
		return nil, false
	}

	m := re.FindStringSubmatch(text)
	if m == nil {
		return nil, false
	}

	// m[0] is full match; m[1..] are capture groups — apply reflection.
	captures := make([]string, len(m)-1)
	for i, c := range m[1:] {
		captures[i] = reflect(strings.TrimSpace(c))
	}
	return captures, true
}

// assemble substitutes (1), (2), … placeholders in the response template
// with the corresponding reflected capture groups.
func assemble(template string, captures []string) string {
	result := template
	for i, c := range captures {
		placeholder := fmt.Sprintf("(%d)", i+1)
		result = strings.ReplaceAll(result, placeholder, c)
	}
	// Clean up any leftover placeholders.
	result = regexp.MustCompile(`\(\d+\)`).ReplaceAllString(result, "")
	return strings.TrimSpace(result)
}
