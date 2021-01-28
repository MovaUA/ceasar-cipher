package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"unicode"
)

func main() {
	shiftFlag := flag.Int("shift", 0, "number of character to shift input")
	file := flag.String("file", "", "input file (if empty - stdin)")
	flag.Parse()

	f := os.Stdin
	if *file != "" {
		f, err := os.Open(*file)
		if err != nil {
			log.Fatalf("could not open file: %v\n", err)
		}
		defer f.Close()
	}

	n := *shiftFlag % 26
	if n < 0 {
		n += 26
	}

	reader := bufio.NewReader(f)

	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalln(err)
		}
		if r == unicode.ReplacementChar {
			log.Println("invalid code point is found")
			continue
		}
		os.Stdout.WriteString(string(cipher(r, n)))
	}
}

func cipher(r rune, n int) rune {
	switch {
	case r >= 'a' && r <= 'z':
		return rotate(r, 'a', n)
	case r >= 'A' && r <= 'B':
		return rotate(r, 'A', n)
	default:
		return r
	}
}

func rotate(r, base rune, n int) rune {
	return base + (r-base+rune(n))%26
}
