# "GoTTA Htmx" Stack

This started off simple Go & HTMX stack but it has expanded as I wanted to try out a few other technologies.
It follows the known "GoTTH" stack but also adds Alpine.js to become "GoTTA H", you know... you've "gotta htmx".

```plain
Go - Golang powered back-end
T  - Templ
T  - TailwindCSS
A  - Alpine.js

Htmx
```

It also utilizes [Air](https://github.com/air-verse/air?tab=readme-ov-file) for hot reloading during development. Out of the box Air works with go but needed to be adjusted to work with the other parts of the stack and to enable automatic browser reload on file changes.

To automate the multiple build processes, e.g. css bundling, js buncling, alongside templ code generation and Go compilation use leverage [templ's built-in proxy server](https://templ.guide/developer-tools/live-reload-with-other-tools/#templ-watch-mode). We do this by setting up a `Makefile` to run all of the necessary commands.

To start simply use `make run`.
