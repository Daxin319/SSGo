## Documentation for SSGo

### Packages

#### main

The main package contains only the actual main program. This bit of code is responsible for the top level commands, calling all sub-packages/functions in order to properly process the input.

It starts by setting the `basePath` variable depending on whether or not the `serve` optional argument was called. The `basePath` variable is usually a directory, and it should be expressed as `/dirName`, as the internal logic of the program will handle the trailing `/`.

If it was not provided, the program will set the base path to the provided argument, otherwise it is left blank. This allows you to call different base paths depending on where you would like to host your webpage. I'm using gitpages for my personal use, so I call the program with `SSGo /SSGo` in order to have the transpiler build the html with the appropriate path.

After setting the base path it gets the working directory and saves it to the `path` variable, then passes that variable to `fileio.CopyStaticToDocs(path)` which is the function that handles copying all data from the static dir to the docs dir.

It then calls the recursive page generation function that crawls the entire `/content` directory and parses all markdown into a mirrored structure in the `/docs` directory as HTML.

Finally, if the `serve` argument was used, then the program will spin up a server on `localhost:3000` and serve the contents of the `/docs` directory to it so you can check changes to your markdown before committing the final upload to your host of choice.

#### fileio

The fileio package handles all file system interactions. This includes clearing all old copies of the file structure, copying static files, reading all of the input and writing all of the output for the program.

##### `func fileio.CopyStaticToDocs(path string) error`

This function takes the current working directory as a string as input, and returns nothing unless there's a problem. It creates the origin and target path strings and copies all data from the origin to the target.

##### `func GeneratePagesRecursive(fromDirPath, destDirPath, templatePath, basePath string)`

This is the recursive function that handles the actual parse calls. It opens the origin directory, and for each entry in the directory it checks whether it's a markdown file or a directory. If it's a directory it creates a copy of the directory in the target location, then calls `GeneratePagesRecursive` with the updated origin and target directories. Otherwise it simply calls `generatePage` and returns to the previous call in the recursion stack.

  ##### `func generatePage(fromDirPath, destDirPath, templatePath, basePath string)`

  This is the actual individual page generation function. It opens the file at the origin location and reads it, creates a temporary variable that stores the data from the template and read that data, extract the H1 header from the top of the file and separate it into the title and content, call `html.MarkdownToHTMLNode` on the content to creat the AST, then walk the tree with `node.ToHTML()` finally it injects the title and content into the temporary template file, makes corrections to links and images as well as the CSS URIs in the HTML, and finally writes it all to disk.

#### html

This package handles the actual conversion of markdown to the HTML AST. 

##### `func MarkdownToHTMLNode(input string) nodes.Textnode`

The first thing it does is it replaces newlines with the html line break, then sanitizes null characters for security reasons using `blocks.SanitizeNulls`.

Now it gets into the meat of the work, it starts by initializing an empty `nodes.TextNode` struct, breaks the markdown into block levels with `blocks.MarkdownToBlocks`, and initializes an empty slice of `nodes.TextNode`s.

It then cycles through each block, and assigns a block type, then creates the appropriate node and appends it to the slice. Some nodes have children and some are explicitly leaf nodes and have no children and all of this is handled in the switch statement.

#### blocks

This package handles all block level parsing. Currently a WIP, I'll add documentation for this bit when I finish it.

#### inline

This package handles all of the inline level parsing.

##### `func TextToTextNodes(s string) []nodes.TextNode` 

This function takes a block of markdown as input, tokenizes the block, parses the stack into an AST, and returns the AST in the form of a slice of `nodes.TextNode`s

##### `func TextToChildren(s string) []nodes.TextNode`

Simply calls `TextToTextNodes`, only really here to help me keep the logic straight during development.

##### `func ParseInlineStack(tokens []tokenizer.Token) []nodes.TextNode`

Probably the most complex function in the project at the moment.

It begins by initializing empty slices for the AST output and a small stack to track delimiter runs.

Then I define a function that wraps any children in the parent html tag using a simple switch statement

Next I have a function to handle repeating delimiters. It started as an asterisk only func and has gotten more complicated but I haven't renamed it yet. It simply creates a delimiter run for repeating delimiters and finds a matching ending delimiter, if it's unable to find one it treats the opening run as literal characters.

Finally I start the main logic, which is a single pass over the token stack, using a switch statement to build the AST based on the token types.

#### tokenizer

This package is the inline tokenizer.

##### `func TokenizeInline(input string) []Token`

This function takes a block of markdown text and parses the inline formatting into a token stack, so that the inline parser can output the AST. To properly handle unicode characters, I have to cast the input string to a rune slice, and then I single pass the rune slice, first checking for raw HTML and falling back to autolinks if it's notraw HTML, then checking for code spans, and finally handling all of the other delimiters.

#### nodes

This package handles the node level AST work

##### `func UnescapeString(s string) string` 

Since certain elements are only escaped in certain contexts, the `UnescapeString` function is located in the nodes package. This handles contexts where I need to leave the literal text entered by the user contextually.

##### `func MapToHTMLChildren(children []TextNode, depth int) []TextNode`

This function is used to iterate through any children of a parent node and convert them to HTML nodes.

##### `func textNodeToHTMLNodeInternal(n TextNode, depth int) TextNode`

Used to convert a single node to an HTML node, checks text type and creates children/props as needed.

##### `type enum int`

This is Go, there are no real enums, I had to make them.

##### `type TextNode struct`

The TextNode struct that represents all possible data.

##### `type HTMLNode interface`

This interface means that the node is being converted to HTML, since I plan to add other target outputs later, I made this interface to start with the organization.

##### `func escapeHTML(s string) string`

A simple function to escape characters in an HTML string so that they are properly interpreted by the browser.

##### `func String(t enum) string`

This func is only here for debugging, but it returns the string version of the fake enum.

##### `func (h *TextNode) Repr() string`

simple formatting function for debugging so I can see my actual nodes and what data is in them.

##### `func (h *TextNode) PropsToHTML() string`

One of the 2 required methods for the `HTMLNode` interface. This properly formats any properties provided in the markdown (like the url for an `<a href= >` tag) for HTML output.

##### `func (h *TextNode) ToHTML() string`

The function that actually makes the change. This converts an individual node from a node to an HTML string, but it also makes sure that it handles all props and children before returning the final string
