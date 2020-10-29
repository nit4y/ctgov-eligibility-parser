package ctgov

import "bytes"

// writeOpenTag writes to buffer the opening html tag el as syntax should be.
func writeOpenTag(el string, buffer *bytes.Buffer) {
	buffer.WriteString("<")
	buffer.WriteString(el)
	buffer.WriteString(">")
}

// writeCloseTag writes to buffer the closing html tag el as syntax should be.
func writeCloseTag(el string, buffer *bytes.Buffer) {
	buffer.WriteString("</")
	buffer.WriteString(el)
	buffer.WriteString(">")
}
