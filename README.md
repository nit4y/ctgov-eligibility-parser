<h1>CTgov Parser</h1>
Parser detailed flow:
    1. Gets a io.Reader instance 
    2. Start iterating over the reader line by line.
    3. Creates a node instance that is actually a blob for all relevant data over the current line.
    4. In case line is / part of:
        - Empty line:
            1. If node is not the first nodeCloses last <li> tag
        -

