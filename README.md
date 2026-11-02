# ELIZA

A Go implementation of the ELIZA conversational program, running the original DOCTOR script from Joseph Weizenbaum's 1966 paper *ELIZA -- A Computer Program For the Study of Natural Language Communication Between Man And Machine*.

ELIZA simulates a Rogerian psychotherapist. It does not understand language; it reflects your words back at you using pattern matching and a fixed set of rules. The effect is surprisingly convincing.

## Background

ELIZA was written at MIT between 1964 and 1966. The DOCTOR script is the most famous of its rule sets. It works by scanning user input for keywords, ranking them by priority, matching the input against decomposition patterns, and filling response templates with reflected fragments of what the user said.

This implementation bundles the original DOCTOR ruleset as Go data structures. No external dependencies, no network access, no learned behavior.

## Building

You need Go 1.18 or later.

```bash
git clone <repo>
cd eliza
go build -o eliza .
```

Or run without building:

```bash
go run .
```

## Usage

### Interactive mode

```bash
./eliza
```

ELIZA greets you and waits for input. Type a sentence and press Enter. Type `quit`, `bye`, or `goodbye` to end the session. You can also press Ctrl-D.

```
ELIZA: How do you do. Please state your problem.

You:   I am very unhappy these days
ELIZA: Did you come to me because you are very unhappy these days?

You:   My mother never understood me
ELIZA: Tell me more about your family.

You:   I can't seem to make any friends
ELIZA: How do you know you can't seem to make any friends?

You:   quit
ELIZA: Goodbye. It was nice talking to you.
```

### Single-shot mode

Use the `--once` flag to send a single message and print the response. Useful for scripting or quick tests.

```bash
./eliza --once "I think everyone hates me"
# Really, everyone?
```

## How it works

When you send a message, ELIZA:

1. Normalizes the input -- lowercases it and removes punctuation while keeping contractions like `can't` and `i'm`.
2. Scans the input for keywords from the DOCTOR script, each with an assigned priority. `family` words like `mother` and `father` rank highest (100); general words like `I` and `you` rank low (0).
3. Picks the highest-priority matching keyword and tries its decomposition patterns in order. Patterns use `*` as a wildcard, e.g. `* i am *`.
4. Applies pronoun reflection to any captured text: `I am sad` becomes `you are sad` before being inserted into a response.
5. Fills numbered placeholders like `(1)` and `(2)` in the response template with the reflected captures.
6. Some keywords redirect to another keyword's responses using `=keyword` syntax, e.g. `mother` redirects to `=family`.

Responses within a rule are cycled round-robin, so the same input gives a different reply on subsequent turns.

## Project structure

```
eliza/
├── main.go       -- REPL loop and --once flag
├── eliza.go      -- pattern matching and response assembly
├── transform.go  -- input normalization and pronoun reflection
└── script.go     -- bundled DOCTOR ruleset
```

## Reference

Weizenbaum, J. (1966). ELIZA -- A computer program for the study of natural language communication between man and machine. *Communications of the ACM*, 9(1), 36-45.
