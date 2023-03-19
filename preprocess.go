package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/nfnt/resize"
	"image/jpeg"
)

func PreprocessImage(inputDir string, outputDir string) {
	// Define the input and output directories.


	// Create the output directory if it doesn't exist.
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.Mkdir(outputDir, os.ModePerm)
	}

	// Get the list of files in the input directory.
	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		log.Fatal(err)
	}

	count := 0

	// Iterate over each file in the input directory.
	for _, file := range files {
		// Open the input file.
		inputPath := filepath.Join(inputDir, file.Name())
		inputFile, err := os.Open(inputPath)
		if err != nil {
			log.Printf("Error opening file %s: %s", inputPath, err)
			continue
		}
		defer inputFile.Close()

		// Decode the input file.
		img, err := jpeg.Decode(inputFile)
		if err != nil {
			log.Printf("Error decoding file %s: %s", inputPath, err)
			continue
		}

		// Resize the image to 100x100 pixels.
		resizedImg := resize.Resize(100, 100, img, resize.Lanczos3)

		// Save the resized image to the output directory.
		outputPath := filepath.Join(outputDir, file.Name())
		outputFile, err := os.Create(outputPath)
		if err != nil {
			log.Printf("Error creating file %s: %s", outputPath, err)
			continue
		}
		defer outputFile.Close()

		count ++
		if count > 100{
			break
		}

		if err := jpeg.Encode(outputFile, resizedImg, nil); err != nil {
			log.Printf("Error encoding file %s: %s", outputPath, err)
			continue
		}

		// Print a success message.
		fmt.Printf("Resized and saved file %s\n", outputPath)
	}
}