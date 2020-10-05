package ctgov

func newParagraph(n node) node {

	pp := node{
		lineType: n.lineType,
		htmlType: p,
		idx:      n.idx,
	}

	pp.contents = []content{n.content}

	return pp
}

func newComment(n node) node {

	pp := node{
		lineType: n.lineType,
		htmlType: co,
		idx:      n.idx,
	}

	pp.contents = []content{n.content}

	return pp
}
