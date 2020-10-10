package ctgov

import (
	"bufio"
	"bytes"
	"io"
)

// Parser type decleration
type Parser struct {
}

// NewParser initiates a Parser instance (struct).
func NewParser() *Parser {
	return &Parser{}
}

// Parse iterates over ctgov criteria section, builds html presentation of that data and returns as buffer.
// HTML output structure consists of indentation and small irregularities found during testing.
// The tree is built along the way while the tree stack is managed dynamicly.
func (pa *Parser) Parse(r io.Reader) []byte {
	scanner := bufio.NewScanner(r)
	var (
		buffer     bytes.Buffer
		treeStack  []*node
		formerNode *node
	)

	// Add empty root node
	formerNode = newNode(0, unkLine, unk)
	treeStack = append(treeStack, formerNode)

	for scanner.Scan() {
		line := scanner.Bytes()
		lType, textStart := calcNodeProps(line, formerNode)
		n := newNode(calcLevel(line), lType, calcHTMLType(lType)) // level is the indentation counter before text starts.
		n.textStart = textStart

		switch n.lineType {
		case emptyLine:
			if formerNode != newNode(0, unkLine, unk) { // making sure its not the first line
				WriteCloseTag(htmlTypes[li], &buffer)
			}

		case textLine, commentLine:
			for len(treeStack) > 0 {
				if n.level < treeStack[len(treeStack)-1].level {
					WriteCloseTag(htmlTypes[treeStack[len(treeStack)-1].htmlType], &buffer)
					treeStack = treeStack[:len(treeStack)-1]

				} else if n.level > treeStack[len(treeStack)-1].level {
					buffer.WriteString(" ")
					buffer.Write(line[n.level:])
					treeStack = append(treeStack, n)
					break

				} else {
					buffer.WriteString(" ")
					buffer.Write(line[n.level:])
					break

				}
			}

		case numberLine, dashLine:
			if len(treeStack) > 0 {
				if n.level > treeStack[len(treeStack)-1].level {
					WriteOpenTag(htmlTypes[n.htmlType], &buffer)
					WriteOpenTag(htmlTypes[li], &buffer)
					buffer.Write(line[n.textStart:]) // write data only after indentation and node numbering

				} else if n.level == treeStack[len(treeStack)-1].level {
					n.htmlType = li
					WriteOpenTag(htmlTypes[n.htmlType], &buffer)
					buffer.Write(line[n.textStart:])

				} else {
					for len(treeStack) > 1 {
						if n.level < treeStack[len(treeStack)-1].level &&
							treeStack[len(treeStack)-1].level-n.level != 1 { // second condition prevents common anomaly when numbering exceeds 9.
							WriteCloseTag(htmlTypes[treeStack[len(treeStack)-1].htmlType], &buffer)
							treeStack = treeStack[:len(treeStack)-1]

						} else {
							n.htmlType = li
							WriteOpenTag(htmlTypes[n.htmlType], &buffer)
							buffer.Write(line[n.textStart:])
							break
						}
					}
				}
				if n.level > treeStack[len(treeStack)-1].level {
					treeStack = append(treeStack, n)
				}
			}
		}
		formerNode = n
	}
	return buffer.Bytes()
}
