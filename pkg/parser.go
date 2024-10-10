package pkg

import (
	"strings"
	"unicode"
)

func HeaderPresent(line string) bool {
	if len(line) < 6 {
		return false
	}

	caps := true
	for _, check := range line {
		if !unicode.IsUpper(check) && !unicode.IsSpace(check) && !unicode.IsPunct(check) {
			caps = false
			break
		}
	}
	return caps
}

func ParseText(text string) Structure {
	lines := strings.Split(text, "\n")
	doc := Structure{
		Data: make(map[string]string),
	}

	var currentSection *Section
	isData := true

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		if trimmedLine == "" {
			continue
		}

		if isData {
			if strings.Contains(trimmedLine, ":") {
				parts := strings.SplitN(trimmedLine, ":", 2)
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				doc.Data[key] = value
				continue
			} else {
				isData = false
				doc.Title = trimmedLine
				continue
			}
		}

		if HeaderPresent(trimmedLine) {
			if currentSection != nil {
				doc.Sections = append(doc.Sections, *currentSection)
			}
			currentSection = &Section{
				Header: trimmedLine,
			}
		} else {
			if currentSection == nil {
				currentSection = &Section{}
			}

			if len(currentSection.Body) > 0 && !strings.HasSuffix(currentSection.Body[len(currentSection.Body)-1], ".") {
				currentSection.Body[len(currentSection.Body)-1] += " " + trimmedLine
			} else {
				currentSection.Body = append(currentSection.Body, trimmedLine)
			}
		}
	}

	if currentSection != nil {
		doc.Sections = append(doc.Sections, *currentSection)
	}

	return doc
}
