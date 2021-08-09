package main

import (
	"fmt"
	"github.com/caticat/go_save_backup/help"
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
		Print("开始备份:", g_ptrConfig.SavePath, "->", tar)
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
			Print("备份当前存档到:", pathRecycle)
			help.Rm(pathRecycle)
			help.Cp(g_ptrConfig.SavePath, pathRecycle)
		} else {
			Print("获取回收站路径失败:", err)
		}

		// 还原
		Print("开始还原:", pathRestore, "->", g_ptrConfig.SavePath)
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
		Print("开始清理路径:", pathRm)
		if pathKeep, err := g_ptrConfig.GetNewestBackup(); err == nil {
			Print("保留文件夹:", pathKeep)
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
		Title:  "备份",
		Size:   Size{700, 300},
		Layout: VBox{},
		Children: []Widget{
			PushButton{
				Text: "备份",
				OnClicked: func() {
					if err := backup(); err == nil {
						Print("备份成功")
					} else {
						PrintError(err)
					}
				},
			},
			PushButton{
				Text: "还原",
				OnClicked: func() {
					if err := restore(); err == nil {
						Print("还原完成")
					} else {
						PrintError(err)
					}
				},
			},
			PushButton{
				Text: "清理",
				OnClicked: func() {
					if err := clear(); err == nil {
						Print("清理成功")
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

	// 备份
	//if err := backup(); err == nil {
	//	log.Println("备份成功")
	//} else {
	//	log.Fatal(err)
	//}

	// 还原
	//if err := restore(); err == nil {
	//	log.Println("还原完成")
	//} else {
	//	log.Fatal(err)
	//}

	// 清理
	//if err := clear(); err == nil {
	//	log.Println("清理成功")
	//} else {
	//	log.Fatal(err)
	//}
}
