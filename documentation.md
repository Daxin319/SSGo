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
- [SSGo Markdown Specification](#ssgo-markdown-specification)
  - [Block Elements](#block-elements)
    - [Thematic Breaks](#thematic-breaks)
    - [ATX Headings](#atx-headings)
    - [Fenced Code Blocks](#fenced-code-blocks)
    - [Blockquotes](#blockquotes)
    - [Lists](#lists)
    - [Raw HTML Blocks](#raw-html-blocks)
  - [Inline Elements](#inline-elements)
    - [Emphasis and Strong Emphasis](#emphasis-and-strong-emphasis)
    - [Strikethrough](#strikethrough)
    - [Highlight](#highlight)
    - [Subscript and Superscript](#subscript-and-superscript)
    - [Code Spans](#code-spans)
    - [Links and Images](#links-and-images)
    - [Autolinks](#autolinks)
    - [Raw HTML](#raw-html)
    - [HTML Entities](#html-entities)
    - [Hard Line Breaks](#hard-line-breaks)
- [Unsupported Features](#unsupported-features)
- [Contributing](#contributing)

## Quick Start Guide

This program is designed to copy static resources into the `/docs` directory inside the repository from the `/static` directory, then process any markdown files in the `/content` directory into HTML which is then also put in the `/docs` directory to be served to localhost or formatted for your host service of choice.

### Prerequisites

- Go version 1.22 or later
- Git (I refuse to believe you don't already have git installed if you're reading this, but just in case: [Git download](https://git-scm.com/downloads) or on Linux: `sudo pacman -S git` or `sudo apt install git`)
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
   git clone https://github.com/Daxin319/SSGo.git
   cd SSGo
   ```

### Usage

1. **Static Resources**
   - Place static assets in the `/static` directory
   - Organize resources in subdirectories as needed for website navigation

2. **Content**
   - Create markdown files in the `/content` directory
   - Use subdirectories to organize your content

3. **Development**
   - Run `./serve.sh` to generate pages and serve locally:
     ```bash
     ./serve.sh
     ```
   - Access your site at `http://localhost:3000`

4. **Production**
   - Run `./build.sh` to generate pages for deployment:
     ```bash
     ./build.sh
     ```
   - If deploying to a subdirectory (e.g., GitHub Pages), edit `build.sh` to set the correct base path for your site.
   - Deploy the contents of the `/docs` directory to your hosting service.

## Overview

SSGo is built with a modular architecture that separates concerns into distinct packages, making the codebase maintainable and extensible. The conversion process follows these steps:

1.  **File System Operations**: The `fileio` package copies static assets from `/static` to `/docs` and then recursively finds Markdown files in `/content`.
2.  **Markdown Processing**: For each Markdown file, the content is parsed into blocks (paragraphs, headings, etc.). Each block is then processed to parse inline elements (bold, links, etc.).
3.  **AST Generation**: The parsing process generates an Abstract Syntax Tree (AST) where each element of the document is a node.
4.  **HTML Generation**: The `html` renderer traverses the AST and generates the final HTML output, which is then written to the `/docs` directory.

## Architecture

The project is organized into several packages, each handling a specific aspect of the markdown-to-HTML conversion process:

1.  `cmd/transpiler` - The main application entry point.
2.  `fileio` - Handles file system operations, like copying directories and generating pages.
3.  `renderer/html` - Converts the Markdown AST into HTML.
4.  `blocks` - Handles block-level Markdown parsing (e.g., paragraphs, lists).
5.  `inline` - Handles inline Markdown parsing (e.g., bold, italic, links).
6.  `tokenizer` - Converts raw text into a stream of tokens for the inline parser.
7.  `nodes` - Defines the AST node structures used to represent the document.

## Packages

### main

The `main` package in `cmd/transpiler` orchestrates the entire conversion process.

#### `func main()`

The entry point that:
1. Sets the `basePath` for assets based on command-line arguments.
2. Copies static files from `static/` to `docs/`.
3. Recursively processes Markdown files from `content/` to `docs/`.
4. Optionally starts a local development server if `serve` is passed as an argument.

### fileio

Handles all file system operations.

#### `func CopyStaticToDocs(path string) error`

Copies static assets from the `static/` directory to the `docs/` directory.

#### `func GeneratePagesRecursive(fromDirPath, destDirPath, templatePath, basePath string)`

Recursively processes directories and files from `/content` to `/docs`, creating a mirror directory structure.

#### `func generatePage(fromDirPath, destDirPath, templatePath, basePath string)`

Handles individual page generation:
1. Reads markdown content from a source file.
2. Extracts the H1 header to use as the page title.
3. Converts the remaining Markdown content to an HTML AST.
4. Injects the title and content into an HTML template.
5. Adjusts asset paths (`href` and `src`) to include the `basePath`.
6. Writes the final HTML to the destination file.

### html

Located in `renderer/html`, this package converts the Markdown AST to an HTML string.

#### `func MarkdownToHTMLNode(input string) nodes.TextNode`

The main conversion function that:
1. Sanitizes input by removing null characters and normalizing line breaks.
2. Splits the input string into blocks.
3. Creates AST nodes for each block type (headings, paragraphs, lists, etc.).
4. Returns a root `div` node containing the full AST.

### blocks

Handles block-level markdown parsing.

#### `func MarkdownToBlocks(input string) []string`

Splits a Markdown string into a slice of block-level strings. It correctly separates different block elements like paragraphs, lists, and code blocks based on blank lines.

#### `func BlockToBlockType(block string) BlockType`

Determines the `BlockType` (e.g., `Paragraph`, `Heading`, `CodeBlock`) of a given block string. This is used to decide which parsing logic to apply to each block.

### inline

Handles inline markdown parsing.

#### `func TextToTextNodes(s string) []nodes.TextNode`

Converts a string of inline markdown into a slice of `TextNode` structs. It first tokenizes the input, then parses the token stream to build the AST.

#### `func ParseInlineStack(tokens []tokenizer.Token) []nodes.TextNode`

A sophisticated single-pass parser that processes a stream of tokens to handle nested formatting. It uses a delimiter stack to correctly match opening and closing markers for elements like bold, italic, and strikethrough.

### tokenizer

Generates tokens for inline parsing.

#### `func TokenizeInline(input string) []Token`

Converts a raw string into a slice of `Token`s. It identifies different inline elements such as emphasis markers, links, images, code spans, and raw text.

#### `func parseEntity(runes []rune, i int) (string, int)`

Handles HTML entities and character references, including named, decimal, and hexadecimal forms.

### nodes

Defines AST nodes and includes methods for HTML conversion.

#### Types

-   `type enum int`: An enumeration for different inline text types (`Bold`, `Italic`, `Link`, etc.).
-   `type TextNode`: The primary struct for the AST. It can represent a simple text element, a formatted element with children, or an HTML tag with properties.
-   `type HTMLNode`: An interface with `ToHTML()` and `PropsToHTML()` methods (though `TextNode` implements this directly).

#### `func (h *TextNode) ToHTML() string`

A method on `TextNode` that recursively converts an AST node and its children into an HTML string. It handles HTML escaping and correctly formats different tags.

## SSGo Markdown Specification

### Block Elements

#### Thematic Breaks

A line consisting of 3 or more `*`, `-`, or `_` characters, optionally separated by spaces.

**Markdown:**
```markdown
---
* * *
```
**HTML:**
```html
<hr />
```

#### ATX Headings

A line prefixed with 1-6 `#` characters. If more than 6 `#` characters are used, it will be rendered as an `<h6>`.

**Markdown:**
```markdown
# Heading 1
## Heading 2
```
**HTML:**
```html
<h1>Heading 1</h1>
<h2>Heading 2</h2>
```

#### Fenced Code Blocks

A block of code surrounded by lines with 3 or more backticks (`` ` ``) or tildes (`~`). Language-specific syntax highlighting is not currently supported.

**Markdown:**
````markdown
```
func main() {
    fmt.Println("Hello, World!")
}
```
````
**HTML:**
```html
<pre><code>func main() {
    fmt.Println("Hello, World!")
}
</code></pre>
```

#### Blockquotes

Lines prefixed with `> `.

**Markdown:**
```markdown
> This is a blockquote.
```
**HTML:**
```html
<blockquote>
  <p>This is a blockquote.</p>
</blockquote>
```

#### Lists

- **Unordered Lists**: List items prefixed with `-` or `*`.
- **Ordered Lists**: List items prefixed with a number followed by a period (`.`).

**Markdown:**
```markdown
- Apple
- Banana

1. First
2. Second
```
**HTML:**
```html
<ul>
  <li>Apple</li>
  <li>Banana</li>
</ul>
<ol>
  <li>First</li>
  <li>Second</li>
</ol>
```

#### Raw HTML Blocks
HTML blocks are passed through directly without modification.

**Markdown:**
```markdown
<div class="note">
  <p>This is a raw HTML block.</p>
</div>
```
**HTML:**
```html
<div class="note">
  <p>This is a raw HTML block.</p>
</div>
```

### Inline Elements

#### Emphasis and Strong Emphasis

- **Italic**: `*italic*` or `_italic_` becomes `<em>`.
- **Bold**: `**bold**` or `__bold__` becomes `<strong>`.
- **Bold and Italic**: `***text***` or `___text___` becomes `<strong><em>`.

#### Strikethrough

`~~strikethrough~~` becomes `<s>`.

#### Highlight

`==highlight==` becomes `<mark>`.

#### Subscript and Superscript

- **Subscript**: `~subscript~` becomes `<sub>`.
- **Superscript**: `^superscript^` becomes `<sup>`.

#### Code Spans

Inline code is wrapped in single backticks: `` `code` `` becomes `<code>`.

#### Links and Images

- **Links**: `[text](url "title")` becomes `<a href="url" title="title">text</a>`.
- **Images**: `![alt text](src "title")` becomes `<img src="src" alt="alt text" title="title" />`.

#### Autolinks

URLs and email addresses are automatically converted into links.
- This applies to both plaintext links and those enclosed in angle brackets (`<>`).
- If a protocol (like `http://` or `https://`) is missing, `https://` is automatically added.
- For authenticated email links (`<user:pass@example.com>`), the password is removed for security, and the `href` becomes `mailto:user@example.com`.

#### Raw HTML

Inline HTML tags are passed through directly.

**Markdown:**
```markdown
An <span style="color:red;">HTML</span> span.
```
**HTML:**
```html
<p>An <span style="color:red;">HTML</span> span.</p>
```

#### HTML Entities

Named (`&copy;`), decimal (`&#169;`), and hex (`&#xA9;`) entities are decoded.

#### Hard Line Breaks

A backslash (`\`) or two spaces at the end of a line creates a `<br />`.

## Unsupported Features

The following features are not currently supported:

-   **Nested Lists**: Lists inside other list items.
-   **Nested Blockquotes**: Blockquotes inside other blockquotes.
-   **Mixed Block Elements**: Block-level elements (like paragraphs or code blocks) inside list items or blockquotes.
-   **Reference-Style Links and Images**: e.g., `[text][label]`.
-   **Task Lists**: e.g., `- [x] Done`.
-   **Definition Lists**.
-   **Tables**.
-   **Indented Code Blocks**.

## Contributing

This project was developed as a personal learning exercise and a portfolio piece. While I appreciate the community's interest, I am not seeking external contributions at this time. This allows me to maintain a clear vision for the project and focus on my own development goals.

If you have suggestions or find bugs, feel free to open an issue. I will review them as time permits. Thank you for your understanding.