package selection

import (
	"fmt"
	"strconv"

	"github.com/mitchellh/cli"
)

// Prompt takes a list of choices and asks the user to pick one
func Prompt(ui cli.Ui, prompt string, choices []string) int {
	if len(choices) <= 1 {
		panic("not enough choices provided")
	}

	ui.Output(prompt)

	for i, c := range choices {
		ui.Output(fmt.Sprintf("[%d] %s", i+1, c))
	}

	for {
		resp, err := ui.Ask(fmt.Sprintf("[1-%d]:", len(choices)))
		if err != nil {
			ui.Error(err.Error())
			continue
		}

		n, err := strconv.Atoi(resp)
		if err != nil {
			ui.Error(err.Error())
			continue
		}
		if n < 1 || n > len(choices) {
			ui.Error(fmt.Sprintf("%d is not an option", n))
			continue
		}

		return n - 1
	}

	panic("unreachable")
}
