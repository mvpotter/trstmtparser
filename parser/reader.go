package parser

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/ledongthuc/pdf"
	"strings"
)

func readPdf(fileName string) ([]string, error) {
	records, err := readPdfInternal(fileName)
	if err != nil {
		return nil, err
	}

	return records, nil
}

func readPdfInternal(path string) ([]string, error) {
	f, r, err := pdf.Open(path)
	defer func() {
		_ = f.Close()
	}()
	if err != nil {
		return nil, fmt.Errorf("unable to open pdf: %w", err)
	}
	totalPage := r.NumPage()
	var records []string
	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}

		text, _ := p.GetPlainText(nil)
		if !strings.Contains(text, "Trade Republic Bank GmbH") {
			return nil, errors.New("no Trade Republic signature")
		}

		rows, _ := p.GetTextByRow()
		ptsInSymbol := 3.3
		for _, row := range rows {
			lastWordEnd := 0.0
			var rowBuffer bytes.Buffer
			for _, word := range row.Content {
				spaceBetweenLastWord := word.X - lastWordEnd
				rowBuffer.WriteString(strings.Repeat(" ", int(spaceBetweenLastWord/ptsInSymbol)))
				rowBuffer.WriteString(word.Font)
				rowBuffer.WriteString(word.S)
				lastWordEnd = word.X + float64(len(word.S))*ptsInSymbol
			}
			records = append(records, rowBuffer.String())
		}
	}
	return records, nil
}
