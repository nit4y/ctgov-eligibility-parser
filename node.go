package ctgov

// node is a blob struct for storing relevant parsing metadata about a line in file
type node struct {
	level     int
	textStart int
	lineType
	htmlType
}

// newNode initiates a node instance.
func newNode(level int, t lineType, hType htmlType) *node {
	n := node{level: level, lineType: t, htmlType: hType}
	return &n
}

// calcNodeProps calculates relevant parsing metadata of line
func calcNodeProps(line []byte, lastNode *node) (lineType, int) {

	var retType = unkLine
	var start int
	var textStart int

	if len(line) > 0 {
		for i := 0; i < len(line); i++ {
			c := line[i]
			if c == ' ' {
				continue
			} else if c >= '1' && c <= '9' {
				if retType == unkLine { // still dont know, check to see if its number line
					for j := i; j < len(line); j++ {
						var cc = line[j]
						if cc >= '0' && cc <= '9' {
							continue
						} else if cc == '.' {
							start = i
							if lastNode.lineType == emptyLine || lastNode.lineType == unkLine {
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
				if lastNode.lineType == emptyLine || lastNode.lineType == unkLine {
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
	return retType, textStart
}

// calcLevel determines line "level" which is the number of indentations before the actual text is.
func calcLevel(line []byte) int {
	counter := 0
	for i := 0; i < len(line); i++ {
		if line[i] == '\t' {
			counter = counter + 4
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
