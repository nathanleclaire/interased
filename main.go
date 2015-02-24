package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

type InterasedUI interface {
	Clear() error
	Bail() error
	PromptInput() (string, error)
	Render() error
}

type InterasedTextUI struct {
	OriginalInput    string
	TransformedInput string
}

func (ui *InterasedTextUI) Clear() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func (ui *InterasedTextUI) Bail(err error) error {
	log.Panicln(err)
	return nil
}

func (ui *InterasedTextUI) PromptInput() (string, error) {
	var (
		pattern string
	)
	_, err := fmt.Scanf("%s", &pattern)
	return pattern, err
}

func (ui *InterasedTextUI) Render() error {
	fmt.Print(ui.OriginalInput)
	fmt.Print("Pattern: ")
	fmt.Println("Transformed output:")
	fmt.Print(ui.TransformedInput)
	return nil
}

func bail(ui InterasedTextUI, err error) {
	if err := ui.Bail(err); err != nil {
		log.Panicln(err)
	}
}

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	userIn := make(chan string)
	ui := InterasedTextUI{
		OriginalInput:    string(stdinBytes),
		TransformedInput: string(stdinBytes),
	}
	go func(userIn chan<- string) {
		for {
			pattern, err := ui.PromptInput()
			if err != nil {
				bail(ui, err)
				return
			}
			userIn <- pattern
			log.Fatal(in)
		}
	}()
	for {
		if err := ui.Clear(); err != nil {
			bail(ui, err)
			return
		}
		if err := ui.Render(); err != nil {
			bail(ui, err)
			return
		}
		<-userIn
	}
}
