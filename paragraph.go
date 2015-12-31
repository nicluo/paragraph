package paragraph

//go-bindata -pkg paragraph data/

import (
	"bufio"
	"bytes"
	"log"
	"math/rand"
	"strings"
)

func countLines(data *bytes.Reader) int {
	fileScanner := bufio.NewScanner(data)
	defer data.Seek(0, 0)

	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}

	return lineCount
}

func readEnable1() []string {
	data, err := Asset("data/enable1.txt")

	if err != nil {
		log.Fatal(err)
	}

	reader := bytes.NewReader(data)
	lineCount := countLines(reader)
	lines := make([]string, lineCount)

	lineNumber := 0
	fileScanner := bufio.NewScanner(reader)
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
