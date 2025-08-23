package preserver

import (
	"encoding/json"
	"os"
)

type Preserve[T any] interface {
	Save(v T) error
	Load(v *T) error
}

type FilePreserver[T any] struct {
	filePath string
}

func (fp *FilePreserver[T]) Save(v T) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fp.filePath, data, 0644)
}

func (fp *FilePreserver[T]) Load(v *T) error {
	data, err := os.ReadFile(fp.filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func NewFilePreserver[T any](fp string) Preserve[T] {
	return &FilePreserver[T]{filePath: fp}
}
