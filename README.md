# img2gray

Package img2gray lets you generate gray images.

## Installation

```bash
$ go get github.com/po3rin/img2gray/cmd/gogray
```

## Usage

as CLI tool.

```bash
$ img2gray -i testimg/gopher.jpeg
```

When you traverse image files to generate gray images, use -t flag.

```
$ img2gray -i testimg -t
detect image file testimg/dir/gopher.jpg
detect image file testimg/gopher.jpeg
```
