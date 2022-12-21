// Copyright 2021, Matthew Winter
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"image/png"
	"io"
	"os"
	"path/filepath"

	"github.com/andviro/go-libtiff/libtiff"
)

var helpText = `%s 0.1.0
Copyright 2021, Matthew Winter

tiff-convert recursively walks the input path searching for all TIFF files and
exporting each page of these TIFF files to the output path as an image file
using the image encoder stated.

Use --help for more details.


USAGE:
    tiff-convert -e ENCODER -i PATH -o PATH

ARGS:
`

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, helpText, filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}

	// Define the Long CLI flag names
	var fileEncoder = flag.String("e", "png", "Page Encoder [png]  (Required)")
	var inputPath = flag.String("i", "", "Input Path  (Required)")
	var outputPath = flag.String("o", "", "Output Path  (Required)")

	// Parse the flags
	flag.Parse()

	// Validate the Required Flags
	if *inputPath == "" || *outputPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Validate the File Format Selection
	// fileFormatChoices := map[string]bool{"png": true, "jpeg": true, "tiff": true}
	fileFormatChoices := map[string]bool{"png": true}
	if _, validChoice := fileFormatChoices[*fileEncoder]; !validChoice {
		flag.Usage()
		os.Exit(1)
	}

	// Traverse all files and directories under the Input Path searching for TIFF files
	fileCounter := 0
	err := filepath.Walk(*inputPath, func(filename string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fileInfo.IsDir() {
			fileCounter++
			fmt.Fprintln(os.Stdout, "----------------------------------------")
			fmt.Fprintf(os.Stdout, "File %d: %q\n", fileCounter, filename)
			fmt.Fprintf(os.Stdout, "File Size: %d bytes\n\n", fileInfo.Size())

			// Calculate the Relative Path and Pass
			relativePath, _ := filepath.Rel(*inputPath, filename)
			err = convertToPNG(os.Stdout, *inputPath, *outputPath, relativePath)
			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Fprintf(os.Stdout, "\n\n")
		}
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open directory, Error: %s\n", err)
		os.Exit(1)
	}
}

func convertToPNG(w io.Writer, inputPath string, outputPath string, relativePath string) error {
	srcFile, _ := filepath.Abs(filepath.Join(inputPath, relativePath))
	srcTIFF, err := libtiff.Open(srcFile)
	if err != nil {
		return err
	}
	defer srcTIFF.Close()

	srcTIFF.Iter(func(pageNumber int) {
		fmt.Fprintf(w, "Page %03d: ", pageNumber+1)

		img, err := srcTIFF.GetRGBA()
		if err != nil {
			fmt.Fprintln(w, "Error [srcTIFF.GetRGBA]:", err)
			return
		}

		relativeOutputFile := filepath.Join(relativePath, fmt.Sprintf("page-%03d.png", pageNumber+1))
		fmt.Fprintf(w, "%q\n", relativeOutputFile)

		dstFile, _ := filepath.Abs(filepath.Join(outputPath, relativeOutputFile))
		if _, err := os.Stat(dstFile); os.IsNotExist(err) {
			err = os.MkdirAll(filepath.Dir(dstFile), 0700)
			if err != nil {
				fmt.Fprintln(w, "Error [os.MkdirAll]:", err)
				return
			}
		}

		f, err := os.Create(dstFile)
		if err != nil {
			fmt.Fprintln(w, "Error [os.Create]:", err)
			return
		}

		if err := png.Encode(f, &img); err != nil {
			f.Close()
			fmt.Fprintln(w, "Error [png.Encode]:", err)
			return
		}

		if err := f.Close(); err != nil {
			fmt.Fprintln(w, "Error [f.Close]:", err)
			return
		}
	})

	return nil
}
