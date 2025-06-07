# SSGo: A Modern Static Site Generator in Go

SSGo is a fast, extensible static site generator written in Go. It converts Markdown files into browser readable HTML, supporting a wide range of CommonMark and extended Markdown features.

For a detailed breakdown of the architecture, supported Markdown syntax, and more, please see the full [documentation](documentation.md).

## Features

- **Markdown Support**: Handles a broad range of block elements (headings, lists, code blocks) and inline elements (bold, italic, links, images, and more). See the [Markdown Support](#markdown-support) section in the documentation for a complete list.
- **Extended Syntax**: Includes support for strikethrough, highlighting, subscript, and superscript.
- **Autolinking**: Automatically detects and links URLs and email addresses.
- **Raw HTML**: Allows embedding raw HTML directly in Markdown files.
- **Customizable**: Supports custom HTML templates and static asset handling.

## Quick Start

### Prerequisites
- Go 1.22 or later
  - To install, I recommend the [Webi installer](https://webi.sh/golang):
    ```sh
    curl -sS https://webi.sh/golang | sh
    ```
  - Alternatively, see the [official Go installation instructions](https://go.dev/doc/install).
- Git (I refuse to believe you don't already have git installed if you're reading this, but just in case: [Git download](https://git-scm.com/downloads) or on Linux: `sudo pacman -S git` or `sudo apt install git`)

### Installation
1. Clone this repository:
   ```sh
   git clone https://github.com/Daxin319/SSGo.git
   cd SSGo
   ```

### Usage
1.  Place your Markdown content in the `/content` directory.
2.  Add any static assets (CSS, images) to the `/static` directory.
3.  Run the development server:
    ```sh
    ./serve.sh
    ```
    Your site will be available at `http://localhost:3000`.

4.  Build for production:
    ```sh
    ./build.sh
    ```
    The output will be generated in the `/docs` directory, ready for deployment.

## About This Project

This project began as a personal journey into programming and software development. Having previously implemented a similar static site generator in Python3, I chose to rewrite it in Go as a way to demonstrate my growth as a developer and showcase my current skills. You can see the output of my test file [here](https://daxin319.github.io/SSGo/) (Feel free to compare it with the `index.md` file in the `content` directory to see the raw markdown). The output page also includes documentation about the differences between the Commonmark standard and my quality-of-life enhancements, along with details about supported extended markdown features.

Special thanks to [boot.dev](https://www.boot.dev?bannerlord=daxin319) for their excellent programming education platform and for graciously allowing me to use their markdown files for initial testing before expanding the feature set.

This project is not currently accepting contributions. For more details, please see the "Contributing" section in the [documentation](documentation.md).

## License

This project is licensed under the MIT License with the Commons Clause restriction. See the [LICENSE](LICENSE) file for details.

### Commercial Use

Commercial use of this software is not permitted under the current license. If you are interested in a commercial license, please contact Lyle Higginbotham via GitHub or at lylehigg@gmail.com.