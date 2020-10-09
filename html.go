package ctgov

import "bytes"

// WriteOpenTag writes to buffer the opening html tag as syntax should be.
func WriteOpenTag(el string, buffer *bytes.Buffer) {
	buffer.WriteString("<")
	buffer.WriteString(el)
	buffer.WriteString(">")
}

// WriteCloseTag writes to buffer the closing html tag as syntax should be.
func WriteCloseTag(el string, buffer *bytes.Buffer) {
	buffer.WriteString("</")
	buffer.WriteString(el)
	buffer.WriteString(">")
}
