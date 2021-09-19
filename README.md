# TIFF Converter


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
