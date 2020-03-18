package main

import (
	"flag"
	"fmt"
	"os"

	"training.go/cli-dictionary/dictionary"
)

func main() {

	action := flag.String("action", "list", "List Dictionary content")

	d, err := dictionary.New("./badger")
	HandleError(err)
	defer d.Close()

	flag.Parse()

	switch *action {

	case "list":
		actionList(d)

	case "add":
		actionAdd(d, flag.Args())

	case "remove":
		actionDel(d, flag.Args())

	case "define":
		actionDef(d, flag.Args())

	default:
		fmt.Printf("Unknow option %v\n", *action)
	}

	//	d.Add("python", "interpreted blablabla")
	//	words, entries, _ := d.List()
	//	for _, word := range words {
	//	fmt.Println(entries[word])

	//	}
	//	entry, _ := d.Get("golang")
	//	fmt.Println(entry)

}

func actionAdd(d *dictionary.Dictionary, args []string) {
	word := args[0]
	definition := args[1]
	err := d.Add(word, definition)
	HandleError(err)
	fmt.Printf("%v added to the dictionary\n", word)
}

func actionDel(d *dictionary.Dictionary, args []string) {

	word := args[0]
	err := d.Remove(word)
	HandleError(err)
	fmt.Printf("%v removed from dictionary\n", word)

}

func actionDef(d *dictionary.Dictionary, args []string) {

	word := args[0]
	entry, err := d.Get(word)
	HandleError(err)
	fmt.Printf("[*] %v's definition =\n%v\n", word, entry)
}

func actionList(d *dictionary.Dictionary) {

	words, entries, err := d.List()
	HandleError(err)
	fmt.Println("Dictionary content :")
	for _, word := range words {
		fmt.Println(entries[word])

	}
}

func HandleError(err error) {

	if err != nil {
		fmt.Printf("Dictionary error:%v\n", err)
		os.Exit(1)
	}

}
