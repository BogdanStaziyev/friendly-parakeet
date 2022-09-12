package file_storage

import "fmt"

type Orog struct {
	Goro int64
}

func (t *Orog) try() {
	fmt.Sprintf("we haw: %d", Orog{43})
}
