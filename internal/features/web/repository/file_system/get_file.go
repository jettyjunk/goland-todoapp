package web_filesystem_repository

import (
	"errors"
	"fmt"
	"os"

	core_errors "github.com/jettyjunk/goland-todoapp/internal/core/errors"
)

func (r *WebRepository) GetFile(filePath string) ([]byte, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("file: %s: %w", filePath, core_errors.ErrNotFounde)
		}

		return nil, fmt.Errorf("get file: %s: %w", file, err)
	}

	return file, nil
}
