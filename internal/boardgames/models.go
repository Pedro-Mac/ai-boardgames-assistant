package boardgames

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Game struct {
	ID          bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Description string        `bson:"description" json:"description"`
	PDFFileID   bson.ObjectID `bson:"pdfFileId" json:"pdfFileId"`
	CreatedAt   time.Time     `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time     `bson:"updatedAt" json:"updatedAt"`
}

type RuleChunk struct {
	ID         bson.ObjectID     `bson:"_id,omitempty" json:"id"`
	GameID     bson.ObjectID     `bson:"gameId" json:"gameId"`
	ChunkText  string            `bson:"chunkText" json:"chunkText"`
	Embedding  []float32         `bson:"embedding" json:"embedding"`
	PageNumber int               `bson:"pageNumber" json:"pageNumber"`
	Metadata   map[string]string `bson:"metadata,omitempty" json:"metadata,omitempty"`
	CreatedAt  time.Time         `bson:"createdAt" json:"createdAt"`
}
