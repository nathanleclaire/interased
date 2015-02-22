package main

import (
	"fmt"
	"log"
)

type InterasedUI interface {
	Bail() error
	PromptInput() (string, error)
	Render() error
}

type InterasedTextUI struct {
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
	fmt.Print("Pattern: ")
	return nil
}

func main() {
	ui := InterasedTextUI{}
	for {
		if err := ui.Render(); err != nil {
			if err := ui.Bail(err); err != nil {
				log.Panicln(err)
			}
			return
		}
		in, err := ui.PromptInput()
		if err != nil {
			if err := ui.Bail(err); err != nil {
				log.Panicln(err)
			}
			return
		}
		log.Fatal(in)
	}
}
