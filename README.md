# TIFF Converter

[![Workflows](https://github.com/wintermi/tiff-convert/workflows/Go/badge.svg)](https://github.com/wintermi/tiff-convert/actions)
[![Go Report](https://goreportcard.com/badge/github.com/wintermi/tiff-convert)](https://goreportcard.com/report/github.com/wintermi/tiff-convert)
[![License](https://img.shields.io/github/license/wintermi/tiff-convert.svg)](https://github.com/wintermi/tiff-convert/blob/main/LICENSE)
[![Release](https://img.shields.io/github/v/release/wintermi/tiff-convert?include_prereleases)](https://github.com/wintermi/tiff-convert/releases)


## Description

"tiff-convert" recursively walks the input path searching for all TIFF files and exporting each page of these TIFF files to the output path as an image file using the image encoder stated.

```
USAGE:
    tiff-convert -e ENCODER -i PATH -o PATH

ARGS:
  -e string
    	Page Encoder [png]  (Required) (default "png")
  -i string
    	Input Path  (Required)
  -o string
    	Output Path  (Required)
```

## Known Limitations

The current version only exports the pages using the PNG Encoder.  Additional formats will be added in a future release.


## License

**tiff-convert** is released under the [Apache License 2.0](https://github.com/wintermi/tiff-convert/blob/main/LICENSE) unless explicitly mentioned in the file header.
