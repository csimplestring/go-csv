# go-csv

This tiny tool now provides the following functionalities:
- Automatically detects the CSV delimiter. It fully conforms to the specifications outlined on the on the [Wikipedia article][csv].

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
