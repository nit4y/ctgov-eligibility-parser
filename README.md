# CTgov Eligibility Criteria Parser
This is a parser written in Go for the eligibility criteria section in [CTgov]([url](https://clinicaltrials.gov/)).

## How to install
    
1. In ctgov package folder: "go install"

2. Should now be available for import.

## Usage example:

    import (
        ctgov github.com/nit4y/ctgov-eligibility-parser
    )
    ...
    p := ctgov.NewParser()
    ctString := strings.NewReader("some ct gov criteria")
    b, err := p.Parse(ctString)
    if err != nil {
        fmt.Printf("error occured: [%s], err)
    }
    fmt.Print(string(b))
    

## General flow

1. Reads line by line from input buffer.

2. Creates a node instance and stores in it relevant data for parsing the line.

3. For each node, writes html lists and elements considering line indentation in compare to last node in stack (parent) and the former node.

    - If line is indented forward, open tag for another list and write text as a new item.

        - Add current node to stack.

    - If line is indented backwards, pop items from stack and close tags of every node in stack, until indentation is the same.

    - if line is indented the same, write text as a item in this list.

5. Finally, close all open tags left in stack.

Note: parser relies on that every two line is seperated with an empty line.
