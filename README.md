# furit ( <u>F</u>ind <u>U</u>n<u>r</u>eferenced <u>I</u>mages in <u>T</u>ext files )

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT) [![test-and-build](https://github.com/kemokemo/furit/actions/workflows/test-and-build.yml/badge.svg)](https://github.com/kemokemo/furit/actions/workflows/test-and-build.yml)

This tool finds unreferenced images from text files such as markdown.

## Install

### Homebrew

```sh
brew install kemokemo/tap/furit
```

### Scoop

First, add my scoop-bucket.

```sh
scoop bucket add kemokemo-bucket https://github.com/kemokemo/scoop-bucket.git
```

Next, install this app by running the following.

```sh
scoop install furit
```

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

-s, -settings:
    specify the settings file path to exclude files etc..

-t, -type:
    specify the target text format (markdown, html are available)

-h, -help:
    display help

-v, -version:
    display version
```

### Example

```sh
$ furit content
content/posts/assets/some_screen.bmp
content/posts/assets/logo.gif

$ echo $?
1
```

This tool looks recursively for the folder you specify, finds links to images in the text it finds, and enumerates the unreferenced image files from text.

If any unreferenced images were found, it returns `1` as an `ExitCode`. This is also true for the deletion operation described below.

```sh
$ furit -d content
content/posts/assets/some_screen.bmp
content/posts/assets/logo.gif
Are you sure to delete these unlinked images? [y/n]: n
the file deletion process has been canceled by user input
```

If you want to delete any unreferenced images found by this tool while checking them, just specify the `-d` flag.

```sh
$ furit -d -f content
content/posts/assets/some_screen.bmp
content/posts/assets/logo.gif
```

You can also specify the `-f` flag if you want to run the process of deletion automatically without confirmation. In that case, the list of files to be deleted will still be printed.

If the deletion fails, it returns an `ExitCode` other than `0` and `1`.

## Supported

### Text format

The following text are supported.

- Markdown
- HTML

### Image format

The following image extensions are supported. They are not case insensitive.

- png
- jpg, jpeg
- bmp
- gif
- tif, tiff
- emf

## Notes

If you use `img` tag in markdown files to specify the image size, please convert to the html files and use `html` type option.

## License

[MIT](https://github.com/kemokemo/furit/blob/main/LICENSE)

## Author

[kemokemo](https://github.com/kemokemo)

