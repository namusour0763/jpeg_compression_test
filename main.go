package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <input_image.jpg> <number_of_iterations>")
		return
	}

	inputFile := os.Args[1]
	iterations, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid number of iterations:", err)
		return
	}

	for i := 0; i < iterations; i++ {
		outputFile := fmt.Sprintf("output_%d.jpg", i+1)
		err := compressAndSave(inputFile, outputFile)
		if err != nil {
			fmt.Printf("Error processing iteration %d: %v\n", i+1, err)
			return
		}
		inputFile = outputFile
	}

	fmt.Printf("Compression completed. Final output: %s\n", inputFile)
}

func compressAndSave(inputPath, outputPath string) error {
	// Open the input file
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("error opening input file: %v", err)
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("error decoding image: %v", err)
	}

	// Create the output file
	out, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer out.Close()

	// Encode and save the image as JPEG
	options := jpeg.Options{Quality: 75} // You can adjust the quality (0-100)
	err = jpeg.Encode(out, img, &options)
	if err != nil {
		return fmt.Errorf("error encoding image: %v", err)
	}

	return nil
}
