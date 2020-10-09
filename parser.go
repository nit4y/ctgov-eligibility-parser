package ctgov

import (
	"bufio"
	"bytes"
	"io"
)

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

func WriteOpenTag(el string, buffer *bytes.Buffer) {
	buffer.WriteString("<")
	buffer.WriteString(el)
	buffer.WriteString(">")
}
func WriteCloseTag(el string, buffer *bytes.Buffer) {
	buffer.WriteString("</")
	buffer.WriteString(el)
	buffer.WriteString(">")
}

// parse ctgov raw textBlock into better looking html
func (pa *Parser) Parse(r io.Reader) []byte {
	scanner := bufio.NewScanner(r)
	var (
		//newItem    bool
		buffer    bytes.Buffer
		nodeStack []*node
	)

	//newItem = true
	nodeStack = append(nodeStack, newNode(0))
	for scanner.Scan() {
		line := scanner.Bytes()
		n := newNode(calcLevel(line))
		n.lineType = calcLineType(line)
		n.htmlType = calcHTMLType(n.lineType)

		switch n.lineType {

		case emptyLine:
			continue

		case textLine, commentLine:
			continue

		case numberLine, dashLine:
			if len(nodeStack) > 0 {
				if n.level > nodeStack[len(nodeStack)-1].level {
					WriteOpenTag(htmlTypes[n.htmlType], &buffer)
					WriteOpenTag(htmlTypes[li], &buffer)
					buffer.WriteString(string(line))
					//WriteCloseTag(htmlTypes[ol], &buffer)

				} else if n.level == nodeStack[len(nodeStack)-1].level {
					n.htmlType = li
					WriteOpenTag(htmlTypes[n.htmlType], &buffer)
					buffer.WriteString(string(line))
					WriteCloseTag(htmlTypes[n.htmlType], &buffer)
					//

				} else {
					for len(nodeStack) > 0 {
						if n.level < nodeStack[len(nodeStack)-1].level {
							WriteCloseTag(htmlTypes[nodeStack[len(nodeStack)-1].htmlType], &buffer)
							nodeStack = nodeStack[:len(nodeStack)-1]
						} else {
							WriteOpenTag(htmlTypes[li], &buffer)
							buffer.WriteString(string(line))
							break
						}
					}
				}
				if n.lineType != nodeStack[len(nodeStack)-1].lineType {
					nodeStack = append(nodeStack, n)
				}

			}

		}

	}
	return buffer.Bytes()
}
