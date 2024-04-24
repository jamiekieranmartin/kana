package kana

import "fmt"

type GameOptsFunc func(*GameOpts)

type GameMode string

const (
	HiraganaMode GameMode = "hiragana"
	KatakanaMode GameMode = "katakana"
)

type GameOpts struct {
	guesses int
	rounds  int
	mode    GameMode
}

type Game struct {
	GameOpts
}

func defaultOpts() GameOpts {
	return GameOpts{
		guesses: 3,
		rounds:  10,
		mode:    HiraganaMode,
	}
}

func NewGame(optFuncs ...GameOptsFunc) *Game {
	opts := defaultOpts()

	for _, fn := range optFuncs {
		fn(&opts)
	}

	return &Game{opts}
}

func WithGuesses(guesses int) GameOptsFunc {
	return func(opts *GameOpts) {
		opts.guesses = guesses
	}
}

func WithRounds(rounds int) GameOptsFunc {
	return func(opts *GameOpts) {
		opts.rounds = rounds
	}
}

func WithMode(mode GameMode) GameOptsFunc {
	return func(opts *GameOpts) {
		opts.mode = mode
	}
}

func (g *Game) Run() {
	kana := Hiragana

	switch g.mode {
	case "hiragana":
		kana = Hiragana
	case "katakana":
		kana = Katakana
	}

	for i := 0; i < g.rounds; i++ {
		kana := RandomKana(&kana)

		fmt.Printf("What is %s in romaji?\n", kana.Kana)

		for i := 0; i < g.guesses; i++ {
			var guess string
			fmt.Scanln(&guess)

			if guess == kana.Romaji {
				fmt.Println("Correct!")
				break
			} else {
				fmt.Println("Incorrect, try again.")
			}
		}

		fmt.Printf("The answer was %s.\n", kana.Romaji)
	}
}
