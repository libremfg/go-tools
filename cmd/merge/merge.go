package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

var dir string
var extensions string
var output string

func init() {
	flag.StringVar(&dir, "directory", "./path/to/files/", "directory of files to merge")
	flag.StringVar(&extensions, "extension", ".txt", "the file extension of files to merge. Use comma separation for multiple types.")
	flag.StringVar(&output, "output", "./output.txt", "the output file name")
}

func main() {
	tmp := make([]byte, 0)
	directory := os.DirFS(dir)

	acceptedExtensions := strings.Split(extensions, ",")

	walkErr := fs.WalkDir(directory, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		if err != nil {
			return err
		}

		if len(acceptedExtensions) != 0 {
			found := false
			for _, extension := range acceptedExtensions {
				if strings.HasSuffix(path, extension) {
					found = true
					break
				}
			}
			if !found {
				return nil
			}
		}

		fileBytes, readErr := os.ReadFile(dir + path)
		if readErr != nil {
			return readErr
		}

		tmp = append(tmp, fileBytes...)
		return nil
	})

	if walkErr != nil {
		log.Print(walkErr)
	}

	if _, err := os.Stat(output); err == nil {

		inputBytes, readErr := os.ReadFile(output)
		if readErr != nil {
			log.Print(err)
			os.Exit(1)
		}

		if err := os.WriteFile(output+".old", inputBytes, 0777); err != nil {
			log.Print(err)
			os.Exit(1)
		}
	}

	err := os.WriteFile(output, tmp, 0777)

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	fmt.Printf("created %s\n", output)
}
