package main

import (
	"flag"
	"fmt"
	"os"

	"training.go/gencert/cert"
	"training.go/gencert/html"
	"training.go/gencert/pdf"
)

func main() {

	outputType := flag.String("type", "pdf", "Output format.")
	filecsv := flag.String("file", "", "import CSV content.")

	flag.Parse()

	var saver cert.Saver
	var err error
	switch *outputType {

	case "pdf":
		saver, err = pdf.New("output")

	case "html":
		saver, err = html.New("output")
	default:
		fmt.Printf("Unknown output format type %v\n", outputType)
	}

	if len(*filecsv) <= 0 {
		fmt.Println("Invalid file, exiting")
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Pdf generator error %v\n", err)
		os.Exit(1)
	}

	certs, err := cert.Parsecsv(*filecsv)
	if err != nil {
		fmt.Printf("Could not parse %v file, exiting", err)
		os.Exit(1)
	}

	for _, c := range certs {
		err := saver.Save(*c)
		if err != nil {
			fmt.Printf("Could not save the certificat %v\n", err)
		}

	}

	//	saver.Save(*c)
}
