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
	abspath := commandmanage.ConvertRelaivePathToAbsolutePathFile(vd, path)
	pathelem := SplitPath(abspath)
	switch typ{
	case 0:
		 renfolder.RenComponent(vd , pathelem, name)
	case 1:
		fmt.Println(WildCardError())
		renfolder = nil

	case 2:
		fmt.Println(TrueDiskError())
		renfolder = nil
	case 3:
		fmt.Println(TrueDiskError())
		renfolder = nil
	}
	vd.Restart()
}


