package storage

import (
	"demo/cli/bins"
	"demo/cli/file"
	"encoding/json"
	"fmt"
)

func JsonToFile(mybin *bins.Bin, filename string) error {
	fileContent, err := json.Marshal(&mybin)
	if err != nil {
		return err
	}
	err = file.WriteToFile(fileContent, filename)
	return err
}

func ReadBinFromFile(mybin *bins.Bin, filename string) (*bins.Bin, error) {
	var myBin bins.Bin
	fileContent, err := file.ReadFromFile(filename)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(fileContent, &myBin)
	if err != nil {
		fmt.Println("Не удалось разобрать файл data.json")
		return nil, err
	}
	return &myBin, nil
}
