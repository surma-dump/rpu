package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type RPU struct {
	memory []byte
	ip uint16
	accumulator Bit
}

type Bit byte
type Op Bit

const (
    LOAD = 0
    STORE = 1
)

type Phase byte
const (
    LOAD1 = 0
    LOAD2 = 1
    STORE = 2
)
func readFile(filepath string) ([]byte, os.Error) {
        f, e := os.Open(filepath, os.O_RDONLY, 0)
        if e != nil {
                return nil, e
        }
        defer f.Close()

        bqry, e := ioutil.ReadAll(f)
        if e != nil {
                return nil, e
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

	mem, e := readFile(*file)
	if e != nil {
		fmt.Fprintf(os.Stderr, "Could not read file: %s\n", e.String())
		os.Exit(1)
	}

	return mem
}

func New(memory []byte) (r *RPU) {
	r = new(RPU)
	r.memory = make([]byte, 4096)
    for i,v := range memory {
        if i >= len(r.memory) {
            break
        }
        r.memory[i] = v
    }
    r.ip = 0;
	return
}

func (r *RPU) GetOp() Op {
	return (r.memory[r.ip] >> 7) & 1
}

func (r *RPU) GetPhase() Phase {
    return r.IP % 3
}

func (r *RPU) GetOperand() uint16  {
	return ((uint16(r.memory[r.ip]) & 0x7F) << 8) | uint16(r.memory[r.ip+1])
}

func (r *RPU) GetAddress(addr uint16) Bit {
    return (m.memory[addr/8] >> (addr%8)) & 1
}

func (r *RPU) Next() {
    r.ip++
}

func (r *RPU) Exec() {
    switch r.getOp() {
        case LOAD:
            
        case STORE:
    }
}

func (r *RPU) GetIP() uint16 {
    return r.ip
}

func main() {
	memory := parse()

	m := New(memory)
    for i:=0; i < 2048; i++ {
        fmt.Printf("%02x: %x %04x\n", m.getIP(), m.getOp(), m.getOperand())
        m.next()
    }
}
