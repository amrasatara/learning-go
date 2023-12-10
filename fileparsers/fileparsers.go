package fileparsers

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// ReadLines method return slice of string will all lines in the file
func ReadLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		log.Fatalf("readLines: %s", scanner.Err())
	}

	return lines
}

func Write(sb strings.Builder, filename string) {
	file, _ := os.Create(filename)
	defer file.Close()

	file.WriteString(sb.String())
}

func ParseToStringObj(input string, separator string) []string {
	return strings.Split(input, separator)
}
