package main

import (
	"flag"
	"fmt"

	"github.com/jamiekieranmartin/kana"
)

const cliVersion = "0.0.1"

const helpMessage = `
kana v%s

`

func main() {
	flag.Usage = func() {
		fmt.Printf(helpMessage, cliVersion)
		flag.PrintDefaults()
	}

	// cli arguments
	version := flag.Bool("version", false, "Print version string and exit")
	help := flag.Bool("help", false, "Print help message and exit")
	guesses := flag.Int("guesses", 3, "Number of guesses per character")
	rounds := flag.Int("rounds", 10, "Number of rounds to play")
	mode := flag.String("mode", "hiragana", "Mode to play in (hiragana, katakana)")

	flag.Parse()

	// if asked for version, disregard everything else
	if *version {
		fmt.Printf("kana v%s\n", cliVersion)
		return
	} else if *help {
		flag.Usage()
		return
	}

	kana.NewGame(
		kana.WithGuesses(*guesses),
		kana.WithRounds(*rounds),
		kana.WithMode(kana.GameMode(*mode)),
	).Run()
}
