package controllers

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
	"net/http"
	"os"
	"strings"
)

// UploadRequest represents the expected JSON request body
type UploadRequest struct {
	Base64Str string `json:"base64_str"`
}

// UploadResponse represents the JSON response body
type UploadResponse struct {
	URL string `json:"url"`
}

// UploadHandler handles the upload of a base64 image to Cloudinary
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req UploadRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Decode the base64 image
	fileBytes, err := decodeBase64(req.Base64Str)
	if err != nil {
		http.Error(w, "Failed to decode base64 string", http.StatusBadRequest)
		return
	}

	// Initialize Cloudinary credentials
	cld, ctx := credentials()

	// Upload the base64 image and get the URL
	url, err := uploadBase64Image(cld, ctx, fileBytes)
	if err != nil {
		http.Error(w, "Failed to upload image", http.StatusInternalServerError)
		return
	}

	if url == "" {
		http.Error(w, "Received empty URL from Cloudinary", http.StatusInternalServerError)
		return
	}

	// Prepare the response
	response := UploadResponse{URL: url}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// decodeBase64 decodes a base64 string into bytes
func decodeBase64(data string) ([]byte, error) {
	// Remove any data URL prefix
	if strings.HasPrefix(data, "data:") {
		data = strings.SplitN(data, ",", 2)[1]
	}
	return base64.StdEncoding.DecodeString(data)
}

// credentials initializes and returns the Cloudinary client and context
func credentials() (*cloudinary.Cloudinary, context.Context) {
	cld, err := cloudinary.NewFromParams(os.Getenv("CLOUDINARY_NAME"), os.Getenv("API_KEY"), os.Getenv("API_SECRETE"))
	if err != nil {
		fmt.Printf("Error initializing Cloudinary: %v\n", err)
	}
	ctx := context.Background()
	return cld, ctx
}

// uploadBase64Image uploads the binary image data to Cloudinary
func uploadBase64Image(cld *cloudinary.Cloudinary, ctx context.Context, fileBytes []byte) (string, error) {
	// Generate a unique identifier for the file
	publicID := "file-" + uuid.NewString()

	uploadResult, err := cld.Upload.Upload(ctx, bytes.NewReader(fileBytes), uploader.UploadParams{
		Folder:       "recipes", // Optional: Specify a folder in Cloudinary
		PublicID:     publicID,
		ResourceType: "auto",
	})
	if err != nil {
		fmt.Printf("Upload error: %v\n", err)
		return "", fmt.Errorf("upload error: %v", err)
	}

	if uploadResult.SecureURL == "" {
		fmt.Println("Empty URL received from Cloudinary")
	}

	return uploadResult.SecureURL, nil
}
