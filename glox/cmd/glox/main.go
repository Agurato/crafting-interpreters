package glox

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Usage: glox [script]")
		os.Exit(64)
	}

	if len(os.Args) == 1 {
		err := runFile(os.Args[0])
		if err != nil {
			panic(err)
		}
	} else {
		err := runPrompt()
		if err != nil {
			panic(err)
		}
	}
}

func runFile(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = run(string(bytes))
	return err
}

func runPrompt() error {
	return nil
}

func run(source string) error {
	return nil
}
