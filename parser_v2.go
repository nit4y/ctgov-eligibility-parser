package ctgov

import (
	"bufio"
	"io"
)

type ParserV2 struct {
}

func NewParserV2() *ParserV2 {
	return &ParserV2{}
}

// parse ctgov raw textBlock into better looking html
func (pa *ParserV2) Parse(r io.Reader) ([]byte, int) {

	scanner := bufio.NewScanner(r)

	var (
		current *node
		all     []node
		stack   []node
		idx     = 0
	)

	for scanner.Scan() {
		line := scanner.Bytes()

		n := parseLine(line)
		n.htmlType = getOpenCloseTag(n.lineType)
		n.idx = idx

		idx++
		switch n.lineType {
		case emptyLine:


			if current != nil {

				if len(current.children) == 0 {
					current.closed = true
				} else {
					current.children[len(current.children) - 1].closed = true
				}

				//if current.htmlType == p {
				//	all = append(all, *current)
				//	current = nil
				//} else if current.htmlType == co {
				//	if len(stack) > 0 {
				//		var temp = *current
				//		for {
				//			current, stack = &stack[len(stack)-1], stack[:len(stack)-1]
				//
				//			// can be just one comment
				//			// replace it
				//
				//			lastChild := current.children[len(current.children)-1]
				//			if len(lastChild.children) == 0 {
				//				current.children[len(current.children)-1] = temp
				//			} else {
				//				grand := len(lastChild.children) - 1
				//				current.children[len(current.children)-1].children[grand] = temp
				//			}
				//
				//			//temp = *current
				//			if len(stack) == 0 {
				//
				//				// done popping, here and before replace the current
				//				break
				//			}
				//
				//		}
				//	}
				//}
			}
			break

		case textLine, commentLine:

			//if n.lineType == commentLine {
			//	fmt.Println(idx, current)
			//	fmt.Println(current.lineType, current.htmlType, len(current.contents), len(current.children))
			//}

			if current == nil {
				current = &n
			}


			if current.htmlType == p {
				if current.closed {
					all = append(all, *current)
					current = nil
				} else {
					current.contents = append(current.contents, n.content)
				}

			} else if current.htmlType == co {
				//current.contents = append(current.contents, n.content)
			} else { // list

				l := len(current.children) - 1
				if n.textStart != current.textStart {

					if n.textStart == current.textStart ||
						(l != -1 && n.textStart == current.children[l].textStart) {

						current.children[l].contents = append(current.children[l].contents, n.content)
					} else if n.textStart < current.textStart ||
						(l != -1 && n.textStart < current.children[l].textStart) { // back indent

						if len(stack) > 0 {
							var temp = *current
							for {
								current, stack = &stack[len(stack)-1], stack[:len(stack)-1]

								// append the former-current to the current
								current.children = append(current.children, temp)
								temp = *current
								if len(stack) == 0 {

									// done popping, here and before replace the current
									break
								}

							}
						}

						// before replace, add the current to the list..
						all = append(all, *current)

						pp := newParagraph(n)
						current = &pp

					} else {

						//fmt.Println(len(stack), n.lineType)
						if n.lineType == commentLine && len(stack) == 0 { // for comments and not parents,
							pp := newComment(n)
							current.children[l].children = append(current.children[l].children, pp)

							stack = append(stack, *current)

							current = &pp
						} else { //
							current.children[l].contents = append(current.children[l].contents, n.content)
						}

					}

				} else {
					current.children[l].contents = append(current.children[l].contents, n.content)
				}
			}

			break

		case dashLine, numberLine:

			if current != nil {

				if current.htmlType == p { // p after dash indicates replace

					all = append(all, *current)

					current = &n
					current.contents = append(current.contents, n.content)

				} else {

					if n.textStart > current.textStart { // indent ul --> li - creating li
						lli := newListItem(n)
						n.children = append(n.children, lli)

						stack = append(stack, *current)

						current = &n

					} else if n.textStart < current.textStart { // back indent
						var temp = *current

						// pop
						for {
							current, stack = &stack[len(stack)-1], stack[:len(stack)-1]

							// append the former current to the current
							current.children = append(current.children, temp)

							if len(stack) == 0 {
								break
							}

							temp = *current
						}

						lli := newListItem(n)

						// append the new li
						current.children = append(current.children, lli)

					} else { // change type and adjust the content correctly

						n.htmlType = li
						n.contents = append(n.contents, n.content)
						current.children = append(current.children, n)

					}
				}
			} else { // no current, start one

				if n.htmlType == ul || n.htmlType == ol {
					lli := newListItem(n)
					n.children = append(n.children, lli)
					current = &n
				}

			}

			break
		}

	}

	if current != nil {
		all = append(all, *current)
	}

	return generateHtml(all), idx
}
