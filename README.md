# SSGo: A Modern Static Site Generator in Go

SSGo is a fast, extensible static site generator written in Go. It converts Markdown files into well-structured HTML, supporting a wide range of CommonMark and extended Markdown features, including GFM-style autolinks, robust autolink validation, raw HTML passthrough, and advanced line break handling.

## Features

- **CommonMark and Extended Markdown Support**
  - Headers, paragraphs, lists, blockquotes, code blocks, horizontal rules
  - Inline formatting: bold, italic, strikethrough, subscript, superscript, highlight
  - Links, images, code spans, HTML entities
  - GFM-style autolinks: bare domains, subdomains, ports, paths, authentication, and email addresses
  - Autolink validation using the publicsuffix library for real TLDs
  - Email autolinks support optional authentication (user:pass), with passwords stripped from the rendered link
  - Hard line breaks via two spaces or backslash at end of line
  - Raw HTML blocks and comments are preserved and rendered as HTML

- **Modern Architecture**
  - Modular codebase for maintainability and extensibility
  - Clean separation of file I/O, block parsing, inline parsing, AST, and HTML rendering

- **Flexible Output**
  - Generates HTML files ready for deployment or local development
  - Customizable templates and static asset support

## Quick Start

### Prerequisites
- Go 1.22 or later
- Git (I refuse to believe you don't already have git installed if you're reading this, but just in case: [Git download](https://git-scm.com/downloads) or on Linux: `sudo pacman -S git` or `sudo apt install git`)

### Installation
1. Install Go 1.22+ (recommended: [Webi installer](https://webi.sh/golang)):
   ```sh
   curl -sS https://webi.sh/golang | sh
   ```
   Alternatively, see the [official Go installation instructions](https://go.dev/doc/install).
2. Clone this repository:
   ```sh
   git clone https://github.com/Daxin319/SSGo.git
   cd SSGo
   ```

### Usage
1. Place static assets in the `/static` directory.
2. Add your Markdown content to the `/content` directory, mirroring your desired site structure.
3. To develop locally:
   ```sh
   ./serve.sh
   ```
   Visit [http://localhost:3000](http://localhost:3000) to view your site.
4. To build for production:
   ```sh
   ./build.sh
   ```
   Deploy the contents of `/docs` to your hosting provider.
5. If deploying to GitHub Pages or a custom path, edit `build.sh` to set the correct base path for your site.
6. Enjoy your new static site!
7. ???
8. Profit.

## About This Project

This project began as a personal journey into programming and software development. Having previously implemented a similar static site generator in Python3, I chose to rewrite it in Go as a way to demonstrate my growth as a developer and showcase my current skills. You can see the output of my test file [here](https://daxin319.github.io/SSGo/) (Feel free to compare it with the `index.md` file in the `content` directory to see the raw markdown). The output page also includes documentation about the differences between the Commonmark standard and my quality-of-life enhancements, along with details about supported extended markdown features.

Special thanks to [boot.dev](https://www.boot.dev?bannerlord=daxin319) for their excellent programming education platform and for graciously allowing me to use their markdown files for initial testing before expanding the feature set.

**Please note:** I am not currently accepting direct contributions or pull requests. For more information, see the "Contributing" section in [documentation.md](documentation.md).

## License

This project is licensed under the MIT License with the Commons Clause restriction. See the [LICENSE](LICENSE) file for details.

## Commercial Use

Commercial use of this software is not permitted under the current license (MIT + Commons Clause).  
If you are interested in obtaining a commercial license, please contact Samuel Lyle Higginbotham (Daxin319) via GitHub or email: lylehigg@gmail.com.