// Rendering the extracted text in the file

package pkg

import (
	"bufio"
	"fmt"
	"os"
)

func RenderParsedText(doc Structure, filePath string) error {
	// Open the file for writing (or create it if it doesn't exist)
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	_, err = writer.WriteString(fmt.Sprintf("Title: %s\n\n", doc.Title))
	if err != nil {
		return fmt.Errorf("failed to write title: %v", err)
	}

	_, err = writer.WriteString("Text:\n")
	if err != nil {
		return fmt.Errorf("failed to write text label: %v", err)
	}

	for key, value := range doc.Data {
		_, err := writer.WriteString(fmt.Sprintf("%s: %s\n", key, value))
		if err != nil {
			return fmt.Errorf("failed to write metadata: %v", err)
		}
	}

	_, err = writer.WriteString("\n")
	if err != nil {
		return fmt.Errorf("failed to write new line after metadata: %v", err)
	}

	for _, section := range doc.Sections {
		_, err := writer.WriteString(fmt.Sprintf("Section: %s\n", section.Header))
		if err != nil {
			return fmt.Errorf("failed to write section header: %v", err)
		}

		for _, paragraph := range section.Body {
			_, err := writer.WriteString(fmt.Sprintf("%s\n", paragraph))
			if err != nil {
				return fmt.Errorf("failed to write paragraph: %v", err)
			}
		}

		_, err = writer.WriteString("\n")
		if err != nil {
			return fmt.Errorf("failed to write new line after section: %v", err)
		}
	}

	// Flush the buffered writer to ensure all data is written to the file
	if err := writer.Flush(); err != nil {
		return fmt.Errorf("failed to flush writer: %v", err)
	}

	return nil
}
