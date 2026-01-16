package utils

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/ledongthuc/pdf"
)

func ExtractPdfContent(fileBytes []byte) (string, error) {
	// Placeholder for PDF content extraction logic

	if fileBytes == nil {
		return "", fmt.Errorf("fileBytes is nil")
	}

	reader := bytes.NewReader(fileBytes)

	pdfReader, err := pdf.NewReader(reader, int64(len(fileBytes)))

	if err != nil {
		fmt.Printf("failed to open PDF: %v\n", err)
		return "", fmt.Errorf("failed to open PDF: %w", err)
	}

	var textBuilder strings.Builder
	numPages := pdfReader.NumPage()
	for i := 1; i <= numPages; i++ {
		page := pdfReader.Page(i)
		if page.V.IsNull() {
			continue
		}
		pageText, err := page.GetPlainText(nil)
		if err != nil {
			fmt.Printf("failed to extract text from page %d: %v\n", i, err)
			continue
		}

		textBuilder.WriteString(pageText)
		textBuilder.WriteString("\n\n")
	}

	text := textBuilder.String()

	if text == "" {
		return "", fmt.Errorf("no text extracted from PDF")
	}

	return text, nil
}
