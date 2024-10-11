// Rendering the extracted text in the file

package pkg

import (
	"os"
)

func RenderParsedTextToFile(paragraphs []string, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write each paragraph to the file, with a double newline to preserve spacing
	for _, para := range paragraphs {
		_, err := file.WriteString(para + "\n\n")
		if err != nil {
			return err
		}
	}

	return nil
}
