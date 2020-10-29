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

	var (
		buffer     bytes.Buffer
		treeStack  []*node
		formerNode *node
	)

	// Add empty root node
	formerNode = newNode(0, unkLine, unk)
	treeStack = append(treeStack, formerNode)

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Bytes()
		lType, textStart := calcNodeProps(line, formerNode)
		n := newNode(calcLevel(line), lType, calcHTMLType(lType)) // level is the indentation counter before text starts.
		n.textStart = textStart

		switch n.lineType {

		case emptyLine:

			if formerNode.lineType == numberLine || formerNode.lineType == dashLine ||
				formerNode.lineType == textLine && formerNode.textStart == treeStack[len(treeStack)-1].textStart { // making sure its not the first line

				if formerNode.level == treeStack[len(treeStack)-1].level && treeStack[len(treeStack)-1].htmlType == p { // close p tag only if it is a paragraph

					writeCloseTag(htmlTypes[p], &buffer)

					treeStack = treeStack[:len(treeStack)-1]

					if treeStack[len(treeStack)-1].lineType == numberLine &&
						formerNode.level <= treeStack[len(treeStack)-1].level ||
						treeStack[len(treeStack)-1].lineType == dashLine &&
							formerNode.level <= treeStack[len(treeStack)-1].level { // if paragraph is a part of a list item, close list tags properly as well.

						writeCloseTag(htmlTypes[li], &buffer)
						writeCloseTag(htmlTypes[treeStack[len(treeStack)-1].htmlType], &buffer)

						treeStack = treeStack[:len(treeStack)-1]

					}

				} else {
					writeCloseTag(htmlTypes[n.htmlType], &buffer)
				}

			}

		case textLine, commentLine:

			for len(treeStack) > 0 { // while stack still contains nodes

				if n.level < treeStack[len(treeStack)-1].level { // if text is indented backwards, close node tag and pop from stack
					writeCloseTag(htmlTypes[treeStack[len(treeStack)-1].htmlType], &buffer)
					treeStack = treeStack[:len(treeStack)-1]

				} else {

					if len(treeStack) == 1 || formerNode.lineType == emptyLine { // if only root node found in stack, the line is a header text line, preventing root node pop
						n.htmlType = p
						writeOpenTag(htmlTypes[n.htmlType], &buffer)
						treeStack = append(treeStack, n)

					} else {
						if n.lineType == commentLine {
							buffer.WriteString("<br>") // comment lines appear below items. br = line break
						} else {
							buffer.WriteString(" ") // all other data appear inline.
						}

					}

					buffer.Write(line[n.level:]) // write from level instead of textStart because a text line might start with numbering or dash.

					break

				}
			}

		case numberLine, dashLine:

			if n.level > treeStack[len(treeStack)-1].level {
				// indent forward
				// in every second consecutive numbering indent, make the ol element an "a, b, c" ordered list.
				if treeStack[len(treeStack)-1].lineType == numberLine && n.lineType == numberLine {
					writeOpenTag("ol type=\"a\"", &buffer)

				} else {
					writeOpenTag(htmlTypes[n.htmlType], &buffer)
				}

				writeOpenTag(htmlTypes[li], &buffer)
				buffer.Write(line[n.textStart:]) // write text found only after indentation and node numbering

			} else if n.level == treeStack[len(treeStack)-1].level { // text is indented the same
				// write line as a item in current open list.
				n.htmlType = li
				writeOpenTag(htmlTypes[n.htmlType], &buffer)
				buffer.Write(line[n.textStart:])

			} else { // has lower level, indent backwards

				for len(treeStack) > 1 { // continue popping items until reaching first parent node (lowest level possible), meant to prevent popping of root node

					if n.level < treeStack[len(treeStack)-1].level &&
						treeStack[len(treeStack)-1].level-n.level != 1 { // second condition prevents common anomaly when numbering exceeds 9.
						writeCloseTag(htmlTypes[treeStack[len(treeStack)-1].htmlType], &buffer)
						treeStack = treeStack[:len(treeStack)-1]

					} else { // reached back to correct level, write item in list.
						n.htmlType = li
						writeOpenTag(htmlTypes[n.htmlType], &buffer)
						buffer.Write(line[n.textStart:])

						break
					}

				}

			}

			if n.level > treeStack[len(treeStack)-1].level { // only parent nodes are needed in stack
				treeStack = append(treeStack, n)
			}

		}

		formerNode = n

	}
	// Close all open html tags.
	// closing open li tag if present
	if treeStack[len(treeStack)-1].lineType == numberLine || treeStack[len(treeStack)-1].lineType == dashLine {
		writeCloseTag(htmlTypes[li], &buffer)
	}

	// closing ol, ul open tags if present
	for len(treeStack) > 1 {
		writeCloseTag(htmlTypes[treeStack[len(treeStack)-1].htmlType], &buffer)
		treeStack = treeStack[:len(treeStack)-1]
	}

	return buffer.Bytes()
}
