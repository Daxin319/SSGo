# SSGo, a Static Site Generator written in Go

Currently offering basic markdown support, this can take markdown as input and will output HTML formatted to be either served to localhost:3000 for testing or to be uploaded to a domain host like gitpages.

I decided to start this project because I decided to change careers and get into programming. I had written a version of this in Python3 earlier in my learning, so I decided to rewrite it in Go without using a guide to see how far I've come and to provide me with a way to show off my skill I've learned. What better way to show you know what you're doing than to make your [professional blog/information page](https://daxin319.github.io/SSGo/) something that was generated from a program that you wrote yourself?

It's rather basic at the moment, but I'm working on adding full commonmark support at this time. If you wish to use it yourself, it's quite simple, just check the Quick Start Section Below.

Big shout out to [boot.dev](https://www.boot.dev?bannerlord=daxin319) for helping me learn to program, and for not being mad that I stole your markdown files to test the basic functions before I started expanding features.

## Quick Start/Usage Guide

This program is designed to copy static resources into the `/docs` directory inside the repository from the `/static` directory, then process any markdown files in the `/content` directory into HTML which is then also put in the `/docs` directory to be served to localhost or formatted for your host service of choice.

To get started:

1. Fork this repository, and clone your fork to your own machine.
2. Put any static resources that will always be present into the static directory, you can organize them into dirs for website navigation as well.
3. In the `content` directory, place all of the `index.md` files that you would like to convert to html in the same file structure as your website.
4. If you're testing offline, run `. serve.sh` from the root of the repository to generate all pages into the `/docs` directory and serve to `localhost:3000`.
5. If you're uploading to git pages or something similar, run `. build.sh` from the root of the repository to build the html files into the `/docs` directory without serving to `localhost:3000`. **You will need to edit the file and change my directory `/SSGo` to your own repository/directory name to make it build the correct HTML for your site.**
6. ???
7. Profit.

### Contributing

This has been a solo project, but it wouldn't have been possible without the efforts of the community at [boot.dev](https://www.boot.dev?bannerlord=daxin319) teaching me how to program from scratch, and supporting me in my journey to break into the tech scene.
