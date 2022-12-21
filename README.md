# TIFF Converter
[![Go Workflow Status](https://github.com/wintermi/tiff-convert/workflows/Go/badge.svg)](https://github.com/wintermi/tiff-convert/actions/workflows/go.yml)&nbsp;[![Go Report Card](https://goreportcard.com/badge/github.com/wintermi/tiff-convert)](https://goreportcard.com/report/github.com/wintermi/tiff-convert)&nbsp;[![license](https://img.shields.io/github/license/wintermi/tiff-convert.svg)](https://github.com/wintermi/tiff-convert/blob/main/LICENSE)&nbsp;[![GitHub release (latest by date including pre-releases)](https://img.shields.io/github/v/release/wintermi/tiff-convert?include_prereleases)](https://github.com/wintermi/tiff-convert/releases)


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
