package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SureToProceed(message string, params ...any) bool {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf(message, params...)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Error reading input: %s\n", err)
		return false
	}

	response := strings.ToLower(strings.TrimSpace(input))

	return response == "y"
}
