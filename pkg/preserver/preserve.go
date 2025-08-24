package preserver

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"

	yaml "sigs.k8s.io/yaml/goyaml.v3"
)

// Codec defines marshal/unmarshal behavior.
type Codec interface {
	Marshal(v any) ([]byte, error)
	Unmarshal(data []byte, v any) error
}

// JSONCodec implements Codec for JSON.
type JSONCodec struct{}

// YAMLCodec implements Codec for YAML.
type YAMLCodec struct{}

func (JSONCodec) Marshal(v any) ([]byte, error)      { return json.MarshalIndent(v, "", "  ") }
func (JSONCodec) Unmarshal(data []byte, v any) error { return json.Unmarshal(data, v) }

func (YAMLCodec) Marshal(v any) ([]byte, error)      { return yaml.Marshal(v) }
func (YAMLCodec) Unmarshal(data []byte, v any) error { return yaml.Unmarshal(data, v) }

type Preserve[T any] interface {
	Save(v T) error
	Load(v *T) error
}

type FilePreserver[T any] struct {
	filePath string
	codec    Codec
}

func (fp *FilePreserver[T]) Save(v T) error {
	data, err := fp.codec.Marshal(v)
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
	return fp.codec.Unmarshal(data, v)
}

// NewConfigHandler selects codec based on file extension.
func NewConfigHandler[T any](fp string) (Preserve[T], error) {
	ext := strings.ToLower(filepath.Ext(fp))
	var codec Codec

	switch ext {
	case ".json":
	case ".dot":
		codec = JSONCodec{}
	case ".yaml", ".yml":
		codec = YAMLCodec{}
	default:
		return nil, errors.New("unsupported file type: " + ext)
	}

	return &FilePreserver[T]{filePath: fp, codec: codec}, nil
}
