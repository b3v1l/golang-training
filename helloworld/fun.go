package main

import (
	"fmt"
	"strings"
)

func TolowerStr(name string) (string, int) {

	return strings.ToLower(name), len(name)
}

func main() {

	lowername, len := TolowerStr("DFdfsFDSGFDSGSDgsdg")
	fmt.Printf("%s len(%d)\n", lowername, len)

	_, len = TolowerStr("TestTtestTTTTT")
	fmt.Printf("%s len(%d)\n", "TestTtestTTTTT", len)
}
