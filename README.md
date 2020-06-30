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

-d, -delete:
    delete unlinked image files (with confirmation)

-f, -force:
    delete unlinked image files without prompting for confirmation

-h, -help:
    display help

-v, -version:
    display version
```

Currently, only Markdown format is supported as text to find links to images.

### Example

```sh
$ furit content
content/posts/assets/some_screen.bmp
content/posts/assets/logo.gif
```

This tool looks recursively for the folder you specify, finds links to images in the text it finds, and enumerates the unlinked image files from text.

```sh
$ furit -d content
content/posts/assets/some_screen.bmp
content/posts/assets/logo.gif
Are you sure to delete these unlinked images? [y/n]: n
the file deletion process has been canceled by user input
```

If you want to delete unwanted images found by the tool while reviewing them, specify only the `-d` flag.

```sh
$ furit -d -f content
content/posts/assets/some_screen.bmp
content/posts/assets/logo.gif
```

You can also specify the `-f` flag if you want to run the process of deletion automatically without confirmation. In that case, the list of files to be deleted will still be printed.
## License

[MIT](https://github.com/kemokemo/furit/blob/master/LICENSE)

## Author

[kemokemo](https://github.com/kemokemo)

