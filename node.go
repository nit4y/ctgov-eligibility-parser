package ctgov

import "strconv"

// node is a blob struct for storing relevant parsing metadata about a line in file
type node struct {
	level     int // number of and tabs before line (space counts as 1, tab as 4)
	textStart int // actualStart + numbering / dashing characters
	lineType
	htmlType
	numberingValue int
}

// newNode initiates a node instance.
func newNode(level int, t lineType, hType htmlType) *node {
	n := node{level: level, lineType: t, htmlType: hType}
	return &n
}

// calcNodeProps calculates relevant parsing metadata of line
func calcNodeProps(line []byte, lastNode *node) (lineType, int, int) {

	var retType = unkLine
	var start int
	var end int
	var textStart int

	if len(line) > 0 {
		for i := 0; i < len(line); i++ {
			c := line[i]

			if c == ' ' || c == '\t' { // tabs may appear as indentation as well
				continue

			} else if c >= '1' && c <= '9' {
				if retType == unkLine { // still dont know, check to see if its a number line
					for j := i; j < len(line); j++ {
						var cc = line[j]

						if cc >= '0' && cc <= '9' {
							continue

						} else if cc == '.' {
							start = i
							end = j
							if lastNode.lineType == emptyLine || lastNode.lineType == unkLine { // a text line might start with a numbering as text
								retType = numberLine
							} else {
								retType = textLine
							}
							i = j

							break

						} else {
							i = j

							break

						}
					}

				} else if retType == dashLine || retType == commentLine { // already know, quit
					textStart = i
					break

				}

			} else if c == '-' {
				start = i
				if lastNode.lineType == emptyLine || lastNode.lineType == unkLine { // a text line might start with a dash
					retType = dashLine
				} else {
					retType = textLine
				}

			} else if c == '*' {
				start = i
				retType = commentLine

			} else { // found text somewhere down the road

				if start == 0 { // first time
					start = i
				}

				textStart = i

				if start == textStart { // text == start -> this is a text line
					retType = textLine
				}

				break
			}
		}

	} else {
		retType = emptyLine
	}

	numberingValue := 0

	if retType == numberLine {
		numberingValue, _ = strconv.Atoi(string(line[start:end]))
	}

	return retType, textStart, numberingValue
}

// calcLevel determines line "level" which is the number of indentations before the actual text is.
func calcLevel(line []byte) int {

	counter := 0

	for i := 0; i < len(line); i++ {

		if line[i] == '\t' {
			counter = counter + 4 // tabs are equal to 4 spaces in terms of indention.

		} else if line[i] == ' ' {
			counter = counter + 1

		} else {
			break

		}
	}

	return counter
}

// calcHTMLType determines html tag of input line with lineType t.
func calcHTMLType(t lineType) htmlType {
	switch t {

	case numberLine:
		return ol

	case dashLine:
		return ul

	case emptyLine:
		return li

	}

	return unk
}
