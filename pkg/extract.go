// extracting the plain text from pdf

package pkg

import (
	"strings"

	"github.com/ledongthuc/pdf"
)

// extracts text from the PDF and return it as a slice of paragraphs
func ExtractText(path string) ([]string, error) {
	pdfFileStream, pdfDoc, err := pdf.Open(path)
	if err != nil {
		return nil, err
	}
	defer pdfFileStream.Close()

	var paragraphs []string

	for i := 1; i <= pdfDoc.NumPage(); i++ {
		currentPage := pdfDoc.Page(i)
		if currentPage.V.IsNull() {
			continue
		}

		pageText, err := currentPage.GetPlainText(nil)
		if err != nil {
			return nil, err
		}

		// Split the text by newlines to preserve paragraph structure
		pageParagraphs := strings.Split(pageText, "\n\n")

		for _, para := range pageParagraphs {
			trimmedPara := strings.TrimSpace(para)
			if len(trimmedPara) > 0 {
				paragraphs = append(paragraphs, trimmedPara)
			}
		}
	}

	return paragraphs, nil
}
