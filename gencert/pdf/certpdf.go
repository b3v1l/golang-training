package pdf

import (
	"fmt"
	"os"
	"path"

	"github.com/signintech/gopdf"
	"training.go/gencert/cert"
)

type PdfSaver struct {
	OutputDir string
}

func New(OutputDir string) (*PdfSaver, error) {

	var p *PdfSaver
	err := os.MkdirAll(OutputDir, os.ModePerm)
	if err != nil {
		return p, err
	}

	p = &PdfSaver{
		OutputDir: OutputDir,
	}
	return p, nil
}

func (p *PdfSaver) Save(cert cert.Cert) error {

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	//save file
	filename := fmt.Sprintf("%v.pdf", cert.LabelTitle)
	path := path.Join(p.OutputDir, filename)

}
