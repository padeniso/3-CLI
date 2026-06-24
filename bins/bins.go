package bins

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

// NewBin creates new Bin element
func NewBin(id string, private bool, createdAt time.Time, name string) (*Bin, error) {
	if id == "" || name == "" {
		return nil, errors.New("Incorrect id or name field")
	}
	mynewBin := Bin{id: id, private: private, createdAt: time.Now(), name: name}
	return &mynewBin, nil

}
