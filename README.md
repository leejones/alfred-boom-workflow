# Alfred Boom Workflow

_Reducing your Mean Time To Gif_

Search your boom text snippets from Alfred and copy the selected snippet to the clipboard.

## Requirements

* [Alfred v2](http://www.alfredapp.com/)
* [boom](https://github.com/holman/boom)

## Installation

* download the [latest release](https://github.com/leejones/alfred-boom-workflow/releases)
* open the downloaded file and Alfred will import it automatically

## Usage

From Alfred:

```
boom [list name] [snippet name]
```

Lists and snippets will be filtered as you type.

## Development

The script filter is written in [Go](http://golang.org) so you'll need to have that installed before starting development.

### Running the tests

```
bin/test
```

### Building the binary

```
bin/build
```

### Packaging and installing into Alfred

```
bin/install
```

## Release Process

* push the code changes to GitHub
* build a clean package (`bin/install`)
* create a [new release on GitHub](https://github.com/leejones/alfred-boom-workflow/releases/new)
* attach the freshly built package (`boom.alfredworkflow`) to the release
* note the change(s) in the release description
* publish the release
* profit
