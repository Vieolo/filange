package filange

import (
	"fmt"
	"os"
)

// Creates a directory if doesn't exist
func CreateDirIfNotExists(dir string, perm os.FileMode) error {
	if FileExists(dir) {
		return nil
	}

	if err := os.MkdirAll(dir, perm); err != nil {
		return fmt.Errorf("failed to create directory: %s: '%w'", dir, err)
	}

	return nil
}
