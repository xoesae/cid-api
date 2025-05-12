package importer

import (
	"encoding/xml"
	"fmt"
	"golang.org/x/text/encoding/ianaindex"
	"io"
	"os"
	"strings"
)

type Subcategory struct {
	Code string `xml:"codsubcat,attr"`
	Name string `xml:"nome"`
}

type Category struct {
	Code          string        `xml:"codcat,attr"`
	Name          string        `xml:"nome"`
	Subcategories []Subcategory `xml:"subcategoria"`
}

type Group struct {
	Code       string     `xml:"codgrupo,attr"`
	Initial    string     `xml:"inicial,attr"`
	Final      string     `xml:"final,attr"`
	Name       string     `xml:"nome"`
	Categories []Category `xml:"categoria"`
}

type Chapter struct {
	Num       string  `xml:"numcap,attr"`
	CodeRange string  `xml:"codcap,attr"`
	Roman     string  `xml:"romano,attr"`
	Initial   string  `xml:"inicial,attr"`
	Final     string  `xml:"final,attr"`
	Name      string  `xml:"nome"`
	Groups    []Group `xml:"grupo"`
}

type CID10 struct {
	Chapters []Chapter `xml:"capitulo"`
}

func StreamChapters(xmlPath string, onChapter func(ch Chapter) error) error {
	f, err := os.Open(xmlPath)
	if err != nil {
		return fmt.Errorf("error on open XML file: %w", err)
	}
	defer f.Close()

	decoder := xml.NewDecoder(f)

	// handle the charset ISO-8859-1
	decoder.CharsetReader = func(charset string, reader io.Reader) (io.Reader, error) {
		enc, err := ianaindex.IANA.Encoding(charset)
		if err != nil {
			return nil, fmt.Errorf("charset %s: %s", charset, err.Error())
		}
		if enc == nil {
			// Assume it's compatible with (a subset of) UTF-8 encoding
			// Bug: https://github.com/golang/go/issues/19421
			return reader, nil
		}
		return enc.NewDecoder().Reader(reader), nil
	}

	for {
		tok, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error on read token: %w", err)
		}

		switch se := tok.(type) {
		case xml.StartElement:
			if se.Name.Local == "capitulo" {
				var ch Chapter
				fmt.Println(se)

				if err := decoder.DecodeElement(&ch, &se); err != nil {
					return fmt.Errorf("error on decode chapter: %w", err)
				}
				if err := onChapter(ch); err != nil {
					return err
				}
			}
		case xml.CharData:
			tok = xml.CharData(strings.ReplaceAll(string(se), "&cruz;", ""))
		}
	}

	return nil
}
