package help

import "os"

func Rm(path string) error {
	path = FormatPath(path)
	return os.RemoveAll(path)
}
