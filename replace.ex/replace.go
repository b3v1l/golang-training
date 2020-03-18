package main

import (
	"fmt"
	"strings"
)

func procLine(line, old, new string) (found bool, res string, occ int) {

	//lowerOld := strings.ToLower(old)
	//lowerNew := strings.ToLower(new)

	res = line
	if strings.Contains(line, old) {
		found = true
		occ += strings.Count(line, old)
		res = strings.Replace(line, old, new, -1)

	}
	return false, res, occ
}

//#funcFindReplaceFile(src, old, new string) (occ int, lines []int , err error) {
// :l
//	f := f.open(filename)
//	defer = f.Close()

func main() {

	_, res, occ := procLine("Go on va tous go vers la mer, gogoog", "go", "switch")
	fmt.Println(res, occ)

}
