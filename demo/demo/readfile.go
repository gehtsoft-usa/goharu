package demo

import (
	"bufio"
	"os"
)

func readFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil, statsErr
	}

	var size int64 = stats.Size()
	bytes := make([]byte, size)

	buffReader := bufio.NewReader(file)
	_, err = buffReader.Read(bytes)

	return bytes, err
}
