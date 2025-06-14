package cmd

import (
	"fmt"

	"github.com/bmf-san/ggc/git"
)

func Push(args []string) {
	if len(args) > 0 {
		switch args[0] {
		case "current":
			err := git.PushCurrentBranch()
			if err != nil {
				fmt.Println("Error:", err)
			}
			return
		case "force":
			err := git.PushForceCurrentBranch()
			if err != nil {
				fmt.Println("Error:", err)
			}
			return
		}
	}
	ShowPushHelp()
}

func ShowPushHelp() {
	fmt.Println("Usage: ggc push current | ggc push force")
}
