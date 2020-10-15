package ctgov

type lineType int16

const (
	unkLine          lineType = 0
	emptyLine        lineType = 1
	textLine         lineType = 2
	dashLine         lineType = 3
	numberLine       lineType = 4
	commentLine      lineType = 5
	romanNumeralLine lineType = 6
)

type htmlType int16

const (
	unk htmlType = 0
	p   htmlType = 1
	ul  htmlType = 2
	li  htmlType = 3
	ol  htmlType = 4
	co  htmlType = 5
	olr htmlType = 6
)

var htmlTypes = []string{
	"",
	"p",
	"ul",
	"li",
	"ol",
	"p",
	"ol type=\"i\"",
}
