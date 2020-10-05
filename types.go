package ctgov

type lineType int16

const (
	unkLine     lineType = 0
	emptyLine   lineType = 1
	textLine    lineType = 2
	dashLine    lineType = 3
	numberLine  lineType = 4
	commentLine lineType = 5
)

type htmlType int16

const (
	unk htmlType = 0
	p   htmlType = 1
	ul  htmlType = 2
	li  htmlType = 3
	ol  htmlType = 4
	co  htmlType = 5
)

var htmlTypes = []string{
	"",
	"p",
	"ul",
	"li",
	"ol",
	"p",
}

func getOpenCloseTag(t lineType) htmlType {

	switch t {

	case textLine:
		return p

	case dashLine:
		return ul

	case numberLine:
		return ol

	case commentLine:
		return co
	}

	return unk
}

func getHtmlTag(t htmlType) string {
	return htmlTypes[t]
}
