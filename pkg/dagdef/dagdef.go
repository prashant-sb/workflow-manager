package dagdef

import (
	"os"
)

func GetDAGFromDefination(in string) (string, error) {
	data, err := os.ReadFile(in)
	if err != nil {
		return in, err
	}

	return string(data), nil
}
