package main

import (
	"caticat.github.com/go_save_backup/help"
	"log"
)

var g_ptrConfig = NewConfig()

func backup() error {
	if tar, err := g_ptrConfig.GetBackupDir(); err == nil {
		log.Println("开始备份:", g_ptrConfig.SavePath, "->", tar)
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
			log.Println("备份当前存档到:", pathRecycle)
			help.Rm(pathRecycle)
			help.Cp(g_ptrConfig.SavePath, pathRecycle)
		} else {
			log.Println("获取回收站路径失败:", err)
		}

		// 还原
		log.Println("开始还原:", pathRestore, "->", g_ptrConfig.SavePath)
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
		log.Println("开始清理路径:", pathRm)
		if pathKeep, err := g_ptrConfig.GetNewestBackup(); err == nil {
			log.Println("保留文件夹:", pathKeep)
			if sliFolder, err := help.GetDirFolder(pathRm); err == nil {
				for _, folder := range sliFolder {
					if folder == pathKeep {
						continue
					}
					log.Println("[X]", folder)
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

func main() {
	if err := g_ptrConfig.Load("config.yaml"); err != nil {
		log.Fatal(err)
	}

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
