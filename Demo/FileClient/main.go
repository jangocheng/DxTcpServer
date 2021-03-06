package main

import (
	"github.com/suiyunonghen/DxTcpServer/Demo/FileClient/DxFileClient"
	"fmt"
	"github.com/suiyunonghen/GVCL/Components/Controls"
	"github.com/suiyunonghen/GVCL/Components/NVisbleControls"
)

var (
	fclient *FileClient.FileClient
)

func main() {
	app := controls.NewApplication()
	app.ShowMainForm = false
	mainForm := app.CreateForm()
	PopMenu := NVisbleControls.NewPopupMenu(mainForm)

	fclient = FileClient.NewFileClient()
	fclient.OnDownLoad = func(client *FileClient.FileClient, FileName string, TotalSize, Position int64) {
		fmt.Println(FileName, "正在文件下载", Position*100/TotalSize, "%")
	}
	mItem := PopMenu.Items().AddItem("下载文件")
	mItem.OnClick = func(sender interface{}) {
		fclient.DownLoadFile("YNote.exe", "d:\\YNote.exe")
		fclient.DownLoadFile("cloudmusicsetup_2_2_1[192801].exe", "d:\\cloudmusicsetup_2_2_1[192801].exe")
	}
	mItem = PopMenu.Items().AddItem("-")
	mItem = PopMenu.Items().AddItem("退出")
	mItem.OnClick = func(sender interface{}) {
		fclient.Close()
		mainForm.Close()
	}
	trayicon := NVisbleControls.NewTrayIcon(mainForm)
	trayicon.PopupMenu = PopMenu
	trayicon.SetVisible(true)
	fclient.Connect("127.0.0.1:8340")
	app.Run()
}
