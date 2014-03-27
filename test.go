package main

import (
	"ctohtml"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func filenamesFromCommandLine() (inFilename, outFilename string, err error) {
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		err = fmt.Errorf("usage: %s [<]infile.zcd [>]outfile.html", filepath.Base(os.Args[0]))
		return "", "", err
	}
	if len(os.Args) > 1 {
		inFilename = os.Args[1]
		if len(os.Args) > 2 {
			outFilename = os.Args[2]
		}
	}
	if inFilename != "" && inFilename == outFilename {
		log.Fatal("won't overwrite the infile")
	}
	return inFilename, outFilename, nil
}

func main() {
	inFilename, outFilename, err := filenamesFromCommandLine()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inFile, outFile := os.Stdin, os.Stdout
	if inFile, err = os.Open(inFilename); err != nil {
		log.Fatal(err)
	}
	defer inFile.Close()
	if outFile, err = os.Create(outFilename); err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	ctohtml.FileToHtml(inFile, outFile)
}
