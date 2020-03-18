package cert

import (
	"fmt"
	"strings"
	"time"
)

var Maxlencourse = 20
var Maxlenname = 30

type Cert struct {
	Course string
	Name   string
	Date   time.Time

	LabelTitle         string
	LabelCompletion    string
	LabelPresented     string
	LabelParticipation string
	LabelDate          string
}

type Saver interface {
	Save(c Cert) error
}

func New(course, name, date string) (*Cert, error) {

	c, err := validate(course)
	if err != nil {
		return nil, err
	}

	n, err := validatename(name)
	if err != nil {
		return nil, err
	}
	d, err := parseDate(date)

	cert := &Cert{
		Course:             c,
		Name:               n,
		LabelTitle:         fmt.Sprintf("%v Certificat - %v", c, n),
		LabelCompletion:    "Certificat of Completion",
		LabelPresented:     "This certificat is presented to",
		LabelParticipation: fmt.Sprintf("For participation in the %v", c),
		LabelDate:          fmt.Sprintf("Date: %v", d.Format("02/01/2006")),
	}
	return cert, nil
}

func validate(course string) (string, error) {
	c, err := TestValidStr(course, Maxlencourse)
	if err != nil {
		return "", err
	}

	if !strings.HasSuffix(c, " course") {
		c = c + " course"
	}
	return strings.ToUpper(c), nil
}

func TestValidStr(str string, Maxlencourse int) (string, error) {
	c := strings.TrimSpace(str)
	if len(c) <= 0 {
		return c, fmt.Errorf("Invalid string. got='%s' -- len(%d)", c, len(c))
	} else if len(c) >= Maxlencourse {
		return c, fmt.Errorf("Too long. got='%s' -- len(%d)", c, len(c))
	}
	return c, nil
}

func validatename(name string) (string, error) {
	n := name
	if len(n) == 0 {
		return "", fmt.Errorf("errrrrrrrror")
	} else if len(n) >= Maxlenname {
		return "", fmt.Errorf("errrrrrrrror")
	}
	return strings.ToTitle(n), nil
}

func parseDate(date string) (time.Time, error) {

	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return t, err
	}
	return t, nil

}
