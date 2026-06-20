package main

import (
	"errors"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

// newBin creates new Bin element
func newBin(id string, private bool, createdAt time.Time, name string) (*Bin, error) {
	if id == "" || name == "" {
		return nil, errors.New("Incorrect id or name field")
	}
	mynewBin := Bin{id: id, private: private, createdAt: time.Now(), name: name}
	return &mynewBin, nil

}

func main() {
	binList := []Bin{}

}
