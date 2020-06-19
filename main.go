// tail | (c) 2020 NETWAYS GmbH | GPLv2+

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	linesParameter := flag.Int("n", 10, "specify how many last lines should be printed.")
	flag.Parse()
	buf := bufio.NewReader(os.Stdin)
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
