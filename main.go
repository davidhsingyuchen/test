package main

import (
	"fmt"

	"github.com/onsi/gomega"
	"github.com/spf13/afero"
)

func main() {
	err := afero.ErrFileClosed
	fmt.Println(err)
	var a gomega.Assertion
	fmt.Println(a)
}
