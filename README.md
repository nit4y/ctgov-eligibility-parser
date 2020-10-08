# ctgov_parser

flow
1. iterate over every line
2. first item - this is base indent, children of root
3. save in stack last taken care of children's parent
    a. if line has higher level, append as child of last item in stack
    b. if line has lower level
        - pop stack
        - once hits a node with lower level then his, append as child
    c. else (same level) put as brother (child of parent)

no nodes - just render per case
back indent - close the list and remember and indent index

two steps
1. indendenatation normalization
2. tree building
 mark down auto parsing

