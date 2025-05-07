# SSGo

After learning Go and falling in love with it, I decided to rewrite a project I did earlier in my education, a Static Site Generator.

I am using this program to generate my professional webpage [link here when finished]

It's rather basic at the moment, but I'm working on adding full commonmark support at this time. If you wish to use it yourself, it's quite simple.

Big shout out to [boot.dev](https://www.boot.dev/) for helping me learn to program, and for not being mad that I stole your markdown files to test this.

## Directions:

1. put any static resources that will always be present into the static directory, you can organize them into dirs for website navigation as well
2. in the `content` directory, place all of the index.md files that you would like to convert to html
3. if you're testing offline, run `. serve.sh` to serve to localhost:3000
4. if you're uploading to git pages or something similar, run `. build.sh` to build the html files. You will need to edit the file and change my directory `/SSGo` to your own repository/directory name to make it build the correct HTML for your site.
5. ???
6. Profit.
