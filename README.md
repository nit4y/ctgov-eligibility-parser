# CTgov Parser

How to install
    
    1. In ctgov package folder: "go install"

    2. Should now be available for import.

Parser flow:

    1. Reads line by line a buffer.

    2. Creates a node instance that is relevant data for parsing the line.

    3. For each node, creates elements and html lists as for line indentation.

        - If text is indented forward, open tag for another list and write text as a new item.

        - If text is indented backwards, close tags of every parent node in stack, and pop items from stack until indentation is corrent.

        - if text is indented the same, write text as a list item.

    4. Inserts tree structure "parent" nodes only to a stack.

    5. Finally, closes all open tags left in stack.
