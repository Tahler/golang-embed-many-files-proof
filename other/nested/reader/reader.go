package reader

import (
	"github.com/Tahler/golang-read-relative-file-proof/tables"
)

func ReadTable() ([]byte, error) {
	return tables.ReadTable("samples")
}
