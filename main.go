// tail | (c) 2020 NETWAYS GmbH | GPLv2+

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var linesParameter = flag.Int("n", 10, "specify how many last lines should be printed.")

func main() {
	flag.Parse()

	if len(flag.Args()) > 0 {
		for filePOS, fileName := range flag.Args() {
			file, fileErr := os.Open(fileName)
			if fileErr != nil {
				fmt.Fprintln(os.Stderr, fileErr)
				os.Exit(1)
			}

			printTailedFile(file)

			if filePOS != len(flag.Args())-1 {
				fmt.Print("\n")
			}

			file.Close()
		}
	} else {
		printTailedFile(os.Stdin)
	}
}

func printTailedFile(file *os.File) {
	if len(flag.Args()) > 1 {
		fmt.Printf("==> %s <==\n", file.Name())
	}

	buf := bufio.NewReader(file)
	var allData [][]byte

	for {
		data, dataErr := buf.ReadBytes('\n')
		if dataErr != nil && dataErr != io.EOF {
			fmt.Fprintln(os.Stderr, dataErr)
			os.Exit(1)
		}
		allData = append(allData, data)

		if len(allData) > *linesParameter && len(data) != 0 {
			allData = allData[1:]
		}

		if dataErr == io.EOF {
			break
		}
	}

	if len(allData[len(allData)-1]) == 0 {
		allData = allData[:len(allData)-1]
	}

	for _, dataName := range allData {
		_, _ = os.Stdout.Write(dataName)
	}
}
