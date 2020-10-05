package ctgov

func newListItem(n node) node {

	lli := node{
		lineType: n.lineType,
		htmlType: li,
		idx:      n.idx,
	}

	lli.contents = []content{n.content}

	return lli
}
