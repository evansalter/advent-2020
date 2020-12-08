package part1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/evansalter/advent-2020/helpers"
)

type Op string

const (
	OpAcc Op = "acc"
	OpJmp    = "jmp"
	OpNop    = "nop"
)

func OpFromString(str string) Op {
	return Op(str)
}

type Instruction struct {
	Op
	Arg      int
	Executed bool
}

func InstructionFromLine(line string) *Instruction {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		panic("invalid line")
	}
	arg, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err.Error())
	}
	return &Instruction{
		Op:  OpFromString(parts[0]),
		Arg: arg,
	}
}

type Computer struct {
	Instructions []*Instruction
	Counter      int
	Accumulator  int
}

func InitializeComputer(lines []string) *Computer {
	instrs := make([]*Instruction, len(lines)-1)
	for i, l := range lines {
		if l == "" {
			continue
		}
		instrs[i] = InstructionFromLine(l)
	}

	return &Computer{
		Instructions: instrs,
	}
}

func (c *Computer) RunUntilLoop() (acc int) {
	for {
		instr := c.Instructions[c.Counter]
		if instr.Executed {
			return c.Accumulator
		}
		switch instr.Op {
		case OpJmp:
			c.Counter += instr.Arg
		case OpAcc:
			c.Accumulator += instr.Arg
			fallthrough
		default:
			c.Counter++
		}
		instr.Executed = true
	}
}

func Run() {
	lines := helpers.ReadInputFile(8)
	c := InitializeComputer(lines)
	acc := c.RunUntilLoop()
	fmt.Println(acc)
}
