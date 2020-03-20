package html

import (
	"fmt"
	"os"
	"path"
	"text/template"

	"training.go/gencert/cert"
)

type HtmlSaver struct {
	Outdir string
}

func New(outputdir string) (*HtmlSaver, error) {

	var h *HtmlSaver
	err := os.MkdirAll(outputdir, os.ModePerm)
	if err != nil {
		return h, err
	}

	h = &HtmlSaver{
		Outdir: outputdir,
	}
	return h, nil
}

func (p *HtmlSaver) Save(cert cert.Cert) error {

	t, err := template.New("certificate").Parse(tpl)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%v.html", cert.LabelTitle)
	path := path.Join(p.Outdir, filename)

	f, err := os.Create(path)
	if err != nil {
		return nil
	}
	defer f.Close()
	err = t.Execute(f, cert)
	if err != nil {
		return nil
	}
	fmt.Printf("Saved Certificat in %v\n", path)
	return nil
}

var tpl = `
<!DOCTYPE html>
<html>
        <head>
                <meta charset="UTF-8">
        <title>{{.LabelTitle}}</title>
        <style>
            body {
                text-align: center;
                font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif
            }

            h1 {
                font-size: 3em;
            }

        </style>
    </head>
    <body>
        <h1>{{.LabelCompletion}}</h1>
        <h2><em>{{.LabelPresented}}</em></h2>
        <h1>{{.Name}}</h1>
        <h2>{{.LabelParticipation}}</h2>
        <p>
            <em>{{.LabelDate}}</em>
        </p>
    </body>
</html>
`
