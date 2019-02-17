# img2gray

Package img2gray lets you generate gray images.

## Installation

```bash
$ go get github.com/po3rin/img2gray/cmd/gogray
```

If you do not install go, binary file is here!
https://github.com/po3rin/img2gray/releases

## Usage

as CLI tool.

```bash
# required i flag
$ img2gray -i testimg/gopher.jpeg
```

When you traverse image files to generate gray images, use -t flag.

```
$ img2gray -i testimg -t
generate grayscale image file testimg/dir/gray_gopher.jpg
generate grayscale image file testimg/gray_gopher.jpeg
```
