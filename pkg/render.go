// Rendering the extracted text on the console

package pkg

import "fmt"

func RenderParsedText(doc Structure) {
	fmt.Printf("Title: %s\n\n", doc.Title)
	fmt.Println("Text: ")
	for key, value := range doc.Data {
		fmt.Printf("%s: %s\n", key, value)
	}

	fmt.Println()

	for _, section := range doc.Sections {
		fmt.Printf("Section: %s \n", section.Header)
		for _, paragraph := range section.Body {
			fmt.Printf("%s \n", paragraph)
		}
		fmt.Println()
	}
}
