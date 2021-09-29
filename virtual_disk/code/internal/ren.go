package internal

import (
	"fmt"
	"strings"
)

func (ren Ren) CommandExecute(vd *VirtualDisk, commandmanage CommandManage, typ int,path string) {
	var renfolder = &Component{}
	pathlist := strings.Split(path," ")
	path = pathlist[0]
    name := pathlist[1]
	abspath := commandmanage.ConvertRelaivePathToAbsolutePath(vd, path)
	pathelem := SplitPath(abspath)
	switch typ{
	case 0:
		 renfolder.RenComponent(vd , pathelem, name)
		vd.UpdateCurrentFolder(&vd.RootComponent)
		OutputRootDrive()
		vd.Execute()
	case 1:
		fmt.Println(WildCardError())
		renfolder = nil
		OutputRootDrive()
		vd.Execute()
	case 2:
		fmt.Println(TrueDiskError())
		renfolder = nil
		OutputRootDrive()
		vd.Execute()
	case 3:
		fmt.Println(TrueDiskError())
		renfolder = nil
		OutputRootDrive()
		vd.Execute()
	}

}


