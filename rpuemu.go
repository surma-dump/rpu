package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type RPU struct {
	Memory []byte
	IP uint16
	Accumulator byte
}

func readFile(filepath string) ([]byte, os.Error) {
        f, e := os.Open(filepath, os.O_RDONLY, 0)
        if e != nil {
                return "", e
        }
        defer f.Close()

        bqry, e := ioutil.ReadAll(f)
        if e != nil {
                return "", e
        }
        return bqry, e
}

func parse() ([]byte) {
	file := flag.String("f", "", "Imagefile to load")
	help := flag.Bool("h", false, "Show help")
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	mem, e := readFile(file)
	if e != nil {
		fmt.Fprintf(os.Stderr, "Could not read file: %s\n", e.String())
		os.Exit(1)
	}

	return mem
}

func New(memory []byte) (r *RPU) {
	r = new(RPU)
	r.Memory = make([]byte, 4096)
	r.Memory[0:len(memory)] = memory[0:]
	return
}

func (r *RPU) getOp() byte {
	return (r.Memory[r.ip] >> 7) & 1
}

func (r *RPU) getOperand() uint16  {
	return ((r.Memory[r.ip] & 0x7F) << 8) | r.Memory[r.ip+1]
}

func main() {
	memory := parse()

	i := New(memory)
	
}
