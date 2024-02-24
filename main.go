package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var opcodes = map[string]uint8{
	"imm1":   0b10000000,
	"imm2":   0b01000000,
	"null":   0b11111111, // read zero
	"mov":    0b00000000,
	"add":    0b00000000,
	"sub":    0b00000001,
	"and":    0b00000010,
	"or":     0b00000011,
	"not":    0b00000100,
	"xor":    0b00000101,
	"push":   0b00000110,
	"pop":    0b00000111,
	"jmpe":   32,
	"jmpne":  33,
	"jmplt":  34,
	"jmplte": 35,
	"jmpgt":  36,
	"jmpgte": 37,
}

var places = map[string]uint8{
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

	labels := map[string]uint8{}

	var address uint8 = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineog := scanner.Text()
		line := strings.ReplaceAll(lineog, ",", "")
		fields := strings.Fields(line)

		if len(fields) == 0 {
			continue
		}

		// If this is a label or comment, ignore it
		if strings.HasPrefix(line, "#") {
			fmt.Println(lineog)
			continue
		}

		instruction := fields[0]
		if instruction == "label" {
			fmt.Println("# " + lineog)
			labels[fields[1]] = address
			continue
		}

		address += 4

		inst, ok := opcodes[instruction]
		if !ok {
			log.Fatalln("instruction undefined", instruction)
		}

		var dest, s1, s2 uint8
		if len(fields) > 1 {
			a := fields[1]
			if strings.HasPrefix(instruction, "jmp") {
				dest, ok = labels[a]
				if !ok {
					log.Fatalln("label not defined", a)
				}
			} else {
				var ok bool
				dest, ok = places[a]
				if !ok {
					log.Fatalln("destination not defined", a)
				}
			}
		}
		if len(fields) > 2 {
			a := fields[2]
			aint64, err := strconv.ParseUint(a, 10, 8)
			aisimm := nil == err
			if aisimm {
				inst |= opcodes["imm1"]
				s1 = uint8(aint64)
			} else {
				s1, ok = places[a]
				if !ok {
					log.Fatalln("first place not defined", a)
				}
			}
		}
		if len(fields) > 3 {
			a := fields[3]
			aint64, err := strconv.ParseUint(a, 10, 8)
			aisimm := nil == err

			if aisimm {
				inst |= opcodes["imm2"]
				s2 = uint8(aint64)
			} else {
				s2, ok = places[a]
				if !ok {
					log.Fatalln("second place not defined", a)
				}
			}
		}

		format := "%#08b %#08b %#08b %#08b\n"
		null := opcodes["null"]
		fmt.Printf("# %v\n", lineog)
		switch instruction {
		case "push":
			fmt.Printf(format, inst, s1, null, null)
		case "pop":
			fmt.Printf(format, inst, null, null, dest)
		case "mov", "not":
			fmt.Printf(format, inst, s1, null, dest)
		case "add", "sub", "and", "or", "xor", "jmpe", "jmpne", "jmplt", "jmplte", "jmpgt", "jmpgte":
			fmt.Printf(format, inst, s1, s2, dest)
		default:
			fmt.Println("Unknown command:", instruction)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}