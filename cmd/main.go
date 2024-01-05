package main

import (
	"os"

	"github.com/taylormonacelli/bloomtail"
)

func main() {
	code := bloomtail.Execute()
	os.Exit(code)
}
