package main

import (
	"os"

	"github.com/andytom/tmpltr/cmd"
)

func main() {

	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
