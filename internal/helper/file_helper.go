package helper

import (
	"bufio"
	"os"
	"strings"
)

func ReadDataFromFile(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "|")
		data = append(data, row)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return data, nil
}

func WriteDataToFile(filename string, data [][]string) {
	file, err := os.Create(filename)
	if err != nil {
		LogError("Error membuat file :", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, row := range data {
		line := strings.Join(row, "|")
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			LogError("Error menulis ke file :", err)
		}
	}
}
