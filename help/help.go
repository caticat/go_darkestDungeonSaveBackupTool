package help

import (
	"os"
	"path/filepath"
	"strings"
)

// 格式化路径
// 统一路径分隔符
func FormatPath(path string) string {
	return strings.ReplaceAll(path, "/", "\\")
}

// 获取目录下的文件夹
func GetDirFolder(dir string) ([]string, error) {
	sliFolder := make([]string, 0, 1)
	err := filepath.Walk(dir, func(p string, info os.FileInfo, err error) error {
		if dir == p {
			return nil
		}
		staP, err := os.Stat(p)
		if err != nil {
			return err
		}
		if !staP.IsDir() {
			return nil
		}
		if strings.LastIndex(strings.TrimPrefix(FormatPath(p), dir), "\\") > 0 {
			return nil
		}
		sliFolder = append(sliFolder, p)
		return nil
	})

	return sliFolder, err
}
