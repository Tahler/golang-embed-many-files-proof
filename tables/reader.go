package tables

import (
	"embed"

	"github.com/ghodss/yaml"
)

//go:embed *.yaml
var tablesFS embed.FS

// ReadTable returns {tableName}.yaml as JSON.
func ReadTable(tableName string) ([]byte, error) {
	path := tableName + ".yaml"
	yamlContent, err := tablesFS.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return yaml.YAMLToJSON(yamlContent)
}
