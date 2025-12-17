package utils

import (
	"bytes"
	"io"
	"mime/multipart"

	"github.com/ledongthuc/pdf"
)

func ExtractPdfContent(file *multipart.File) error {
	// Placeholder for PDF content extraction logic

	fileBytes, err := io.ReadAll(*file)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(fileBytes)

	pdfReader, err := pdf.NewReader(reader, int64(len(fileBytes)))

	if err != nil {
		return err
	}
}
