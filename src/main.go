package main

import (
	"fmt"
	"io/ioutil"
	"github.com/vezril/go-cookiejar/collections/tape"
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

		dat, err := ioutil.ReadFile("/home/cference/.go/src/github.com/vezril/golang-brainfuck/src/hello.bf")
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
					fmt.Println(string(interpreter.d.Read()))
				} else if token == 91 { // [
					if interpreter.d.Read() == 0 {
						interpreter.i.SyncToClosing()
					}
				} else if token == 93 { // ]
					if interpreter.d.Read() != 0 {
						interpreter.i.SyncToOpening()
						// This is so the catch all instruction pointer gets set at the correct position.
						// I should probably refactor this to something less ugly. switch-case maybe?
						interpreter.i.Left()
					}
				} else if token == 0 {
					signal = false
				}
				// Increment instruction pointer
				interpreter.i.Right()
		}
		//

}
