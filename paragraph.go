package paragraph

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strings"
)

func countLines(file *os.File) int {
	fileScanner := bufio.NewScanner(file)
	defer file.Seek(0, 0)

	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}

	return lineCount
}

func readEnable1() []string {
	file, err := os.Open("enable1.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lineCount := countLines(file)
	lines := make([]string, lineCount)

	lineNumber := 0
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		lines[lineNumber] = fileScanner.Text()
		lineNumber++
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func GenerateSentence(sentenceLength int) string {
	words := readEnable1()
	sentence := make([]string, sentenceLength)

	for i := 0; i < 10; i++ {
		index := rand.Intn(len(words))
		sentence[i] = words[index]
	}

	return strings.Join(sentence, " ")
}
