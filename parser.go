package ctgov

import (
	"bufio"
	"io"

	"github.com/gomarkdown/markdown"
)

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

func calcIndentLevel(line []byte) int {
	counter := 0
	for i := 0; i < len(line); i++ {
		if line[i] == ' ' {
			counter = counter + 1
		} else {
			break
		}
	}
	return counter
}

func NormalizeIndent(r io.Reader) []byte {
	scanner := bufio.NewScanner(r)
	var (
		baseIndent int
		first      bool
		retBuffer  []byte
	)
	first = true
	for scanner.Scan() {
		line := scanner.Bytes()
		if first && len(line) > 0 {
			baseIndent = calcIndentLevel(line)
			first = false
		} else if first {
			continue
		}
		if len(line) > 0 {
			line = line[baseIndent:]
			retBuffer = append(retBuffer, string('\n')...)
			retBuffer = append([]byte(retBuffer), []byte(line)...)
		} else {
			continue
		}
	}
	return retBuffer
}

// parse ctgov raw textBlock into better looking html
func (pa *Parser) Parse(r io.Reader) []byte {
	buffer := NormalizeIndent(r)
	return []byte(markdown.ToHTML(buffer, nil, nil))
}
