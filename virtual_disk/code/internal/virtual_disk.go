package internal

import (
	"fmt"
)

type VirtualDiskManageSystem struct{

}//虚拟磁盘管理系统类

func (vdms VirtualDiskManageSystem) CreateVirtualDisk() {
	//VirtualDisk.VirtualDiskInit()
}

func (vdms VirtualDiskManageSystem) GetFatherContents(c Component) {
	//c.AddNode()
}

type VirtualDisk struct {
	CommandManageSystem CommandManage
	RootContents        Folder
	WorkingPath         Path
}//虚拟磁盘类

func (vd VirtualDisk)  VirtualDiskInit() {
	vd.RootContents.Name = GetRootDrive()
	vd.RootContents.FatherContents = nil
	vd.RootContents.IsFolder = true
	vd.WorkingPath.Name = vd.RootContents.Name
}

func (vd VirtualDisk) Execute() {
	fmt.Print("GetRootDrive()GetSeparatorChar()")
}



