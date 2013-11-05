package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var zipFile = flag.Bool("z", false, "Set to true to zip file as it's being written.")

func main() {
	customFileName := false
        dateTpl := ""
	flag.Parse()

	// get the file name template
	fn := flag.Arg(0)

	// if there is not a filename, print the usage and exit
	if fn == "" {
		printUsage()
		os.Exit(0)
	}

	// check and see if a date template has been included
	fileParts := strings.Split(fn, "%")
	if len(fileParts) > 2 {
		// dateTpl := fileParts[1]
		customFileName = true
		// if there is a date template set the currentTime variable
                dateTpl := fileParts[1]
                fileParts[1] = time.Now().Format(dateTpl)
		fn = strings.Join(fileParts, "")
	}
	fd, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Printf("Error: %q\n", err)
		os.Exit(1)
	}
	defer fd.Close()

	// start the loop
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadBytes('\n')
                if customFileName {
                    if time.Now().Format(dateTpl) != fileParts[1] {
                        fileParts[1] = time.Now().Format(dateTpl)
                        fn = strings.Join(fileParts, "")
                        fd.Close()
                        fd, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
                        if err != nil {
                                fmt.Printf("Error: %q\n", err)
                                os.Exit(1)
                        }
                        defer fd.Close()
                    }
                }

		if err != nil {
			break
		}

		fd.Write(line)
	}
}

func printUsage() {
	fmt.Printf(`Usage: godl [OPTIONS] FILENAME

Writes content from stdin to FILENAME.

FILENAME can include a date template using the date
Mon Jan 2 15:04:05 -0700 MST 2006 as the layout surrounded
by %%. For example
    
godl http_%%20060101%%.log

OPTIONS
    -z flag sets the file to be compressed as it's written.
`)
}
