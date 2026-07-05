package storage

import (
	"demo/cli/bins"
	"encoding/json"
	"fmt"
)

func JsonToFile(mybin *bins.Bin, file string) ([]byte, error) {
	fileContent, err := json.Marshal(&mybin)
	if err != nil {
		return nil, err
	}
	err = WriteToFile(fileContent, file)
}

func ReadBinFromFile(mybin *bins.Bin, file string) (*bins.Bin, error) {
	var myBin bins.Bin
	fileContent, err := ReadFromFile(file)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(fileContent, &myBin)
	if err != nil {
		fmt.Println("Не удалось разобрать файл data.json")
	}
	return &myBin, nil
}
