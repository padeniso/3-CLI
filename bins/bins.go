package bins

import (
	"errors"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
}

// NewBin creates new Bin element
func NewBin(id string, private bool, createdAt time.Time, name string) (*Bin, error) {
	if id == "" || name == "" {
		return nil, errors.New("Incorrect id or name field")
	}
	mynewBin := Bin{id: id, private: private, createdAt: time.Now(), name: name}
	return &mynewBin, nil

}
