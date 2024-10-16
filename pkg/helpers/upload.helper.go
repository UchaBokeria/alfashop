package helpers

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"main/internal/model"
	"main/pkg/globals"
	"main/pkg/storage"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
)

type UploadResponse struct {
	ID      int
	Message string
	Success bool
}

func File(file *multipart.FileHeader) *UploadResponse {
	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return &UploadResponse{ID: -1, Message: "Error opening received file", Success: false}
	}
	defer src.Close()

	// Calculate SHA-256 hash of the file contents
	hash := sha256.New()
	if _, err := io.Copy(hash, src); err != nil {
		return &UploadResponse{ID: -1, Message: "Error calculating hash", Success: false}
	}

	// Reset src to the beginning to read again
	src.Seek(0, 0)
	extension := GetFileExtension(file)
	hashName := hex.EncodeToString(hash.Sum(nil))

	// Create a new file on the server to store the uploaded file
	dst, err := os.Create("./public" + globals.Env.Uploads + hashName + extension)
	if err != nil {
		return &UploadResponse{ID: -1, Message: "Error creating file on server", Success: false}
	}
	defer dst.Close()

	// Copy the file from the form data to the destination file
	if _, err = io.Copy(dst, src); err != nil {
		return &UploadResponse{ID: -1, Message: "Error copying file to destination", Success: false}
	}

	if len(extension) < 2 {
		return &UploadResponse{ID: -1, Message: "File type " + extension + " has a problem", Success: false}
	}

	var Type model.File_types
	result := storage.DB.Where(&model.File_types{Ext: extension[1:]}).Last(&Type)

	if result.Error != nil {
		log.Print(result)
		return &UploadResponse{ID: -1, Message: "Server can't accept " + extension + " type files", Success: false}
	}

	var File model.Files = model.Files{
		Name:       hashName + extension,
		Original:   file.Filename,
		Size:       int(file.Size),
		Location:   globals.Env.Uploads,
		Path:       globals.Env.Uploads + hashName + extension,
		Compressed: false,
		Base64:     "",
		TypeID:     int(Type.ID),
	}

	Result := storage.DB.Create(&File)
	if Result.Error != nil || Result.RowsAffected < 1 {
		log.Print(Result)
		return &UploadResponse{ID: -1, Message: "File uploaded but was not saved in database", Success: false}
	}

	return &UploadResponse{ID: int(File.ID), Message: "Successfully uploaded", Success: true}
}

func GetDbTypeIdByExtension(extension string) int {
	types := map[string]string{
		"jpeg": "2",
		"jpg":  "3",
		"png":  "4",
		"gif":  "5",
		"pdf":  "8",
	}

	indeger, _ := strconv.Atoi(types[extension])
	return indeger
}

func GetFileExtension(file *multipart.FileHeader) string {
	header, err := file.Open()
	if err != nil {
		return ""
	}
	defer header.Close()

	extension := ""
	for i := len(file.Filename) - 1; i >= 0; i-- {
		if file.Filename[i] == '.' {
			extension = file.Filename[i:]
			break
		}
	}
	return strings.ToLower(extension)
}
