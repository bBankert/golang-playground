package saver

import "fmt"

type Saver interface {
	Save() error
}

func SaveData(data Saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Failed to save")
		return err
	}

	fmt.Println("Saved successfully")

	return nil
}
