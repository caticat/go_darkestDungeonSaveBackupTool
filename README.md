# 黑暗地牢存档备份工具

[English](./README_ENG.md)

## 功能

快速的备份还原黑暗地牢的存档
PS: 还可以备份其他文件夹, 功能上没有限制

## 功能

- 备份指定文件夹到备份文件夹
- 备份会被按照备份时间重命名,同时可以有多个时间点的备份
- 还原时,直接使用最新的备份进行还原
- 还原前,自动备份回进行覆盖的目录,防止误操作(备份文件放在recycle目录中)
- 清理功能可以快速清理除了最新备份以外的所有其他冗余备份

## 使用方法

- 修改`config.yaml`中的`savePath`,改成自己存档的对应目录(基本上只要修改`账户ID`为自己的ID即可)
- 运行`DarkestDungeonSaveBackupTool.exe`
- **还原时需要先退出游戏**,其他操作无需关闭游戏

## 项目

[地址](https://github.com/caticat/go_darkestDungeonSaveBackupTool/releases)
