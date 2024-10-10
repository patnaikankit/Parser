// extracting the plain text from pdf

package pkg

import "github.com/ledongthuc/pdf"

func ExtractText(path string) (string, error) {
	pdfFileStream, pdf, err := pdf.Open(path)

	if err != nil {
		return "", err
	}

	defer pdfFileStream.Close()

	var text string
	for i := 1; i <= pdf.NumPage(); i++ {
		currentPage := pdf.Page(i)

		if currentPage.V.IsNull() {
			continue
		}

		pageText, err := currentPage.GetPlainText(nil)
		if err != nil {
			return " ", err
		}

		text += pageText
	}
	return text, nil
}
