package ctgov

import (
	"bytes"
)

func sanitize(c content) []byte {

	if c.textStart != c.start {
		from := c.textStart
		to := len(c.line)

		return c.line[from:to]

	}

	return c.line

}

func generateEl(n node, buffer *bytes.Buffer) {

	el := getHtmlTag(n.htmlType)

	//start
	buffer.WriteString("<")
	buffer.WriteString(el)
	buffer.WriteString(">")

	if n.htmlType == co {
		buffer.WriteString("* ")
	}

	for i := range n.contents {
		c := n.contents[i]
		// sanitize
		txt := sanitize(c)
		buffer.Write(txt)
	}

	for i := range n.children {
		child := n.children[i]
		generateEl(child, buffer)
	}

	// end
	buffer.WriteString("</")
	buffer.WriteString(el)
	buffer.WriteString(">")

}

func generateHtml(nodes []node) []byte {

	buffer := bytes.Buffer{}

	for i := range nodes {
		n := nodes[i]
		generateEl(n, &buffer)

	}

	return buffer.Bytes()
}
