package utils

import (
	"compress/gzip"
	"encoding/xml"
	"os"
)

type document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int
}

// Read file from the path, then put it in an array
func LoadDocuments(path string) ([]document, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer f.Close()
	gz, err := gzip.NewReader(f)

	if err != nil {
		return nil, err
	}

	defer gz.Close()
	dec := xml.NewDecoder(gz)
	// an array of documents
	dump := struct {
		Documents []document `xml:"doc"`
	}{}
	err = dec.Decode(&dump)
	if err != nil {
		return nil, err
	}
	docs := dump.Documents
	for i := range docs {
		docs[i].ID = i
	}

	return docs, nil
}
