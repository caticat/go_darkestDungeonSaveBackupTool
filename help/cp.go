package help

import (
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// 强制覆盖复制
func Cp(src, tar string) error {
	src = FormatPath(src)
	tar = FormatPath(tar)
	staSrc, err := os.Stat(src)
	if err != nil {
		return err
	}

	if staSrc.IsDir() {
		filepath.Walk(src, func(srcFile string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			staPath, err := os.Stat(srcFile)
			if err != nil {
				return err
			}

			tarPath := getTarPath(src, tar, srcFile)
			if staPath.IsDir() {
				//log.Println("复制文件夹:", srcFile, ",到:", tarPath)
				err := os.MkdirAll(tarPath, 0777)
				if err != nil {
					return err
				}
			} else {
				err := os.MkdirAll(path.Base(tarPath), 0777)
				if err != nil {
					return err
				}
				return cpFile(srcFile, tarPath)
			}
			return nil
		})
		return nil
	} else {
		return cpFile(src, tar)
	}
}

// ============================ 内部函数

// 相对路径计算
func getTarPath(src, tar, obj string) string {
	rel := strings.TrimPrefix(obj, src)
	return path.Join(tar, rel)
}

// 单个文件复制
func cpFile(src, tar string) error {
	//log.Println("复制文件:", src, ",到:", tar)
	filSrc, err := os.Open(src)
	if err != nil {
		return err
	}
	defer filSrc.Close()

	filTar, err := os.Create(tar)
	if err != nil {
		return err
	}
	defer filTar.Close()

	_, err = io.Copy(filTar, filSrc)
	if err != nil {
		return err
	}

	return nil
}
