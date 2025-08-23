package dagdef

import "os"

// common DAG definition configuration if any

func GetDAGFromDefination(in string) (string, error) {
	data, err := os.ReadFile(in)
	if err != nil {
		return in, err
	}

	return string(data), nil
}
