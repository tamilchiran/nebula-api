package comics

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ImageFile struct {
	FileID string `json:"file_id"`
}

type Text struct {
	Value       string   `json:"value"`
	Annotations []string `json:"annotations"`
}

type Content struct {
	Type      string     `json:"type"`
	ImageFile *ImageFile `json:"image_file,omitempty"`
	Text      *Text      `json:"text,omitempty"`
}

type Data struct {
	Content []Content `json:"content"`
}

func PublishArticle(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusInternalServerError)
		return
	}

	// Unmarshal JSON
	var data Data
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the parsed data
	fmt.Printf("Content:\n")
	for _, content := range data.Content {
		fmt.Printf("Type: %s\n", content.Type)
		if content.ImageFile != nil {
			fmt.Printf("File ID: %s\n", content.ImageFile.FileID)
		}
		if content.Text != nil {
			fmt.Printf("Text: %s\n", content.Text.Value)
		}
		fmt.Println()
	}

	w.WriteHeader(http.StatusOK)
	w.Write(([]byte("Received request")))
}
