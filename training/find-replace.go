package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func replaceChar(line, old, new string) (res string, found bool, occ int) {

	res = line
	lowerOld := strings.ToLower(old)
	lowerNew := strings.ToLower(new)

	if strings.Contains(res, old) || strings.Contains(res, lowerOld) {
		found = true
		occ = strings.Count(res, old)
		occ = strings.Count(res, lowerOld)

		res = strings.Replace(res, old, new, -1)
		res = strings.Replace(res, lowerOld, lowerNew, -1)
		return res, found, occ
	}
	return res, found, occ
}

func openFile(src, dst, old, new string) (count int, lines []int, err error) {

	srcFile, err := os.Open(src)
	if err != nil {
		return count, lines, err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)

	if err != nil {
		return count, lines, err
	}

	defer dstFile.Close()

	old = old + " "
	new = new + " "
	var lineIdx int = 1
	scanner := bufio.NewScanner(srcFile)
	writer := bufio.NewWriter(dstFile)
	for scanner.Scan() {
		res, found, occ := replaceChar(scanner.Text(), old, new)
		if found {
			count += occ
			lines = append(lines, lineIdx)

		}
		fmt.Fprintf(writer, res)
		lineIdx++

	}
	return count, lines, err
}

func main() {

	old := "Go"
	new := "python"

	count, lines, err := openFile("test.txt", "testmodif.txt", old, new)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Println("===== Summary ====")
	defer fmt.Println("==== End ====")
	fmt.Printf("Word %v has been found %v times\n", old, count)
	fmt.Printf("test lines: %v\n", lines)

}
