// Package prompt provides a simple prompt for user input in CLI applications.
// while there are other fantastic packages that provide such functionality eg. as https://github.com/c-bata/go-prompt
// or https://github.com/AlecAivazis/survey most of them provide functionality one might not need while hacking together simple CLI appplications
// simple prompt aims to be quick to understand and use, or to just provide an example you can copy and paste from the parts necessary for your use-case.
package prompt

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

// AskOptions provides optional arguments for Ask.
// See defaults below when passing empty AskOptions to Ask.
type AskOptions struct {
	Reader               io.RuneReader                           // default: bufio.NewReader(os.Stdin)
	Answers              []rune                                  // default: []rune{'y', 'n'}
	InvalidAnswerMessage string                                  // default: "Invalid answer, please try again"
	FailHandlerFunc      func(string, *AskOptions) (rune, error) // default: prompt.Ask
}

func setDefaults(opts *AskOptions) {
	if opts.InvalidAnswerMessage == "" {
		opts.InvalidAnswerMessage = "Invalid answer, please try again"
	}

	if opts.Answers == nil {
		opts.Answers = []rune{'y', 'n'}
	}

	if opts.Reader == nil {
		opts.Reader = bufio.NewReader(os.Stdin)
	}

	if opts.FailHandlerFunc == nil {
		opts.FailHandlerFunc = Ask
	}
}

// Ask prompts the provided question to the user and waits for a single character input from a io.RuneReader passed in AskOptions.Reader (default: os.Stdin).
// The answer is validated against the contents of AskOptions.Answers, and either the provided input is returned, or the AskOptions.InvalidAnswerMessage is printed
// and AskOptions.FailHandlerFunc is called. If no FailHandlerFunc is provided, Ask is called again, prompting again for a valid input.
//
func Ask(question string, opts *AskOptions) (rune, error) {
	setDefaults(opts)

	fmt.Println(question)

	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()

	char, _, err := opts.Reader.ReadRune()

	if !isRuneContained(char, opts.Answers) {
		fmt.Println(opts.InvalidAnswerMessage)
		return opts.FailHandlerFunc(question, opts)
	}

	return char, err
}

func isRuneContained(r rune, runeSlice []rune) bool {
	for _, item := range runeSlice {
		if item == r {
			return true
		}
	}
	return false
}
