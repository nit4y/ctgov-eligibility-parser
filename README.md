# CTgov Parser

## How to install
    
1. In ctgov package folder: "go install"

2. Should now be available for import.

## Usage example:
    p := ctgov.NewParser()
    ctString := strings.NewReader("some ct gov criteria")
    ret := p.Parse(st)
    fmt.Print(string(ret))

## Parser flow (macro):

1. Reads line by line from input buffer.

2. Creates a node instance and stores in it relevant data for parsing the line.

3. For each node, writes html lists and elements considering line indentation in compare to last node in stack (parent) and the former node.

    - If line is indented forward, open tag for another list and write text as a new item.

        - Add current node to stack.

    - If line is indented backwards, pop items from stack and close tags of every node in stack, until indentation is the same.

    - if line is indented the same, write text as a item in this list.

5. Finally, closes all open tags left in stack.

Note: parser relies on that every two line is seperated with an empty line.
