package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	shiftFlag := flag.Int("shift", 0, "number of character to shift input")
	file := flag.String("file", "", "input file (if empty - stdin)")
	flag.Parse()

	r := os.Stdin
	if *file != "" {
		f, err := os.Open(*file)
		if err != nil {
			log.Fatalf("could not open file: %v\n", err)
		}
		defer f.Close()
		r = f
	}

	n := *shiftFlag % 26

	s := bufio.NewScanner(r)
	for s.Scan() {
		in := s.Bytes()
		out := make([]byte, len(in))
		for i := 0; i < len(in); i++ {
			out[i] = cipher(in[i], n)
		}
		fmt.Println(string(out))
	}

	if err := s.Err(); err != nil {
		log.Fatalf("could not read: %v\n", err)
	}
}

func cipher(b byte, n int) byte {
	if new, ok := shift(b, byte('a'), byte('z'), n); ok {
		return new
	}
	new, _ := shift(b, byte('A'), byte('Z'), n)
	return new
}

func shift(b, min, max byte, n int) (byte, bool) {
	if b < min || b > max {
		return b, false
	}
	b = byte(int(b) + n)
	if b > max {
		b -= 26
	}
	if b < min {
		b += 26
	}
	return b, true
}
