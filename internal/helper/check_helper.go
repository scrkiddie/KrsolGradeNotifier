package helper

import "fmt"

func DetectAndLogChanges(existingData, newData [][]string) []string {
	var changes []string

	for i, newRow := range newData {
		if i < len(existingData) {
			existingRow := existingData[i]
			if !equalRows(existingRow, newRow) {
				changes = append(changes, fmt.Sprintf("Perubahan terdeteksi pada: %v menjadi %v", existingRow, newRow))
			}
		} else {
			changes = append(changes, fmt.Sprintf("Baris baru terdeteksi pada baris %d: %v", i, newRow))
		}
	}

	for i := len(newData); i < len(existingData); i++ {
		changes = append(changes, fmt.Sprintf("Baris %d dihapus: %v", i, existingData[i]))
	}

	return changes
}

func equalRows(row1, row2 []string) bool {
	if len(row1) != len(row2) {
		return false
	}
	for i := range row1 {
		if row1[i] != row2[i] {
			return false
		}
	}
	return true
}
