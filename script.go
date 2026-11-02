package main

// Rule represents a single keyword entry from the DOCTOR script.
// Priority is used to select among multiple matching keywords — higher wins.
type Rule struct {
	Keyword    string
	Priority   int
	Decomps    []Decomp
}

// Decomp is a decomposition pattern paired with a list of reassembly templates.
// The pattern uses `*` as a wildcard capture group (following original ELIZA notation).
// Responses are cycled through in order.
type Decomp struct {
	Pattern   string
	Responses []string
	index     int // tracks which response to use next (round-robin)
}

// script is the bundled DOCTOR ruleset, faithfully derived from
// Weizenbaum's 1966 CACM paper "ELIZA—A Computer Program For the
// Study of Natural Language Communication Between Man And Machine."
var script = []Rule{
	{
		Keyword:  "xnone",
		Priority: 0,
		Decomps: []Decomp{
			{
				Pattern: "*",
				Responses: []string{
					"I am not sure I understand you fully.",
					"Please go on.",
					"What does that suggest to you?",
					"Do you feel strongly about discussing such things?",
					"That is interesting. Please continue.",
					"Tell me more about that.",
				},
			},
		},
	},
	{
		Keyword:  "sorry",
		Priority: 0,
		Decomps: []Decomp{
			{
				Pattern: "*",
				Responses: []string{
					"Please don't apologise.",
					"Apologies are not necessary.",
					"What feelings do you have when you apologise?",
					"I've told you that apologies are not required.",
				},
			},
		},
	},
	{
		Keyword:  "apologise",
		Priority: 0,
		Decomps: []Decomp{
			{Pattern: "*", Responses: []string{"=sorry"}},
		},
	},
	{
		Keyword:  "remember",
		Priority: 5,
		Decomps: []Decomp{
			{
				Pattern: "* i remember *",
				Responses: []string{
					"Do you often think of (2)?",
					"Does thinking of (2) bring anything else to mind?",
					"What else do you recollect?",
					"Why do you recollect (2) just now?",
					"What in the present situation reminds you of (2)?",
					"What is the connection between me and (2)?",
				},
			},
			{
				Pattern: "* do you remember *",
				Responses: []string{
					"Did you think I would forget (2)?",
					"Why do you think I should recall (2) now?",
					"What about (2)?",
					"=what",
					"You mentioned (2)?",
				},
			},
		},
	},
	{
		Keyword:  "if",
		Priority: 3,
		Decomps: []Decomp{
			{
				Pattern: "* if *",
				Responses: []string{
					"Do you think it's likely that (2)?",
					"Do you wish that (2)?",
					"What do you think about (2)?",
					"Really, if (2)?",
				},
			},
		},
	},
	{
		Keyword:  "dreamed",
		Priority: 4,
		Decomps: []Decomp{
			{
				Pattern: "* i dreamed *",
				Responses: []string{
					"Really, (2)?",
					"Have you ever fantasized (2) while you were awake?",
					"Have you dreamed (2) before?",
					"=dream",
				},
			},
		},
	},
	{
		Keyword:  "dream",
		Priority: 3,
		Decomps: []Decomp{
			{
				Pattern: "*",
				Responses: []string{
					"What does that dream suggest to you?",
					"Do you dream often?",
					"What persons appear in your dreams?",
					"Don't you believe that dream has something to do with your problem?",
				},
			},
		},
	},
	{
		Keyword:  "perhaps",
		Priority: 0,
		Decomps: []Decomp{
			{
				Pattern: "*",
				Responses: []string{
					"You don't seem quite certain.",
					"Why the uncertain tone?",
					"Can't you be more positive?",
					"You aren't sure?",
					"Don't you know?",
				},
			},
		},
	},
	{
		Keyword:  "name",
		Priority: 15,
		Decomps: []Decomp{
			{
				Pattern: "*",
				Responses: []string{
					"I am not interested in names.",
					"I've told you before, I don't care about names — please continue.",
				},
			},
		},
	},
	{
		Keyword:  "deutsch",
		Priority: 0,
		Decomps: []Decomp{
			{Pattern: "*", Responses: []string{"=français"}},
		},
	},
	{
		Keyword:  "français",
		Priority: 0,
		Decomps: []Decomp{
			{Pattern: "*", Responses: []string{"=italiano"}},
		},
	},
	{
		Keyword:  "italiano",
		Priority: 0,
		Decomps: []Decomp{
			{Pattern: "*", Responses: []string{"=español"}},
		},
	},
	{
		Keyword:  "español",
		Priority: 0,
		Decomps: []Decomp{
			{
				Pattern: "*",
				Responses: []string{
					"I am sorry, I speak only English.",
				},
			},
		},
	},
	{
		Keyword:  "hello",
		Priority: 0,
		Decomps: []Decomp{
			{
				Pattern: "*",
				Responses: []string{
					"How do you do. Please state your problem.",
				},
			},
		},
	},
	{
		Keyword:  "computer",
		Priority: 50,
		Decomps: []Decomp{
			{
				Pattern: "*",
				Responses: []string{
					"Do computers worry you?",
					"Why do you mention computers?",
					"What do you think machines have to do with your problem?",
					"Don't you think computers can help people?",
					"What about machines worries you?",
					"Say, do you have any psychological problems?",
				},
			},
		},
	},
	{
		Keyword:  "am",
		Priority: 0,
		Decomps: []Decomp{
			{
				Pattern: "* am i *",
				Responses: []string{
					"Do you believe you are (2)?",
					"Would you want to be (2)?",
					"Do you wish I would tell you you are (2)?",
					"What would it mean if you were (2)?",
					"=what",
				},
			},
			{
				Pattern: "* i am *",
				Responses: []string{"=i"},
			},
			{
				Pattern: "*",
				Responses: []string{
					"Why do you say 'am'?",
					"I don't understand that.",
				},
			},
		},
	},
	{
		Keyword:  "are",
		Priority: 0,
		Decomps: []Decomp{
			{
				Pattern: "* are you *",
				Responses: []string{
					"Why are you interested in whether I am (2) or not?",
					"Would you prefer if I weren't (2)?",
					"Perhaps I am (2) in your fantasies.",
					"Do you sometimes think I am (2)?",
					"=what",
				},
			},
			{
				Pattern: "* are *",
				Responses: []string{
					"Did you think they might not be (2)?",
					"Would you like it if they were not (2)?",
					"What if they were not (2)?",
					"Possibly they are (2).",
				},
			},
		},
	},
	{
		Keyword:  "your",
		Priority: 0,
		Decomps: []Decomp{
			{
				Pattern: "* your *",
				Responses: []string{
					"Why are you concerned over my (2)?",
					"What about your own (2)?",
					"Are you worried about someone else's (2)?",
					"Really, my (2)?",
				},
			},
		},
	},
	{
		Keyword:  "was",
		Priority: 2,
		Decomps: []Decomp{
			{
				Pattern: "* was i *",
				Responses: []string{
					"What if you were (2)?",
					"Do you think you were (2)?",
					"Were you (2)?",
					"What would it mean if you were (2)?",
					"What does '(2)' suggest to you?",
					"=what",
				},
			},
			{
				Pattern: "* i was *",
				Responses: []string{
					"Were you really?",
					"Why do you tell me you were (2) just now?",
					"Perhaps I already know you were (2).",
				},
			},
			{
				Pattern: "* was you *",
				Responses: []string{
					"Why do you ask?",
					"What if I was (2)?",
					"=what",
				},
			},
		},
	},
	{
		Keyword:  "i",
		Priority: 0,
		Decomps: []Decomp{
			{
				Pattern: "* i want *",
				Responses: []string{
					"What would it mean to you if you got (2)?",
					"Why do you want (2)?",
					"Suppose you got (2) soon.",
					"What if you never got (2)?",
					"What would getting (2) mean to you?",
				},
			},
			{
				Pattern: "* i am *",
				Responses: []string{
					"Did you come to me because you are (2)?",
					"How long have you been (2)?",
					"Do you believe it is normal to be (2)?",
					"Do you enjoy being (2)?",
				},
			},
			{
				Pattern: "* i'm *",
				Responses: []string{
					"How does being (2) make you feel?",
					"Do you enjoy being (2)?",
					"Why do you tell me you're (2)?",
					"Why do you think you're (2)?",
				},
			},
			{
				Pattern: "* i can't *",
				Responses: []string{
					"How do you know you can't (2)?",
					"Have you tried?",
					"Perhaps you could (2) now.",
					"Do you really want to be able to (2)?",
				},
			},
			{
				Pattern: "* i don't *",
				Responses: []string{
					"Don't you really (2)?",
					"Why don't you (2)?",
					"Do you wish to be able to (2)?",
					"Does that trouble you?",
				},
			},
			{
				Pattern: "* i feel *",
				Responses: []string{
					"Tell me more about such feelings.",
					"Do you often feel (2)?",
					"Do you enjoy feeling (2)?",
					"Of what does feeling (2) remind you?",
				},
			},
			{
				Pattern: "* i think *",
				Responses: []string{
					"Do you really think so?",
					"But you are not sure you (2)?",
					"Do you doubt you (2)?",
				},
			},
			{
				Pattern: "* i have *",
				Responses: []string{
					"Why do you tell me that you've (2)?",
					"Have you really (2)?",
					"Now that you have (2), what will you do next?",
				},
			},
			{
				Pattern: "* i would *",
				Responses: []string{
					"Could you explain why you would (2)?",
					"Who else knows you would (2)?",
					"Why would you (2)?",
					"What makes you think you would (2)?",
				},
			},
			{
				Pattern: "* i *",
				Responses: []string{
					"You say you (2)?",
					"Can you elaborate on that?",
					"Do you say you (2) for some special reason?",
					"That's quite interesting.",
				},
			},
		},
	},
	{
		Keyword:  "you",
		Priority: 0,
		Decomps: []Decomp{
			{
				Pattern: "* you remind me of *",
				Responses: []string{"=alike"},
			},
			{
				Pattern: "* you are *",
				Responses: []string{
					"What makes you think I am (2)?",
					"Does it please you to believe I am (2)?",
					"Do you sometimes wish you were (2)?",
					"Perhaps you would like to be (2).",
				},
			},
			{
				Pattern: "* you * me *",
				Responses: []string{
					"Why do you think I (2) you?",
					"You like to think I (2) you — don't you?",
					"What makes you think I (2) you?",
					"Really, I (2) you?",
					"Does it please you to believe I (2) you?",
				},
			},
			{
				Pattern: "* you *",
				Responses: []string{
					"We were discussing you — not me.",
					"Oh, I (2)?",
					"You're not really talking about me — are you?",
					"What are your feelings now?",
				},
			},
		},
	},
	{
		Keyword:  "yes",
		Priority: 0,
		Decomps: []Decomp{
			{
				Pattern: "*",
				Responses: []string{
					"You seem quite positive.",
					"You are sure.",
					"I see.",
					"I understand.",
				},
			},
		},
	},
	{
		Keyword:  "no",
		Priority: 0,
		Decomps: []Decomp{
			{
				Pattern: "*",
				Responses: []string{
					"Are you saying 'no' just to be negative?",
					"You are being a bit negative.",
					"Why not?",
					"Why 'no'?",
				},
			},
		},
	},
	{
		Keyword:  "my",
		Priority: 2,
		Decomps: []Decomp{
			{
				Pattern: "* my *",
				Responses: []string{
					"Why are you concerned about your (2)?",
					"What about your own (2)?",
					"Does that have anything to do with the fact that your (2)?",
					"Lets discuss further why your (2).",
					"Earlier you said your (2).",
				},
			},
		},
	},
	{
		Keyword:  "mother",
		Priority: 100,
		Decomps: []Decomp{
			{Pattern: "*", Responses: []string{"=family"}},
		},
	},
	{
		Keyword:  "mom",
		Priority: 100,
		Decomps: []Decomp{
			{Pattern: "*", Responses: []string{"=family"}},
		},
	},
	{
		Keyword:  "father",
		Priority: 100,
		Decomps: []Decomp{
			{Pattern: "*", Responses: []string{"=family"}},
		},
	},
	{
		Keyword:  "dad",
		Priority: 100,
		Decomps: []Decomp{
			{Pattern: "*", Responses: []string{"=family"}},
		},
	},
	{
		Keyword:  "sister",
		Priority: 100,
		Decomps: []Decomp{
			{Pattern: "*", Responses: []string{"=family"}},
		},
	},
	{
		Keyword:  "brother",
		Priority: 100,
		Decomps: []Decomp{
			{Pattern: "*", Responses: []string{"=family"}},
		},
	},
	{
		Keyword:  "wife",
		Priority: 100,
		Decomps: []Decomp{
			{Pattern: "*", Responses: []string{"=family"}},
		},
	},
	{
		Keyword:  "husband",
		Priority: 100,
		Decomps: []Decomp{
			{Pattern: "*", Responses: []string{"=family"}},
		},
	},
	{
		Keyword:  "children",
		Priority: 100,
		Decomps: []Decomp{
			{Pattern: "*", Responses: []string{"=family"}},
		},
	},
	{
		Keyword:  "family",
		Priority: 100,
		Decomps: []Decomp{
			{
				Pattern: "*",
				Responses: []string{
					"Tell me more about your family.",
					"Who else in your family (2)?",
					"Your family.",
					"What else comes to mind when you think of your family?",
				},
			},
		},
	},
	{
		Keyword:  "alike",
		Priority: 10,
		Decomps: []Decomp{
			{
				Pattern: "*",
				Responses: []string{
					"In what way?",
					"What resemblance do you see?",
					"What does that similarity suggest to you?",
					"What other connections do you see?",
					"Could there really be some connection?",
					"How?",
				},
			},
		},
	},
	{
		Keyword:  "like",
		Priority: 10,
		Decomps: []Decomp{
			{
				Pattern: "* * like *",
				Responses: []string{"=alike"},
			},
		},
	},
	{
		Keyword:  "what",
		Priority: 0,
		Decomps: []Decomp{
			{
				Pattern: "*",
				Responses: []string{
					"Why do you ask?",
					"Does that question interest you?",
					"What is it you really want to know?",
					"Are such questions much on your mind?",
					"What answer would please you most?",
					"What do you think?",
					"What comes to your mind when you ask that?",
					"Have you asked such questions before?",
					"Have you asked anyone else?",
				},
			},
		},
	},
	{
		Keyword:  "because",
		Priority: 0,
		Decomps: []Decomp{
			{
				Pattern: "*",
				Responses: []string{
					"Is that the real reason?",
					"Don't any other reasons come to mind?",
					"Does that reason seem to explain anything else?",
					"What other reasons might there be?",
				},
			},
		},
	},
	{
		Keyword:  "why",
		Priority: 0,
		Decomps: []Decomp{
			{
				Pattern: "* why don't you *",
				Responses: []string{
					"Do you believe I don't (2)?",
					"Perhaps I will (2) in good time.",
					"Should you (2) yourself?",
					"You want me to (2)?",
					"=what",
				},
			},
			{
				Pattern: "* why can't i *",
				Responses: []string{
					"Do you think you should be able to (2)?",
					"Do you want to be able to (2)?",
					"Do you believe this will help you to (2)?",
					"Have you any idea why you can't (2)?",
					"=what",
				},
			},
			{
				Pattern: "*",
				Responses: []string{"=what"},
			},
		},
	},
	{
		Keyword:  "everyone",
		Priority: 2,
		Decomps: []Decomp{
			{
				Pattern: "* everyone *",
				Responses: []string{
					"Really, everyone?",
					"Surely not everyone.",
					"Can you think of anyone in particular?",
					"Who, for example?",
					"Are you thinking of a very special person?",
					"Who, may I ask?",
					"Someone special perhaps?",
					"You have a particular person in mind, don't you?",
					"Who do you think you're talking about?",
				},
			},
		},
	},
	{
		Keyword:  "everybody",
		Priority: 2,
		Decomps: []Decomp{
			{Pattern: "* everybody *", Responses: []string{"=everyone"}},
		},
	},
	{
		Keyword:  "nobody",
		Priority: 2,
		Decomps: []Decomp{
			{Pattern: "* nobody *", Responses: []string{"=everyone"}},
		},
	},
	{
		Keyword:  "noone",
		Priority: 2,
		Decomps: []Decomp{
			{Pattern: "* noone *", Responses: []string{"=everyone"}},
		},
	},
	{
		Keyword:  "always",
		Priority: 1,
		Decomps: []Decomp{
			{
				Pattern: "*",
				Responses: []string{
					"Can you think of a specific example?",
					"When?",
					"What incident are you thinking of?",
					"Really, always?",
				},
			},
		},
	},
	{
		Keyword:  "think",
		Priority: 1,
		Decomps: []Decomp{
			{
				Pattern: "*",
				Responses: []string{
					"Do you really think so?",
					"But you are not sure?",
					"Do you doubt it?",
				},
			},
		},
	},
	{
		Keyword:  "alike",
		Priority: 10,
		Decomps: []Decomp{
			{
				Pattern: "*",
				Responses: []string{
					"In what way?",
					"What resemblance do you see?",
					"What does that similarity suggest to you?",
					"What other connections do you see?",
					"Could there really be some connection?",
					"How?",
				},
			},
		},
	},
	{
		Keyword:  "how",
		Priority: 0,
		Decomps: []Decomp{
			{Pattern: "*", Responses: []string{"=what"}},
		},
	},
	{
		Keyword:  "who",
		Priority: 0,
		Decomps: []Decomp{
			{Pattern: "*", Responses: []string{"=what"}},
		},
	},
	{
		Keyword:  "when",
		Priority: 0,
		Decomps: []Decomp{
			{Pattern: "*", Responses: []string{"=what"}},
		},
	},
	{
		Keyword:  "where",
		Priority: 0,
		Decomps: []Decomp{
			{Pattern: "*", Responses: []string{"=what"}},
		},
	},
	{
		Keyword:  "can you",
		Priority: 0,
		Decomps: []Decomp{
			{
				Pattern: "* can you *",
				Responses: []string{
					"You believe I can (2) don't you?",
					"=what",
					"You want me to be able to (2)?",
					"Perhaps you would like to be able to (2) yourself.",
				},
			},
		},
	},
	{
		Keyword:  "can i",
		Priority: 0,
		Decomps: []Decomp{
			{
				Pattern: "* can i *",
				Responses: []string{
					"Whether or not you can (2) depends on you more than on me.",
					"Do you want to be able to (2)?",
					"Perhaps you don't want to (2).",
					"=what",
				},
			},
		},
	},
	{
		Keyword:  "quit",
		Priority: 999,
		Decomps: []Decomp{
			{
				Pattern: "*",
				Responses: []string{
					"Goodbye. It was nice talking to you.",
				},
			},
		},
	},
}
