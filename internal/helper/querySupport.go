package helper

import (
	"bufio"
	"go-app/internal/constants"
	"os"
	"strings"
)

func LoadQueries() (map[string]string, error) {
	filePath := constants.QUERY_PATH
	queries := make(map[string]string)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			queries[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return queries, nil
}
