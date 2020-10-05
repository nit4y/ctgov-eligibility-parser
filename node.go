package ctgov

type node struct {
	lineType lineType
	htmlType htmlType
	idx      int
	closed   bool
	content

	contents []content
	children []node
}

func parseLine(line []byte) node {

	n := node{
		lineType: unkLine,
		htmlType: unk,
	}

	n.line = make([]byte, len(line))
	copy(n.line, line)

	if len(line) > 0 {
		for i := 0; i < len(line); i++ {
			c := line[i]
			if c == ' ' {
				continue
			} else if c >= '0' && c <= '9' {
				if n.lineType == unkLine { // still dont know, check to see if its number line
					for j := i; j < len(line); j++ {
						var cc = line[j]
						if cc >= '0' && cc <= '9' {
							continue
						} else if cc == '.' {
							n.start = i
							n.lineType = numberLine
							i = j
							break
						} else {
							i = j
							break
						}
					}
				} else if n.lineType == dashLine || n.lineType == commentLine { // already know, quit
					n.textStart = i
					break
				}
			} else if c == '-' {
				n.start = i
				n.lineType = dashLine
			} else if c == '*' {
				n.start = i
				n.lineType = commentLine
			} else { // found text somewhere down the road

				if n.start == 0 { // first time
					n.start = i
				}

				n.textStart = i
				if n.start == n.textStart { // text == start -> this is a text line
					n.lineType = textLine
				}
				break
			}
		}
	} else {
		n.lineType = emptyLine
	}

	return n

}
