package boardgames

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ServerDependencies interface {
	GetDatabaseClient() *mongo.Client
	GetRouter() *chi.Mux
}

func RegisterRoutes(server ServerDependencies) {
	router := server.GetRouter()
	dbClient := server.GetDatabaseClient()

	router.Post("/boardgames/game-upload", handleRulesUploadFile(dbClient))
}

func handleRulesUploadFile(dbClient *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Handler logic goes here

		err := r.ParseMultipartForm(100 << 20) // 50MB
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		file, _, err := r.FormFile("file")

		if err != nil {
			http.Error(w, "Error getting file", http.StatusBadRequest)
			return
		}

		defer file.Close()

		name := r.FormValue("name")
		description := r.FormValue("description")

		if name == "" {
			http.Error(w, "Game name is required", http.StatusBadRequest)
			return
		}

		db := dbClient.Database("boardgames")
		bucket := db.GridFSBucket()

		uploadStream, err := bucket.OpenUploadStream(r.Context(), name)
		if err != nil {
			http.Error(w, "Error opening upload stream", http.StatusInternalServerError)
			return
		}
		defer uploadStream.Close()

		// Copy file to GridFS
		_, err = io.Copy(uploadStream, file)
		if err != nil {
			http.Error(w, "Error uploading file", http.StatusInternalServerError)
			return
		}

		fileID := uploadStream.FileID.(bson.ObjectID)

		game := Game{
			Name:        name,
			Description: description,
			PDFFileID:   fileID,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		gamesCollection := db.Collection("games")
		result, err := gamesCollection.InsertOne(r.Context(), game)
		if err != nil {
			http.Error(w, "Error saving game", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"gameId":  result.InsertedID,
			"message": "Game rules uploaded successfully",
		})
	}
}
