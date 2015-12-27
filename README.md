# go-csv

This tiny tool is dedicated to automatically detects the CSV delimiter. It fully conforms to the specifications outlined on the on the [Wikipedia article][csv]. Looking through many CSV libraries code and discussion on the stackoverflow, finding that their CSV delimiter detection is limited or incomplete or containing many unneeded features. Hoping this can people solve the CSV delimiter detection problem without importing extra overhead.

[csv]: http://en.wikipedia.org/wiki/Comma-separated_values

## Usage

    package main
    
    import (
    	"github.com/csimplestring/go-csv/detector"
    	"os"
    	"fmt"
    )
    
    func main()  {
    	detector := detector.New()
    
    	file, err := os.OpenFile("example.csv", os.O_RDONLY, os.ModePerm)
    	if err != nil {
    		os.Exit(1)
    	}
    	defer file.Close()
    
    	delimiters := detector.DetectDelimiter(file, '"')
    	fmt.Println(delimiters)
    }

This tool is inspired by [parseCSV][link].
[link]: https://github.com/parsecsv/parsecsv-for-php
