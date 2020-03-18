package main

import (
	"fmt"
	"strings"
)

func lineProc(line, old, new string) (found bool, newline string, count int) {

	newline = line
	if strings.Contains(line, old) {
		found = true
		count += strings.Count(line, old)
		newline = strings.Replace(line, old, new, -1)
		return found, newline, count
	}
	return found, newline, count
}

func main() {

	found, newline, count := lineProc("this is just a Test test for the tEsting part teSt", "test", "polo") //
	fmt.Println(found, newline, count)
}
