package cert

import (
	"encoding/csv"
	"os"
)

func Parsecsv(filename string) ([]*Cert, error) {

	certs := make([]*Cert, 0)
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return certs, err
	}
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	for _, record := range records {
		c, err := New(record[0], record[1], record[2])
		if err != nil {
			return certs, err
		}
		certs = append(certs, c)

	}
	return certs, nil
}
