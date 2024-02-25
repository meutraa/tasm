package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	R   = 1
	IMM = 2
	L   = 4
)

func output(a, b, c, d uint16) string {
	return fmt.Sprintf("%#02x %#02x %#02x %#02x", a, b, c, d)
}

func threeL(code uint16) Instruction {
	return Instruction{
		opcode:  code,
		params:  []uint16{L | IMM, R | IMM, R | IMM},
		mapping: []uint16{3, 1, 2},
		out: func(code uint16, p []uint16) string {
			return output(code, p[1], p[2], p[0])
		},
	}
}

func three(code uint16) Instruction {
	return Instruction{
		opcode:  code,
		params:  []uint16{R, R | IMM, R | IMM},
		mapping: []uint16{3, 1, 2},
		out: func(code uint16, p []uint16) string {
			return output(code, p[1], p[2], p[0])
		},
	}
}

func two(code uint16) Instruction {
	return Instruction{
		opcode:  code,
		params:  []uint16{R, R | IMM},
		mapping: []uint16{3, 1},
		out: func(code uint16, p []uint16) string {
			return output(code, p[1], unused, p[0])
		},
	}
}

func one(code uint16) Instruction {
	return Instruction{
		opcode:  code,
		params:  []uint16{R | IMM},
		mapping: []uint16{1},
		out: func(code uint16, p []uint16) string {
			return output(code, p[0], unused, unused)
		},
	}
}

func oneRd(code uint16) Instruction {
	return Instruction{
		opcode:  code,
		params:  []uint16{R},
		mapping: []uint16{3},
		out: func(code uint16, p []uint16) string {
			return output(code, unused, unused, p[0])
		},
	}
}

var instructions = map[string]Instruction{
	"add":    three(0),
	"sub":    three(1),
	"and":    three(2),
	"or":     three(3),
	"not":    two(4),
	"xor":    three(5),
	"push":   one(6),
	"pop":    oneRd(7),
	"mov":    two(8),
	"mull":   three(9),
	"mulu":   three(10),
	"shl":    three(11),
	"shr":    three(12),
	"jmpe":   threeL(32),
	"jmpne":  threeL(33),
	"jmplt":  threeL(34),
	"jmplte": threeL(35),
	"jmpgt":  threeL(36),
	"jmpgte": threeL(37),
	"call":   Instruction{params: []uint16{L}},
	"ret":    Instruction{params: []uint16{}},
	"jmp":    Instruction{params: []uint16{L}},
	"dec":    Instruction{params: []uint16{R}},
	"inc":    Instruction{params: []uint16{R}},
}

const (
	imm1   = 0b10000000
	imm2   = 0b01000000
	unused = 0b11111111
)

type Instruction struct {
	opcode  uint16
	params  []uint16
	mapping []uint16
	out     func(uint16, []uint16) string
}

var registers = map[string]uint16{
	"r0":  0,
	"r1":  1,
	"r2":  2,
	"r3":  3,
	"r4":  4,
	"ra":  5, // ram address
	"pa":  6, // program address
	"in":  7, // general purpose input
	"out": 7, // general purpose output
	"r5":  8,
	"r6":  9,
	"r7":  10,
	"r8":  11,
	"r9":  12,
	"r10": 13,
	"r11": 14,
	"r12": 15,
	"rm":  16, // ram
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./tasm <filename>")
		return
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	labels := map[string]uint16{}

	var address uint16 = 0
	scanner := bufio.NewScanner(file)

	// first scan for addresses
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) == 0 || strings.HasPrefix(fields[0], ";") {
			continue
		}

		instruction := fields[0]
		isLabel := strings.HasSuffix(instruction, ":")
		if !isLabel {
			if instruction == "call" {
				address += 8
			} else {
				address += 4
			}
		}
		if isLabel {
			labels[strings.Split(instruction, ":")[0]] = address
			//log.Println(strings.Split(instruction, ":")[0], address)
		}
	}

	address = 0
	file.Seek(0, io.SeekStart)
	scanner = bufio.NewScanner(file)
	lineNo := -1
	for scanner.Scan() {
		lineNo++
		lineog := scanner.Text()
		line := strings.ReplaceAll(lineog, ",", "")
		line = strings.Split(line, ";")[0]
		fields := strings.Fields(line)

		if len(fields) == 0 {
			continue
		}

		instruction := fields[0]
		// Check if this is a label line
		if strings.HasSuffix(instruction, ":") {
			fmt.Println("# " + lineog)
			continue
		}

		inst, ok := instructions[instruction]
		if !ok {
			log.Fatalf("%v: instruction undefined: %v\n", lineNo, instruction)
		}

		paramCount := len(inst.params)
		if paramCount != len(fields)-1 {
			log.Fatalf("%v: instruction parameter count mismatch: expected %v, got %v\n", lineNo, paramCount, len(fields)-1)
		}

		p := make([]uint16, paramCount)
		for i := 0; i < paramCount; i++ {
			// 0 - 2 for three params
			parameter := inst.params[i]
			a := fields[i+1]
			// Is this a named destination
			dest, ok := registers[a]
			if ok && (parameter&R == R) {
				p[i] = dest
				continue
			}
			// Is this a label
			dest, ok = labels[a]
			if ok && (parameter&L == L) {
				p[i] = dest
				continue
			}
			// It must be an immediate value
			if parameter&IMM != IMM {
				log.Fatalf("%v: instruction parameter type not allowed: %v\n", lineNo, a)
			}
			value, err := strconv.ParseUint(a, 0, 8)
			if nil != err {
				log.Fatalf("%v: unable to parse parameter: %v\n", lineNo, err)
			}
			switch inst.mapping[i] {
			case 1:
				inst.opcode |= imm1
			case 2:
				inst.opcode |= imm2
			default:
				// this is likely a jump to an address
			}
			p[i] = uint16(value)
		}

		fmt.Printf("#(%v) %v\n", address, lineog)
		fmt.Println(inst.out(inst.opcode, p))
		/*switch instruction {
		case "jmp":
			fmt.Printf(format, opcodes["mov"]|imm1, dest, null, places["pa"])
		case "dec":
			fmt.Printf(format, opcodes["sub"]|imm2, dest, 1, dest)
		case "inc":
			fmt.Printf(format, opcodes["add"]|imm2, dest, 1, dest)
		case "call":
			//fmt.Printf("#(%v) (push | imm1) %v null null\n", address, address+8)
			fmt.Printf(format, opcodes["push"]|imm1, address+8, null, null)
			address += 4
			//fmt.Printf("#(%v) mov %v pa\n", address, dest)
			fmt.Printf(format, opcodes["mov"]|imm1, dest, null, places["pa"])
		case "ret":
			fmt.Printf(format, opcodes["pop"], null, null, places["pa"])
		}*/

		address += 4
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
