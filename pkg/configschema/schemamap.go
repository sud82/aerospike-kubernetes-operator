package configschema

import (
	"embed"
	"io/fs"
	"path/filepath"
	"strings"
)

//go:embed json
var schemas embed.FS

type SchemaMap map[string]string

func NewSchemaMap() (SchemaMap, error) {
	schema := make(SchemaMap)

	if err := fs.WalkDir(
		schemas, ".", func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if !d.IsDir() {
				content, err := fs.ReadFile(schemas, path)
				if err != nil {
					return err
				}

				baseName := filepath.Base(path)
				key := strings.TrimSuffix(baseName, filepath.Ext(baseName))
				schema[key] = string(content)
			}

			return nil
		},
	); err != nil {
		return nil, err
	}

	return schema, nil
}
