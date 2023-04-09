package models

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestBook(t *testing.T) {
	authorID := uuid.New()

	tests := []struct {
		name    string
		payload string
	}{
		{
			name: "CreateBookRequest should create a new book",
			payload: `{
				"title": "The Hitchhiker's Guide to the Galaxy",
				"author": {
					"id": "` + authorID.String() + `",
					"name": "Douglas Adams"
				},
				"created_at": "` + time.Now().Format(time.RFC3339) + `"
			}`,
		},
		{
			name: "UpdateBook should update a book",
			payload: `{
				"title": "The Restaurant at the End of the Universe",
				"author": {
					"id": "` + authorID.String() + `",
					"name": "Douglas Adams"
				},
				"updated_at": "` + time.Now().Format(time.RFC3339) + `"
			}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			switch test.name {
			case "CreateBookRequest should create a new book":
				var req CreateBookRequest
				err := json.Unmarshal([]byte(test.payload), &req)
				if err != nil {
					t.Errorf("Error unmarshaling JSON: %v", err)
				}

				if req.Title != "The Hitchhiker's Guide to the Galaxy" {
					t.Errorf("Expected title to be 'The Hitchhiker's Guide to the Galaxy', but got %s", req.Title)
				}

				if req.Author.ID != authorID {
					t.Errorf("Expected author ID to be %s, but got %s", authorID.String(), req.Author.ID.String())
				}

			case "UpdateBook should update a book":
				var update UpdateBook
				err := json.Unmarshal([]byte(test.payload), &update)
				if err != nil {
					t.Errorf("Error unmarshaling JSON: %v", err)
				}

				if update.Title != "The Restaurant at the End of the Universe" {
					t.Errorf("Expected title to be 'The Restaurant at the End of the Universe', but got %s", update.Title)
				}

				if update.Author.ID != authorID {
					t.Errorf("Expected author ID to be %s, but got %s", authorID.String(), update.Author.ID.String())
				}
			}
		})
	}
}
