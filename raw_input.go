package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/nsf/termbox-go"
)

const (
	backspaceKey = 127
	coldef       = termbox.ColorDefault
)

var (
	current    string
	stdinbytes []byte
)

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func redraw_all(stdinbytes []byte) {
	termbox.Clear(coldef, coldef)
	tbprint(0, 0, termbox.ColorMagenta, coldef, "Press enter to quit")
	tbprint(0, 1, coldef, coldef, fmt.Sprintf("%q", current))
	sedcmd := exec.Command("sed", fmt.Sprintf("%q", current))
	pipe, err := sedcmd.StdinPipe()
	if err != nil {
		log.Panicln("Error creating sed pipe")
	}
	if _, err := io.Write(pipe, stdinbytes); err != nil {
		log.Panicln("Error writing string to sed pipe")
	}
	sedout, err := sedcmd.Output()
	if err != nil {
		log.Panicln("Error getting sed output")
	}
	tbprint(0, 2, coldef, coldef, "sed output:")
	tbprint(0, 3, coldef, coldef, string(sedout))
	termbox.Flush()
}

func main() {
	stdinbytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Panicln("Error reading from stdin")
	}
	err = termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputAlt)
	redraw_all(stdinbytes)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				current = current[:len(current)-1]
			case termbox.KeySpace:
				current += " "
			case termbox.KeyEnter:
				termbox.Clear(coldef, coldef)
				termbox.Flush()
				log.Fatal("BOOM!")
			default:
				current += string(ev.Ch)
			}
		case termbox.EventError:
			panic(ev.Err)
		}

		redraw_all(stdinbytes)
	}
}
