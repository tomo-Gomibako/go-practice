package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// String string input test
func String() {
	stdin := bufio.NewScanner(os.Stdin)
	fmt.Println("Type your name.")
	stdin.Scan()
	text := stdin.Text()
	if len(text) == 0 {
		text = "anonymous"
	}
	text = strings.ToUpper(string(text[0])) + text[1:]
	fmt.Println("Hello, " + text + ".")
}
