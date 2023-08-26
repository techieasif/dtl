package pkg

import (
	"os"
	"path/filepath"
)

func Save2Json(path string, file []byte) error {

	newFileName := filepath.Base(path)
	newFileName = newFileName[0:len(newFileName)-len(filepath.Ext(newFileName))] + ".json"
	r := filepath.Dir(path)

	err := os.WriteFile(filepath.Join(r, newFileName), file, os.FileMode(0644))

	if err != nil {
		return err
	}
	return nil
}
