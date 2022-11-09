package main

import (
	"fmt"

	"github.com/spf13/afero"
)

func main() {
	err := afero.ErrFileClosed
	fmt.Println(err)
}
