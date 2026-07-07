package admin

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func UploadDir() string {
	if d := os.Getenv("UPLOAD_DIR"); d != "" {
		return d
	}
	return "./uploads"
}

var extensionsImages = map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".webp": true}
var extensionsDocs = map[string]bool{".pdf": true}

const (
	maxImageSize int64 = 5 << 20
	maxDocSize   int64 = 10 << 20
)

func SaveUpload(file multipart.File, header *multipart.FileHeader, sousDossier string, kind string) (string, error) {
	ext := strings.ToLower(filepath.Ext(header.Filename))

	var maxSize int64
	if kind == "document" {
		if !extensionsDocs[ext] {
			return "", errors.New("type de document non autorise (PDF uniquement)")
		}
		maxSize = maxDocSize
	} else {
		if !extensionsImages[ext] {
			return "", errors.New("type d'image non autorise (jpg, jpeg, png, webp)")
		}
		maxSize = maxImageSize
	}
	if header.Size > maxSize {
		return "", fmt.Errorf("fichier trop volumineux (max %d Mo)", maxSize>>20)
	}

	tete := make([]byte, 512)
	n, _ := file.Read(tete)
	mimeType := http.DetectContentType(tete[:n])
	if kind == "document" {
		if !strings.HasPrefix(mimeType, "application/pdf") {
			return "", errors.New("contenu invalide : un PDF est attendu")
		}
	} else if !strings.HasPrefix(mimeType, "image/") {
		return "", errors.New("contenu invalide : une image est attendue")
	}
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return "", err
	}

	buf := make([]byte, 16)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	nom := hex.EncodeToString(buf) + ext

	dossier := filepath.Join(UploadDir(), sousDossier)
	if err := os.MkdirAll(dossier, 0755); err != nil {
		return "", err
	}
	dst, err := os.Create(filepath.Join(dossier, nom))
	if err != nil {
		return "", err
	}
	defer dst.Close()
	if _, err := io.Copy(dst, file); err != nil {
		return "", err
	}

	return "uploads/" + sousDossier + "/" + nom, nil
}
