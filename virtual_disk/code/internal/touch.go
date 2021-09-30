package internal

import "fmt"

func (touch Touch) CommandExecute(vd *VirtualDisk, typ int, path string, commandmanage CommandManage) {
	abspath:= commandmanage.ConvertRelaivePathToAbsolutePathFile(vd, path)
	var addfile = &Component{}
	pathelem := SplitPath(abspath)
	switch typ{
	case 0:
		addfile.TouchFile(vd , pathelem)
	case 1:
		fmt.Println(WildCardError())
		addfile = nil
	case 2:
		fmt.Println(TrueDiskError())
		addfile = nil

	case 3:
		fmt.Println(TrueDiskError())
		addfile = nil
	}
	vd.Restart()
}
