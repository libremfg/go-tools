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
var verbose bool

func init() {
	flag.StringVar(&dir, "directory", "./path/to/files/", "directory of files to merge")
	flag.StringVar(&extensions, "extension", ".txt", "the file extension of files to merge. Use comma separation for multiple types.")
	flag.StringVar(&output, "output", "./output.txt", "the output file name")
	flag.BoolVar(&verbose, "verbose", false, "verbose logging")
}

func main() {

	fmt.Println(os.Args)

	if !verbose {
		fmt.Printf("called with directory=%s, extension=%s, output=%s\n", dir, extensions, output)
	}

	acceptedExtensions := strings.Split(extensions, ",")

	if verbose {
		fmt.Printf("accepted extensions: include %s\n", strings.Join(acceptedExtensions, ", "))
	}

	Merge(dir, acceptedExtensions, output, verbose)
}

func Merge(d string, fileExtensions []string, outputFilepath string, verboseLogging bool) {
	directory := os.DirFS(d)

	tmp := make([]byte, 0)

	walkErr := fs.WalkDir(directory, ".", func(path string, entry fs.DirEntry, err error) error {
		if entry.IsDir() {
			if verboseLogging {
				fmt.Printf("skipping directory %s\n", d+path)
			}
			return nil
		}

		if err != nil {
			if verboseLogging {
				fmt.Printf("error checking directory %s\n", d+path)
			}
			return err
		}

		if len(fileExtensions) != 0 {
			found := false
			for _, extension := range fileExtensions {
				if strings.HasSuffix(path, extension) {
					found = true
					fmt.Printf("found matching file extension %s for %s\n", extension, d+path)
					break
				}
			}
			if !found {
				if verboseLogging {
					fmt.Printf("didn't find matching file extension for %s\n", d+path)
				}
				return nil
			}
		} else {
			if verboseLogging {
				fmt.Println("no extensions to check")
			}
		}

		fileBytes, readErr := os.ReadFile(d + path)
		if readErr != nil {
			fmt.Printf("error reading file %s\n", d+path)
			return readErr
		}

		tmp = append(tmp, fileBytes...)
		return nil
	})

	if walkErr != nil {
		fmt.Printf("error walking file system %s\n", walkErr)
	}

	if _, err := os.Stat(outputFilepath); err == nil {

		inputBytes, readErr := os.ReadFile(outputFilepath)
		if readErr != nil {
			fmt.Printf("didn't find existing file %s\n", outputFilepath)
			log.Print(err)
			os.Exit(1)
		}

		if err := os.WriteFile(outputFilepath+".old", inputBytes, 0777); err != nil {
			log.Print(err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("didn't find existing file %s\n", outputFilepath)
	}

	err := os.WriteFile(outputFilepath, tmp, 0777)

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	fmt.Printf("created %s\n", outputFilepath)
}
