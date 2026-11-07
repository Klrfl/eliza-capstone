package main

import "testing"

// exchange is a single prompt/response pair.
type exchange struct {
	prompt   string
	response string
}

// TestConversations runs each exchange through a fresh ELIZA instance and
// checks that the output exactly matches the expected response.
func TestConversations(t *testing.T) {
	conversations := []struct {
		name      string
		exchanges []exchange
	}{
		{
			name: "opening",
			exchanges: []exchange{
				{"Hello", "How do you do. Please state your problem."},
			},
		},
		{
			name: "CACM paper exchange",
			exchanges: []exchange{
				{"Well, my boyfriend made me come here.", "Your boyfriend made you come here?"},
			},
		},
		{
			name: "i am",
			exchanges: []exchange{
				{"I am very unhappy", "Did you come to me because you are very unhappy?"},
			},
		},
		{
			// "like" (priority 10) outranks "i" (priority 0), redirects to =alike.
			name: "i feel like — like keyword wins",
			exchanges: []exchange{
				{"I feel like I am lost", "In what way?"},
			},
		},
		{
			// "cannot" is the pre-transformed form; "i" keyword fires, "* i can't *"
			// pattern no longer matches because the contraction was expanded.
			// "* i *" catches it instead.
			name: "i can't contraction",
			exchanges: []exchange{
				{"I can't sleep at night", "You say you cannot sleep at night?"},
			},
		},
		{
			// "do not" is the pre-transformed form; "* i *" catches "i do not know..."
			// Fresh instance: "* i *" index 0 → "You say you (2)?"
			name: "i don't contraction",
			exchanges: []exchange{
				{"I don't know what to do", "You say you do not know what to do?"},
			},
		},
		{
			// "i am" is the pre-transformed form; "* i am *" decomp fires.
			// Fresh instance: "* i am *" index 0 → "Did you come to me because you are (2)?"
			name: "i'm contraction",
			exchanges: []exchange{
				{"I'm feeling lost", "Did you come to me because you are feeling lost?"},
			},
		},
		{
			name: "family - mother generic",
			exchanges: []exchange{
				{"My mother never listened to me", "Tell me more about your mother."},
			},
		},
		{
			name: "family - mother with tail capture",
			exchanges: []exchange{
				// First call → index 0: "Tell me more about your mother."
				// Second call → index 1: "Who else in your family (2)?" with reflected tail.
				{"My mother takes care of me", "Tell me more about your mother."},
				{"My mother takes care of me", "Who else in your family takes care of you?"},
			},
		},
		{
			name: "remember",
			exchanges: []exchange{
				{"Do you remember me", "Did you think I would forget you?"},
			},
		},
		{
			name: "computer",
			exchanges: []exchange{
				{"Computer says no", "Do computers worry you?"},
			},
		},
		{
			name: "everyone",
			exchanges: []exchange{
				{"I think everyone hates me", "Really, everyone?"},
			},
		},
		{
			name: "perhaps",
			exchanges: []exchange{
				{"Perhaps I need help", "You don't seem quite certain."},
			},
		},
		{
			name: "dreamed",
			exchanges: []exchange{
				{"I dreamed about flying", "Really, about flying?"},
			},
		},
		{
			// "i" (priority 0) and "because" (priority 0) tie; "i" wins because it
			// appears later in the script and overwrites best in the scan.
			// Fresh instance: "* i *" index 0 → "You say you (2)?"
			name: "because with i",
			exchanges: []exchange{
				{"Because I said so", "You say you said so?"},
			},
		},
		{
			name: "quit",
			exchanges: []exchange{
				{"quit", "Goodbye. It was nice talking to you."},
			},
		},
		{
			// Verify round-robin: same input on a fresh instance gives response[0],
			// second call gives response[1].
			name: "round-robin cycles responses",
			exchanges: []exchange{
				{"perhaps", "You don't seem quite certain."},
				{"perhaps", "Why the uncertain tone?"},
			},
		},
		{
			name: "multi-turn: i am then family",
			exchanges: []exchange{
				// Fresh instance: "* i am *" index 0.
				{"I am feeling anxious", "Did you come to me because you are feeling anxious?"},
				// "father" keyword (priority 100), pattern "* father *", index 0.
				{"My father was strict", "Tell me more about your father."},
			},
		},
	}

	for _, conv := range conversations {
		t.Run(conv.name, func(t *testing.T) {
			e := New()
			for _, ex := range conv.exchanges {
				got := e.Respond(ex.prompt)
				if got != ex.response {
					t.Errorf("prompt:   %q\ngot:      %q\nexpected: %q", ex.prompt, got, ex.response)
				}
			}
		})
	}
}
