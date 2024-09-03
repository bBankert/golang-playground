package file_operations

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type FileManager struct {
	ExistingFilePath string
	NewFilePath      string
}

func New(inputFile, outputFile string) *FileManager {
	return &FileManager{
		ExistingFilePath: inputFile,
		NewFilePath:      outputFile,
	}
}

func (fm FileManager) ReadSavedPriceData() ([]float64, error) {
	file, err := os.Open(fm.ExistingFilePath)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("unable to open file")
	}

	scanner := bufio.NewScanner(file)

	var prices []float64

	for scanner.Scan() {
		price, conversionError := strconv.ParseFloat(scanner.Text(), 64)

		if conversionError != nil {
			fmt.Println(conversionError)
			file.Close()
			return nil, errors.New("unable to convert to float")
		}

		prices = append(prices, price)
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println(err)
		file.Close()
		return nil, errors.New("unable to read file")
	}

	return prices, nil
}

func (fm FileManager) SavePriceData(data any) error {
	file, err := os.Create(fm.NewFilePath)

	if err != nil {
		fmt.Println(err)
		return errors.New("unable to write file")
	}

	encoder := json.NewEncoder(file)

	err = encoder.Encode(data)

	if err != nil {
		fmt.Println(err)
		return errors.New("unable to convert data to json")
	}

	file.Close()

	return nil
}
