package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// ** original code: https://gist.github.com/r0l1/3dcbb0c8f6cfe9c66ab8008f55f8f28b
// askForConfirmation returns whether or not the user agrees to the process.
func askForConfirmation(s string, in io.Reader, out io.Writer, retry int) (bool, error) {
	reader := bufio.NewReader(in)

	for ; retry > 0; retry-- {
		fmt.Fprintf(out, "%s [y/n]: ", s)

		res, err := reader.ReadString('\n')
		if err != nil {
			return false, fmt.Errorf("failed to read user input: %s", err)
		}

		res = strings.ToLower(strings.TrimSpace(res))
		if res == "y" || res == "yes" {
			return true, nil
		} else if res == "n" || res == "no" {
			return false, nil
		}
	}

	return false, fmt.Errorf("retry count has been reached")
}
