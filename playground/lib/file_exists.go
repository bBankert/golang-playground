package lib

import "os"

func FileExists(fileName string) (exists bool) {
	_, err := os.Stat(fileName)

	if err == nil {
		exists = true
	} else {
		exists = false
	}

	return exists
}

func ReadFileData(fileName string) (fileData string) {
	fileBytes, _ := os.ReadFile("Balance.txt")

	fileData = string(fileBytes)

	return fileData
}
