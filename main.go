package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/RisingStack/simple-prompt/prompt"
	"github.com/atotto/clipboard"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the string to be cut: ")
	text, _ := reader.ReadString('\n')

	text = strings.TrimSuffix(text, "\n")

	slicesOfFour := createSlices(text, 4)

	getUserInputForChunks(slicesOfFour)
}

func createSlices(text string, n int) []string {
	inputRunes := []rune(text)

	slicesOfFour := []string{}
	substr := ""
	for i, r := range inputRunes {
		substr = substr + string(r)
		if (i+1)%n == 0 {
			slicesOfFour = append(slicesOfFour, substr)
			substr = ""
		} else if (i + 1) == len(inputRunes) {
			slicesOfFour = append(slicesOfFour, substr)
		}
	}

	return slicesOfFour
}

func getUserInputForChunks(slicesOfFour []string) {
	for _, chunk := range slicesOfFour {
		fmt.Println("Next chunk: " + chunk)
		char, err := prompt.Ask("Would you like to (c)opy (s)kip or (e)xit [c/s/e]", &prompt.AskOptions{Answers: []rune{'c', 's', 'e'}})

		if err != nil {
			panic(err)
		}

		switch char {
		case 'c':
			clipboard.WriteAll(chunk)
		case 's':
			continue
		case 'e':
			os.Exit(0)
		}
	}
}
