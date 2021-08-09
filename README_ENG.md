# Darkest Dungeon Save Backup Tool

[中文](./README.md)

## Purpose

Use for backup save files for the game Darkest Dungeon
Aloth it can backup other files also.

## Util

- Backup target Dir to backup dir
- The backups has been renamed by date, so you can get multi backups at same time
- Fast restore the newest backup file to the target Dir
- Before restore, backup target Dir to recycle folder in case mistake operation
- Fast clear backups except the latest one

## Usage

- modify`savePath` to your own save path in `config.yaml`(basically change `账户ID` to your self steam ID will be OK)
- run`DarkestDungeonSaveBackupTool.exe`
- **Exit Game Before Restore**, backup & clear do not need close game
