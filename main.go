// tail | (c) 2020 NETWAYS GmbH | GPLv2+

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var allData [][]byte

	for {
		data, dataErr := buf.ReadBytes('\n')
		if dataErr != nil && dataErr != io.EOF {
			fmt.Fprintln(os.Stderr, dataErr)
			os.Exit(1)
		}
		allData = append(allData, data)

		if len(allData) > 10 && len(data) != 0 {
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
