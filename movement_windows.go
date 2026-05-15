//go:build windows

package filange

import "os"

// Windows does not have a POSIX uid/gid ownership model, so
// we just return nil
func preserveOwnership(sourcePath, destPath string, fileInfo os.FileInfo) error {
	return nil
}
