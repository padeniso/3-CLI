package file

import (
	"fmt"
	"os"
	"strings"
)

func ReadFromFile(file string) (data []byte, err error) {
	isJSON := strings.HasSuffix(file, ".json")
	if !isJSON {
		fmt.Printf("The file %s is not a JSON file", file)
		return nil, err
	}
	data, err = os.ReadFile(file)
	if err != nil {
		fmt.Println("Unable to read from a file")
		return nil, err
	}
	return data, nil
}

func WriteToFile(content []byte, name string) error {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("Unable to create a file")
		return err
	}
	_, err = file.Write(content)
	if err != nil {
		fmt.Println("Unable to write to the file")
		return err
	}
	defer file.Close()
	fmt.Println("File written successfully")
	return nil
}
