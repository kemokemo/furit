# furit ( <u>F</u>ind <u>U</u>n<u>r</u>eferenced <u>I</u>mages in <u>T</u>ext files )

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT) [![CircleCI](https://circleci.com/gh/kemokemo/furit.svg?style=svg)](https://circleci.com/gh/kemokemo/furit) [![codecov](https://codecov.io/gh/kemokemo/furit/branch/master/graph/badge.svg)](https://codecov.io/gh/kemokemo/furit) [![Go Report Card](https://goreportcard.com/badge/github.com/kemokemo/furit)](https://goreportcard.com/report/github.com/kemokemo/furit)

This tool finds unreferenced image assets from text files such as markdown.

## Install

### Binary

Get the latest version from [the release page](https://github.com/kemokemo/furit/releases/latest), and download the archive file for your operating system/architecture. Unpack the archive, and put the binary somewhere in your `$PATH`.

## Usage

```sh
$ furit -h
Usage: furit [<option>...] <1st path> <2nd path>...
 you can set mutiple paths to search invalid images.

  -h	display help
```

Currently, only Markdown format is supported as text to find links to images.

### Example

```
$ furit lib/test-data/markdown
lib/test-data/markdown/assets/画面.bmp
lib/test-data/markdown/テスト.gif
```

This tool looks recursively for the folder you specify, finds links to images in the text it finds, and enumerates the unlinked image files from text.

## License

[MIT](https://github.com/kemokemo/furit/blob/master/LICENSE)

## Author

[kemokemo](https://github.com/kemokemo)

