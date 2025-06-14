package cmd

import (
	"fmt"

	"github.com/bmf-san/ggc/git"
)

func Commit(args []string) {
	if len(args) > 0 {
		switch args[0] {
		case "allow-empty":
			err := git.CommitAllowEmpty()
			if err != nil {
				fmt.Println("Error:", err)
			}
			return
		case "tmp":
			err := git.CommitTmp()
			if err != nil {
				fmt.Println("Error:", err)
			}
			return
		}
	}
	ShowCommitHelp()
}

func ShowCommitHelp() {
	fmt.Println("Usage: ggc commit allow-empty | ggc commit tmp")
}
