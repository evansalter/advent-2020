package part2

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/evansalter/advent-2020/helpers"
)

var executionError = errors.New("execution error")

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

func (c *Computer) ResetRegisters() {
	c.Counter = 0
	c.Accumulator = 0
	for _, instr := range c.Instructions {
		instr.Executed = false
	}
}

func (c *Computer) Run() (acc int, err error) {
	for {
		if c.Counter == len(c.Instructions) {
			return c.Accumulator, nil
		}
		if c.Counter < 0 {
			return 0, executionError
		}
		instr := c.Instructions[c.Counter]
		if instr.Executed {
			return 0, executionError
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

	for _, instr := range c.Instructions {
		original := instr.Op
		if instr.Op == OpJmp {
			instr.Op = OpNop
		} else if instr.Op == OpNop {
			instr.Op = OpJmp
		} else {
			continue
		}
		acc, err := c.Run()
		if err == executionError {
			instr.Op = original
			c.ResetRegisters()
			continue
		} else {
			fmt.Println(acc)
			break
		}
	}
}
