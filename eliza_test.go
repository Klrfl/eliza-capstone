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
			name: "i don't contraction",
			exchanges: []exchange{
				{"I don't know what to do", "Can you elaborate on that?"},
			},
		},
		{
			// "i am" is the pre-transformed form; "* i am *" decomp fires.
			name: "i'm contraction",
			exchanges: []exchange{
				{"I'm feeling lost", "How long have you been feeling lost?"},
			},
		},
		{
			name: "family - mother",
			exchanges: []exchange{
				{"My mother never listened to me", "Tell me more about your family."},
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
			// "i" (priority 0) outranks "because" (priority 0) but appears later
			// in the script scan; actual winner is "i" via "* i *".
			name: "because with i",
			exchanges: []exchange{
				{"Because I said so", "Do you say you said so for some special reason?"},
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
				{"perhaps", "Why the uncertain tone?"},
				{"perhaps", "Can't you be more positive?"},
			},
		},
		{
			name: "multi-turn: i am then family",
			exchanges: []exchange{
				// "i am" decomp cycles — second response on a fresh ELIZA instance.
				{"I am feeling anxious", "Do you believe it is normal to be feeling anxious?"},
				// "father" (priority 100) wins, redirects to =family.
				// The family rule's round-robin index is already at 1 (shared via ruleMap)
				// so the second family response fires.
				{"My father was strict", "Who else in your family?"},
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
