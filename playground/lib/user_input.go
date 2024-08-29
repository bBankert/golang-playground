package lib

import (
	"fmt"
	"strings"
)

// Re-usable function to display text and store input from stdinput
func GetUserInput(messageText string, variable any) {
	//better readability, ensure there is a space after input message
	if !strings.HasSuffix(messageText, " ") {
		messageText = messageText + " "
	}

	fmt.Print(messageText)

	fmt.Scan(variable)
}
