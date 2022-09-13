package main

import (
	"fmt"

	"github.com/Tahler/golang-read-relative-file-proof/tables"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	for _, name := range []string{"samples", "slides"} {
		b, err := tables.ReadTable(name)
		if err != nil {
			return err
		}
		fmt.Printf("Read %s as JSON: %s\n", name, string(b))
	}
	return nil
}
