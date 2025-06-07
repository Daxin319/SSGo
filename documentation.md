# SSGo Documentation

SSGo is a static site generator that converts Markdown files to HTML, with support for extended Markdown features and custom styling. It aims to provide a simple yet powerful tool for creating static websites with a focus on maintainability and extensibility.

## Table of Contents

- [Quick Start Guide](#quick-start-guide)
- [Overview](#overview)
- [Architecture](#architecture)
- [Packages](#packages)
  - [Main](#main)
  - [FileIO](#fileio)
  - [HTML](#html)
  - [Blocks](#blocks)
  - [Inline](#inline)
  - [Tokenizer](#tokenizer)
  - [Nodes](#nodes)
- [Markdown Support](#markdown-support)
- [Contributing](#contributing)

## Quick Start Guide

This program is designed to copy static resources into the `/docs` directory inside the repository from the `/static` directory, then process any markdown files in the `/content` directory into HTML which is then also put in the `/docs` directory to be served to localhost or formatted for your host service of choice.

### Prerequisites

- Go version 1.22 or later
- Git
- A text editor
- Basic knowledge of Markdown

### Installation

1. Install Go version 1.22 or later:
   ```bash
   curl -sS https://webi.sh/golang | sh
   ```
   Alternatively, follow the [official Go installation guide](https://go.dev/doc/install).

2. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/SSGo.git
   cd SSGo
   ```

### Usage

1. **Static Resources**
   - Place static assets in the `/static` directory
   - Organize resources in subdirectories as needed for website navigation

2. **Content**
   - Create markdown files in the `/content` directory
   - Use `index.md` as the main file for each page
   - Maintain the same directory structure as your desired website

3. **Development**
   - Run `. serve.sh` to generate pages and serve locally:
     ```bash
     ./serve.sh
     ```
   - Access your site at `http://localhost:3000`

4. **Production**
   - Run `. build.sh` to generate pages for deployment:
     ```bash
     ./build.sh
     ```
   - Edit `build.sh` to set your repository name (replace `/SSGo` with your path)
   - Deploy the contents of `/docs` to your hosting service

## Overview

SSGo is built with a modular architecture that separates concerns into distinct packages, making the codebase maintainable and extensible. The conversion process follows these steps:

1. File System Operations
   - Copy static assets
   - Process markdown files
   - Generate output structure

2. Markdown Processing
   - Parse blocks
   - Process inline elements
   - Generate AST

3. HTML Generation
   - Convert AST to HTML
   - Apply templates
   - Handle assets

## Architecture

The project is organized into several packages, each handling a specific aspect of the markdown-to-HTML conversion process:

1. `main` - Entry point and high-level program flow
2. `fileio` - File system operations
3. `html` - HTML AST generation
4. `blocks` - Block-level markdown parsing
5. `inline` - Inline markdown parsing
6. `tokenizer` - Token generation for inline parsing
7. `nodes` - AST node definitions and HTML conversion

## Packages

### main

The main package orchestrates the entire conversion process.

#### `func main()`

The entry point that:
1. Sets the `basePath` based on the `serve` argument
2. Copies static files to the docs directory
3. Recursively processes markdown files
4. Optionally starts a development server

The `basePath` is used to handle different hosting scenarios (e.g., GitHub Pages requires `/SSGo`).

### fileio

Handles all file system operations.

#### `func CopyStaticToDocs(path string) error`

Copies static assets from the source directory to the destination directory.

#### `func GeneratePagesRecursive(fromDirPath, destDirPath, templatePath, basePath string)`

Recursively processes directories and files:
- Creates mirror directory structure
- Converts markdown files to HTML
- Preserves static assets

#### `func generatePage(fromDirPath, destDirPath, templatePath, basePath string)`

Handles individual page generation:
1. Reads markdown content
2. Extracts title from H1 header
3. Converts content to HTML AST
4. Applies template
5. Fixes asset paths
6. Writes output file

### html

Converts markdown to HTML AST.

#### `func MarkdownToHTMLNode(input string) nodes.TextNode`

The main conversion function that:
1. Sanitizes input
2. Splits into blocks
3. Creates AST nodes
4. Handles block types:
   - Headers
   - Paragraphs
   - Lists
   - Code blocks
   - Blockquotes
   - HTML blocks
   - Horizontal rules

### blocks

Handles block-level markdown parsing. Currently supports basic block types with some limitations:

#### `func SanitizeNulls(input string) string`

Security function that removes null characters from input.

#### `func MarkdownToBlocks(input string) []string`

Splits markdown into blocks, handling:
- Headers (H1-H6)
- Paragraphs
- Blockquotes (single level only)
- Lists (single level only)
  - Ordered lists (1. 2. 3.)
  - Unordered lists (- * +)
- Code blocks
  - Fenced code blocks (``` or ~~~)
- Horizontal rules (- * _)
- Empty lines (block separators)

Limitations (at this time):
- Nested lists are not supported
- Nested blockquotes are not supported
- List items cannot contain block-level elements
- Task lists are not supported
- Definition lists are not supported
- HTML blocks are not supported
- Indented code blocks are not supported

### inline

Handles inline markdown parsing.

#### `func TextToTextNodes(s string) []nodes.TextNode`

Converts inline markdown to AST nodes:
1. Tokenizes input
2. Parses token stack
3. Returns AST nodes

#### `func TextToChildren(s string) []nodes.TextNode`

Helper function that calls `TextToTextNodes`.

#### `func ParseInlineStack(tokens []tokenizer.Token) []nodes.TextNode`

Complex function that:
1. Initializes AST output and delimiter stack
2. Processes tokens in a single pass
3. Handles nested formatting
4. Builds AST based on token types

### tokenizer

Generates tokens for inline parsing.

#### `func TokenizeInline(input string) []Token`

Converts markdown text to tokens, handling:
- Raw HTML elements
- Autolinks (URLs and email addresses)
- Code spans
- HTML entities and character references
- Emphasis (bold, italic, bolditalic)
- Strikethrough
- Subscript and superscript
- Highlighting
- Links and images
- Escaped characters

#### `func parseEntity(runes []rune, i int) (string, int)`

Handles HTML entities and character references:
- Named entities (e.g., `&amp;`, `&lt;`, `&copy;`)
- Decimal character references (e.g., `&#169;`, `&#8212;`)
- Hexadecimal character references (e.g., `&#xA9;`, `&#x2014;`)
- Invalid entities (treated as literal text)

### nodes

Defines AST nodes and handles HTML conversion.

#### Types

##### `type enum int`

Represents node types in the AST.

##### `type TextNode struct`

Core AST node structure containing:
- Node type
- Content
- Children
- Properties
- HTML conversion methods

##### `type HTMLNode interface`

Interface for HTML conversion, allowing for future output formats.

#### Functions

##### `func UnescapeString(s string) string`

Handles context-specific string unescaping.

##### `func MapToHTMLChildren(children []TextNode, depth int) []TextNode`

Recursively converts child nodes to HTML.

##### `func textNodeToHTMLNodeInternal(n TextNode, depth int) TextNode`

Converts a single node to HTML.

##### `func escapeHTML(s string) string`

Escapes HTML special characters.

##### `func String(t enum) string`

Debug helper for enum string representation.

#### Methods

##### `func (h *TextNode) Repr() string`

Debug helper for node string representation.

##### `func (h *TextNode) PropsToHTML() string`

Formats node properties as HTML attributes.

##### `func (h *TextNode) ToHTML() string`

Converts node and its children to HTML string.

## Markdown Support

SSGo currently supports approximately 70% of the CommonMark specification, plus several extended features:

### Supported Features

- Basic block elements
  - Headers (H1-H6)
  - Paragraphs
  - Blockquotes
  - Lists
  - Code blocks
  - Horizontal rules

- Inline elements
  - Emphasis (bold, italic)
  - Links and images
  - Code spans
  - HTML entities
  - Autolinks

### Extended Features

- Strikethrough (`~~text~~`)
- Subscript (`~text~`)
- Superscript (`^text^`)
- Highlighting (`==text==`)
- Hard line breaks (two spaces or backslash at end of line)

### Autolinks and Raw HTML

- **GFM-style autolinks** are supported:
  - Bare domains (e.g., `<example.com>`, `<sub.example.com>`, `<example.com/path>`, `<example.com:8080>`, etc.)
  - Protocol autolinks (e.g., `<http://example.com>`, `<https://example.com>`, etc.)
  - Email autolinks (e.g., `<user@example.com>`, `<user:pass@example.com>`, etc.)
  - Authentication in email autolinks is supported, but passwords are stripped from the rendered link and only the username and domain are shown/linked.
  - Autolinks are only created if the TLD is a real, recognized public suffix (using the publicsuffix library).
  - Protocol autolinks require at least one character after the protocol (e.g., `<http://>` is not autolinked).
  - Invalid/fake TLDs (like `.url`, `.test`, `.invalid`, `.localhost`) are not autolinked.
  - Reserved/fake TLDs are not autolinked.

- **Raw HTML blocks and comments** are preserved and rendered as HTML, not escaped. This includes:
  - HTML tags (e.g., `<div>`, `<span style="color: red;">`, etc.)
  - HTML comments (e.g., `<!-- comment -->`)
  - HTML blocks are not wrapped in `<p>` or escaped, and are passed through to the output as-is.

- **Block splitting** has been improved:
  - Each autolink or raw HTML tag/comment on its own line is treated as its own block, ensuring correct parsing and rendering.

### Planned Features

- Nested lists and blockquotes
- Task lists
- Definition lists
- HTML blocks (multi-line)
- Indented code blocks

## Contributing

This has been a solo project for my portfolio, but if you see something please feel free to reach out to me and let me know you've seen something, then I'll happily discuss it with you, but I don't want you to fix it for me.