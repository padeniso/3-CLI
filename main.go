package main

import (
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func newBin(id string, private bool, createdAt time.Time, name string) *Bin {

}

func main() {
	binList := []Bin{}

}
