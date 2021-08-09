package main

import (
	"fmt"
	"github.com/caticat/go_darkestDungeonSaveBackupTool/help"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
	"time"
)

// 备注
// rsrc -manifest go_save_backup.manifest -ico ./res/394ca7c20fc9d944d31be8dd36f7108fd3d8f64f.ico -o rsrc.syso
// 关闭控制台
// go build -ldflags="-H windowsgui"

var g_ptrConfig = NewConfig()
var g_ptrOutTE *walk.TextEdit

func backup() error {
	if tar, err := g_ptrConfig.GetBackupDir(); err == nil {
		Print(g_ptrConfig.GetLanguage("backup_begin"), g_ptrConfig.SavePath, "->", tar)
		if err := help.Cp(g_ptrConfig.SavePath, tar); err == nil {
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}

func restore() error {
	if pathRestore, err := g_ptrConfig.GetNewestBackup(); err == nil {
		// 备份
		if pathRecycle, err := g_ptrConfig.GetRecyclePath(); err == nil {
			Print(g_ptrConfig.GetLanguage("backup_save_to"), pathRecycle)
			help.Rm(pathRecycle)
			help.Cp(g_ptrConfig.SavePath, pathRecycle)
		} else {
			Print(g_ptrConfig.GetLanguage("get_recycle_path_fail"), err)
		}

		// 还原
		Print(g_ptrConfig.GetLanguage("restore_begin"), pathRestore, "->", g_ptrConfig.SavePath)
		if err := help.Cp(pathRestore, g_ptrConfig.SavePath); err == nil {
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}

func clear() error {
	if pathRm, err := g_ptrConfig.GetBackupPath(); err == nil {
		Print(g_ptrConfig.GetLanguage("clear_begin"), g_ptrConfig.GetLanguage("path"), pathRm)
		if pathKeep, err := g_ptrConfig.GetNewestBackup(); err == nil {
			Print(g_ptrConfig.GetLanguage("keep_dir"), pathKeep)
			if sliFolder, err := help.GetDirFolder(pathRm); err == nil {
				for _, folder := range sliFolder {
					if folder == pathKeep {
						continue
					}
					Print("[X]", folder)
					help.Rm(folder)
				}
				return nil
			} else {
				return err
			}
		} else {
			return err
		}
	} else {
		return err
	}
}

func PrintError(sliLog ...interface{}) {
	var sliLogN []interface{}
	sliLogN = append(sliLogN, "[ERROR]", sliLog)
	Print(sliLogN...)
}

func Print(sliLog ...interface{}) {
	g_ptrOutTE.AppendText("[" + time.Now().Format("2006/01/02 15:04:05") + "]" + fmt.Sprint(sliLog...) + "\r\n")
}

func main() {
	if err := g_ptrConfig.Load("config.yaml"); err != nil {
		log.Fatal(err)
	}

	MainWindow{
		Title:  g_ptrConfig.GetLanguage("title"),
		Size:   Size{700, 300},
		Layout: VBox{},
		Children: []Widget{
			PushButton{
				Text: g_ptrConfig.GetLanguage("backup"),
				OnClicked: func() {
					if err := backup(); err == nil {
						Print(g_ptrConfig.GetLanguage("backup_done"))
					} else {
						PrintError(err)
					}
				},
			},
			PushButton{
				Text: g_ptrConfig.GetLanguage("restore"),
				OnClicked: func() {
					if err := restore(); err == nil {
						Print(g_ptrConfig.GetLanguage("restore_done"))
					} else {
						PrintError(err)
					}
				},
			},
			PushButton{
				Text: g_ptrConfig.GetLanguage("clear"),
				OnClicked: func() {
					if err := clear(); err == nil {
						Print(g_ptrConfig.GetLanguage("clear_done"))
					} else {
						PrintError(err)
					}
				},
			},
			TextEdit{
				AssignTo: &g_ptrOutTE,
				ReadOnly: true,
				VScroll:  true,
			},
		},
	}.Run()
}
