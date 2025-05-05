package main

import (
	"crypto/fips140"
	"crypto/rand"
	"flag"
	"fmt"
	"os"
)

func main() {
	size := flag.String("size", "1G", "size of random data (supports M, G suffixes), minimum 1M")
	output := flag.String("output", "./rng.bin", "output file path")
	flag.Parse()

	var totalSize int64
	switch s := *size; s[len(s)-1] {
	case 'M':
		totalSize = int64(1024 * 1024 * atoi(s[:len(s)-1]))
	case 'G':
		totalSize = int64(1024 * 1024 * 1024 * atoi(s[:len(s)-1]))
	default:
		totalSize = int64(atoi(s))
	}

	if !fips140.Enabled() {
		fmt.Fprintln(os.Stderr, "Warning: FIPS 140-3 mode is NOT enabled")
	}

	const chunkSize = 1024 * 1024

	file, err := os.Create(*output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create file: %v\n", err)
		return
	}
	defer file.Close()

	buf := make([]byte, chunkSize)
	var written int64

	for written < totalSize {
		n, err := rand.Read(buf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading random data: %v\n", err)
			return
		}

		if n != chunkSize {
			fmt.Fprintf(os.Stderr, "Unexpected chunk size: %d\n", n)
			return
		}

		nw, err := file.Write(buf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to file: %v\n", err)
			return
		}

		written += int64(nw)
		fmt.Printf("\rWritten: %d / %d bytes", written, totalSize)
	}

	fmt.Println("\nDone.")
}

func atoi(s string) int {
	n := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			fmt.Fprintf(os.Stderr, "Invalid number: %s\n", s)
			os.Exit(1)
		}
		n = n*10 + int(c-'0')
	}
	return n
}
