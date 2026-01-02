package utils

import "strings"

type TextChunk struct {
	Text       string
	ChunkIndex int
}

func ChunkText(pdfContent string) []TextChunk {
	return paginateText(pdfContent)
}

func paginateText(pdfContent string) []TextChunk {
	currentChunkTokenSize := 0
	var currentChunk TextChunk
	chunks := make([]TextChunk, 0)
	currentIndex := 0

	stringsList := strings.Split(pdfContent, "\n\n")

	for i := range stringsList {
		// Handle huge single paragraphs (over 800 tokens)
		if currentChunkTokenSize == 0 && (len(stringsList[i])/4) > 800 {
			chunks = append(chunks, TextChunk{stringsList[i], currentIndex})
			currentIndex++
			continue
		}

		// Check if adding this paragraph stays within limit
		if currentChunkTokenSize+(len(stringsList[i])/4) <= 800 {
			// Add paragraph separator if needed
			if len(currentChunk.Text) > 0 {
				currentChunk.Text += "\n\n"
			}
			currentChunk.Text += stringsList[i]
			currentChunkTokenSize += len(stringsList[i]) / 4
		} else {
			// Current chunk is full, save it and start new one
			chunks = append(chunks, currentChunk)
			currentIndex++

			// Create overlap from previous chunk
			if len(chunks) > 0 {
				overlapText := chunks[len(chunks)-1].Text
				if len(overlapText) > 600 { // 150 tokens * 4 chars
					overlapText = overlapText[len(overlapText)-600:]
				}
				currentChunk = TextChunk{overlapText, currentIndex}
				currentChunkTokenSize = len(overlapText) / 4
			} else {
				currentChunk = TextChunk{"", currentIndex}
				currentChunkTokenSize = 0
			}

			// Add current paragraph to new chunk
			if len(currentChunk.Text) > 0 {
				currentChunk.Text += "\n\n"
			}
			currentChunk.Text += stringsList[i]
			currentChunkTokenSize += len(stringsList[i]) / 4
		}
	}

	if len(currentChunk.Text) > 0 {
		chunks = append(chunks, currentChunk)
	}

	return chunks
}
