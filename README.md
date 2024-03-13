# OpenCanvas-go

Golang library for parsing the [JSON Canvas 1.0](https://jsoncanvas.org) file
format and rendering it to HTML.

**Work in progress:** the project is in a very early stage right now and doesn't
*do much.

## Testing

Tests rely on sample data from the
[JSON Canvas GitHub repository](https://github.com/obsidianmd/jsoncanvas).

Download submodules to obtain it:
```bash
git submodule update --init --recursive
```

## Licensing

All OpenCanvas-go code is licensed under the [MIT License](LICENSE).

Parts of the code are based on the
[JSON Canvas 1.0 specification](https://jsoncanvas.org/spec/1.0.html)
(mainly I've taken the docstrings mostly as is from the spec). That is also MIT,
[copyright Obsidian.md](LICENSE.jsoncanvas).
