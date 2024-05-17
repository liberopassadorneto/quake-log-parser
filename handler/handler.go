package handler

import (
	"encoding/json"
	"fmt"
	"github.com/liberopassadorneto/quake-parser/logger"
	"github.com/liberopassadorneto/quake-parser/parser"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"strings"
)

type Handler struct {
	parser *parser.Parser
}

func NewHandler() *Handler {
	return &Handler{
		parser: parser.NewParser(),
	}
}

// UploadFile processes the uploaded log file and returns parsed game data as JSON for CLI use.
func (h *Handler) UploadFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	mimetype := http.DetectContentType(buffer)
	if !strings.HasPrefix(mimetype, "text/plain") || filepath.Ext(filePath) != ".log" {
		return nil, fmt.Errorf("invalid file type")
	}

	tempFile, err := os.CreateTemp("", "upload-*.log")
	if err != nil {
		return nil, fmt.Errorf("failed to create temporary file: %w", err)
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, file)
	if err != nil {
		return nil, fmt.Errorf("failed to write to temporary file: %w", err)
	}

	games, err := h.parser.ParseLog(tempFile.Name())
	if err != nil {
		return nil, fmt.Errorf("failed to parse log: %w", err)
	}

	js, err := json.MarshalIndent(games, "", "\t")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON: %w", err)
	}

	logger.Log.Info("Finished writing report")
	return js, nil
}
