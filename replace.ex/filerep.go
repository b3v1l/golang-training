package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func replaceChar(line, old, new string) (found bool, newline string, count int) {

	newline = line
	if strings.Contains(line, old) {
		found = true
		count += strings.Count(line, old)
		newline = strings.Replace(line, old, new, -1)
		return found, newline, count
	}
	return found, newline, count
}

func findReplace(fsrc, old, new string) (count int, lines []int, err error) {

	srcFile, err := os.Open(fsrc)
	if err != nil {
		fmt.Printf("error : %v\n", err)
		//break
		//return count, lines, err
	}
	defer srcFile.Close()

	lineIdx := 1
	scanner := bufio.NewScanner(srcFile)
	for scanner.Scan() {
		found, newline, o := replaceChar(scanner.Text(), old, new)
		if found {
			count += o
			lines = append(lines, lineIdx)

		}
		fmt.Println(newline)
		lineIdx++
	}
	return count, lines, nil
}

func main() {
	old := "Go"
	new := "python"
	occ, lines, err := findReplace("tedst.txt", old, new)
	if err != nil {
		fmt.Println("Error in %v\n", err)
	}

	fmt.Println("== Summary ==")
	defer fmt.Println("== End of process ==")
	fmt.Printf("Number of occurence for the word %v: %v\n", old, occ)
	fmt.Printf("Number of lines: %v\n", len(lines))
}
