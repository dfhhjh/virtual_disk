package internal

import (
	"fmt"
)

type VirtualDisk struct {
	CommandManageSystem CommandManage
	RootComponent       Component
	CurrentFolder       *Component
}//虚拟磁盘类

func (vd *VirtualDisk) CreateVirtualDisk() {
	vd.VirtualDiskInit()
}

func (vd *VirtualDisk)  VirtualDiskInit() {
	vd.RootComponent.Name = GetRootDrive()
	vd.RootComponent.FatherComponent = nil
	vd.RootComponent.IsFolder = true
	vd.RootComponent.Path = vd.RootComponent.Name
	vd.CurrentFolder = &vd.RootComponent
	vd.RootComponent.FatherComponent = &vd.RootComponent
}

func (vd *VirtualDisk) Execute() {
	var commandmanage CommandManage
	inputstring := commandmanage.CreateCommand()
	sizeerror := SizeError()
	invalidcharaerror := InvalidCharactersError()
	if inputstring == "\n"{
		vd.Restart()
	}
	if !commandmanage.CheckSize(inputstring) { //检查输入字符大小，输入过大则重新输入
		fmt.Println(sizeerror)
		vd.Restart()
	}
	if commandmanage.HaveInvalidCharacters(inputstring) { //检查输入有无无效字符，有则重新输入
		fmt.Println(invalidcharaerror)
		vd.Restart()
	}
	inputstring = commandmanage.CheckSeparator(inputstring)
	itd := commandmanage.IsTrueDiskPath(inputstring)
	hwc := commandmanage.HaveWildCard(inputstring)
	typ := commandmanage.GetCommandType(itd, hwc)
	parapath, cmd := commandmanage.SeparateCommand(inputstring, vd)
	path, para := commandmanage.SeparateParameter(parapath, cmd)
	singlepath := commandmanage.IsSinglePath(path)
	havepara := commandmanage.HaveParameter(para)
	switch cmd {
	case dir:
		var dir Dir
		dir.CommandExecute(vd, commandmanage, typ, para, singlepath, havepara, path)
	case md:
		var md Md
		md.CommandExecute(vd, typ, path, commandmanage)
	case cd:
		var cd Cd
		cd.CommandExecute(vd, typ, path, commandmanage)
	case copy:
		var copy Copy
		copy.CommandExecute(vd, commandmanage,typ, path)
	case touch:
		var touch Touch
		touch.CommandExecute(vd, typ, path, commandmanage)
	case del:
		var del Del
		del.CommandExecute(vd, typ, para, path, commandmanage)
	case rd:
		var rd Rd
		rd.CommandExecute(vd, typ, para, path, commandmanage)
	case ren:
		var ren Ren
		ren.CommandExecute(vd, commandmanage, typ, path)
	case move:
		var move Move
		move.CommandExecute(vd, commandmanage, typ, para, path)
	case save:
		var save Save
		save.CommandExecute(vd, path)
	case load:
		var load Load
		load.CommandExecute(vd, path)
	case cls:
		var cls Cls
		cls.CommandExecute(vd)
	}
}


func (vd *VirtualDisk) GetChildNodeByName(startnode *Component, name string ) {
	if startnode.SonComponent == nil &&startnode.Name == GetRootDrive(){
		vd.CurrentFolder = &vd.RootComponent
	}else{
		for _,v := range startnode.SonComponent {
			if IsExits(v.Name, name) {
				vd.CurrentFolder = v
				break
			}
			if (v.IsFolder) {
				vd.GetChildNodeByName(v, name)
			}
		}
	}
}

func (vd VirtualDisk) GetCurrentFolderPath() string{
	return vd.CurrentFolder.Path
}

func (vd *VirtualDisk) UpdateCurrentFolder(folder *Component){
	vd.CurrentFolder = folder
}

func (vd *VirtualDisk) Restart(){
	vd.UpdateCurrentFolder(&vd.RootComponent)
	OutputRootDrive()
	vd.Execute()
}