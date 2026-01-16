package services

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type EmbeddingService struct {
	client *openai.Client
}

func NewEmbeddingService(client *openai.Client) *EmbeddingService {
	return &EmbeddingService{
		client: client,
	}

}

func (s *EmbeddingService) GetEmbedding(ctx context.Context, text string) ([]float32, error) {
	resp, err := s.client.CreateEmbeddings(ctx, openai.EmbeddingRequest{
		Input: []string{text},
		Model: openai.SmallEmbedding3, // This is text-embedding-3-small
	})

	if err != nil {
		return nil, err
	}

	return resp.Data[0].Embedding, nil
}
