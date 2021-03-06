package main

import (
	"errors"
	"github.com/caticat/go_darkestDungeonSaveBackupTool/help"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"
)

// 命名前缀
const SAVE_PREFIX string = "save_"

// 备份目录
const PATH_BACKUP string = "backup"

// 回收站
const PATH_RECYCLE string = "recycle"

type Config struct {
	SavePath    string            `yaml:"savePath"`
	MapLanguage map[string]string `yaml:"language"`
}

func NewConfig() *Config {
	return &Config{
		MapLanguage: make(map[string]string),
	}
}

func (this *Config) Load(path string) error {
	// 读取配置
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(f, this)
	if err != nil {
		return err
	}

	// 目录整理
	this.SavePath = this.fixPathLastIndex(this.SavePath)

	return nil
}

// 获取语言配置
func (this *Config) GetLanguage(key string) string {
	if this.MapLanguage == nil {
		PrintError(errors.New("m_mapLanguage == nil"))
		return key
	}

	if val, ok := this.MapLanguage[key]; ok {
		return val
	} else {
		PrintError(errors.New("no key:" + key + " found"))
		return key
	}
}

// 获取备份路径
func (this *Config) GetBackupPath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return help.FormatPath(path.Join(cwd, PATH_BACKUP)), nil
}

// 获取回收站路径
func (this *Config) GetRecyclePath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return help.FormatPath(path.Join(cwd, PATH_RECYCLE)), nil
}

// 获取备份文件夹名
func (this *Config) GetBackupDir() (string, error) {
	backupPath, err := this.GetBackupPath()
	if err != nil {
		return "", err
	}
	return help.FormatPath(path.Join(backupPath, this.genBackupDirName())), nil
}

// 获取最新的备份路径
func (this *Config) GetNewestBackup() (string, error) {
	backupPath, err := this.GetBackupPath()
	if err != nil {
		return "", err
	}

	sliPathDate := make([]uint64, 0, 1)
	pathPrefix := path.Join(backupPath, SAVE_PREFIX)
	pathPrefix = help.FormatPath(pathPrefix)
	if sliFolder, err := help.GetDirFolder(backupPath); err == nil {
		for _, folder := range sliFolder {
			strSuffix := strings.TrimPrefix(folder, pathPrefix)
			if i, err := strconv.ParseUint(strSuffix, 10, 64); err == nil {
				sliPathDate = append(sliPathDate, i)
			} else {
				Print("skip path:", folder) // 这里的报错只需要记录过掉的目录即可
			}
		}
	} else {
		return "", err
	}

	if len(sliPathDate) == 0 {
		return "", errors.New("find no backup file")
	}

	sort.Slice(sliPathDate, func(i, j int) bool {
		if sliPathDate[i] > sliPathDate[j] {
			return true
		} else if sliPathDate[i] == sliPathDate[j] {
			return i < j
		} else {
			return false
		}
	})

	return help.FormatPath(path.Join(backupPath, SAVE_PREFIX+strconv.FormatUint(sliPathDate[0], 10))), nil
}

// 生成备份文件夹名
func (this *Config) genBackupDirName() string {
	return SAVE_PREFIX + time.Now().Format("20060102150405")
}

// 去掉目录最后的分隔符
func (this *Config) fixPathLastIndex(p string) string {
	if len(p) == 0 {
		return ""
	}
	if (p[len(p)-1] == '/') || (p[len(p)-1] == '\\') {
		return p[0 : len(p)-1]
	}
	return p
}
