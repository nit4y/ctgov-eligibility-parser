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
		lastNode  *node
	)

	//newItem = true
	nodeStack = append(nodeStack, newNode(0))
	for scanner.Scan() {
		line := scanner.Bytes()
		n := newNode(calcLevel(line))
		n.lineType, n.textStart = calcLineProps(line)
		n.htmlType = calcHTMLType(n.lineType)

		switch n.lineType {

		case emptyLine:
			if lastNode != nil {
				WriteCloseTag(htmlTypes[li], &buffer)
			}

		case textLine, commentLine:
			for len(nodeStack) > 0 {
				if n.level < nodeStack[len(nodeStack)-1].level {
					WriteCloseTag(htmlTypes[nodeStack[len(nodeStack)-1].htmlType], &buffer)
					nodeStack = nodeStack[:len(nodeStack)-1]
				} else {
					buffer.WriteString(" ")
					buffer.Write(line[n.textStart:])
					break
				}
			}

		case numberLine, dashLine:
			if len(nodeStack) > 0 {
				if n.level > nodeStack[len(nodeStack)-1].level {
					WriteOpenTag(htmlTypes[n.htmlType], &buffer)
					WriteOpenTag(htmlTypes[li], &buffer)
					buffer.Write(line[n.textStart:])

				} else if n.level == nodeStack[len(nodeStack)-1].level {
					n.htmlType = li
					WriteOpenTag(htmlTypes[n.htmlType], &buffer)
					buffer.Write(line[n.textStart:])

					//WriteCloseTag(htmlTypes[n.htmlType], &buffer)

				} else {
					for len(nodeStack) > 1 {
						if n.level < nodeStack[len(nodeStack)-1].level {
							WriteCloseTag(htmlTypes[nodeStack[len(nodeStack)-1].htmlType], &buffer)
							nodeStack = nodeStack[:len(nodeStack)-1]
						} else {
							n.htmlType = li
							WriteOpenTag(htmlTypes[n.htmlType], &buffer)
							buffer.Write(line[n.textStart:])

							//WriteCloseTag(htmlTypes[n.htmlType], &buffer)
							break
						}
					}
				}
				if n.lineType != nodeStack[len(nodeStack)-1].lineType && n.lineType != textLine {
					nodeStack = append(nodeStack, n)
				}

			}

		}
		lastNode = n

	}
	return buffer.Bytes()
}
