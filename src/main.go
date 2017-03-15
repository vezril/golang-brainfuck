package main

import (
	"fmt"
	"io/ioutil"
	"github.com/vezril/golang-cookiejar/collections/tape"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Interpreter struct {
	// Instructions
	i *tape.Tape
	// Data
	d *tape.Tape
}

func New() *Interpreter {
	interpreter := new(Interpreter)
	interpreter.i = tape.New()
	interpreter.d = tape.New()
	return interpreter
}

func main() {

		dat, err := ioutil.ReadFile("test/hello.bf")
		check(err)
		interpreter := New()

		// Load instruction data into memory
		for _,b := range dat {
			interpreter.i.Write(b)
			interpreter.i.Right()
		}
		interpreter.i.Sync(0)
		signal := true

		for signal {
				token := interpreter.i.Read()
				if token == 62 { // >
					interpreter.d.Right()
				} else if token == 60 { // <
					interpreter.d.Left()
				} else if token == 43 { // +
					interpreter.d.Inc()
				} else if token == 45 { // -
					interpreter.d.Dec()
				} else if token == 46 { // .
					fmt.Print(string(interpreter.d.Read()))
				} else if token == 91 { // [
					if interpreter.d.Read() == 0 {
						fmt.Println("SyncToClosing")
						interpreter.i.SyncToClosing()
					}
				} else if token == 93 { // ]
					if interpreter.d.Read() != 0 {
						interpreter.i.SyncToOpening()
					}
				} else if token == 0 {
					signal = false
				}
				// Increment instruction pointer
				interpreter.i.Right()
		}
}
