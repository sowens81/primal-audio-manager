package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Prompt(message string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(message)

	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func PromptRequired(message string) string {
	for {
		input := Prompt(message)
		if input != "" {
			return input
		}
		fmt.Println("Value is required, please try again.")
	}
}
