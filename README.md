# CTgov Parser

How to install
    
    1. In ctgov package folder: "go install"

    2. Should now be available for import.

Parser flow (macro):

    1. Reads line by line from input buffer.

    2. Creates a node instance and stores in it relevant data for parsing the line.

    3. For each node, creates elements and html lists as for line indentation, considering the last opened list, and the former node.

        - If text is indented forward, open tag for another list and write text as a new item.

                - Add node to stack.

        - If text is indented backwards, close tags of every node in stack, and pop items from stack until indentation is the same.

        - if text is indented the same, write text as a item in this list.

    5. Finally, closes all open tags left in stack.
