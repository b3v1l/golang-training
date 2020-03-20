package pdf

import (
	"fmt"
	"os"
	"path"

	"github.com/jung-kurt/gofpdf"

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

	pdf := gofpdf.New("L", "", "A4", "arial")

	pdf.SetTitle(cert.LabelTitle, false)
	pdf.AddPage()

	//background
	background(pdf)

	//Header
	header(pdf, &cert)
	pdf.Ln(30)

	//body
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, cert.LabelPresented, "C")
	pdf.Ln(30)
	//body uername

	pdf.SetFont("Times", "B", 40)
	pdf.WriteAligned(0, 50, cert.Name, "C")
	pdf.Ln(30)
	//body participation

	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, cert.LabelParticipation, "C")
	pdf.Ln(30)
	//body date :

	pdf.SetFont("Helvetica", "I", 15)
	pdf.WriteAligned(0, 50, cert.LabelDate, "C")
	pdf.Ln(30)

	footer(pdf, &cert)
	//save file
	filename := fmt.Sprintf("%v.pdf", cert.LabelTitle)
	path := path.Join(p.OutputDir, filename)
	err := pdf.OutputFileAndClose(path)
	if err != nil {
		return err
	}
	fmt.Printf("Saved cert in '%v'\n", path)
	return nil
}

func header(pdf *gofpdf.Fpdf, cert *cert.Cert) {

	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}

	margin := 30.0
	x := 0.0
	imageWidth := 30.0
	filename := "img/gopher.png"
	x, _ = pdf.GetPageSize()

	pdf.ImageOptions("img/gopher.png", x+margin, 20, imageWidth, 0, false, opts, 0, "")

	pageWidth, _ := pdf.GetPageSize()
	x = pageWidth - imageWidth
	pdf.ImageOptions(filename, x-margin, 20, imageWidth, 0, false, opts, 0, "")

	pdf.SetFont("Helvetica", "", 40)
	pdf.WriteAligned(0, 50, cert.LabelCompletion, "C")
}

func footer(pdf *gofpdf.Fpdf, c *cert.Cert) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	ImageWidth := 50.0
	filename := "img/stamp.png"
	x, y := pdf.GetPageSize()
	//fmt.Printf("%v //// x = %v, y = %v", ImageWidth, x, y)
	l := (x - ImageWidth - 20.0)
	h := y - ImageWidth - 10.0
	pdf.ImageOptions(filename, l, h, ImageWidth, 0, false, opts, 0, "")

}

func background(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	x, y := pdf.GetPageSize()
	pdf.ImageOptions("img/background.png",
		0, 0, x, y, false, opts, 0, "",
	)

	//pdf.Image()

}
