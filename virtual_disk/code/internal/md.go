package internal

import "fmt"

func (md Md) CommandExecute(vd *VirtualDisk, typ int, path string, commandmanage CommandManage) {
	abspath := commandmanage.ConvertRelaivePathToAbsolutePathFile(vd, path)
	var addfolder = &Component{}
	pathelem := SplitPath(abspath)
	switch typ{
	case 0:
		addfolder.AddFolder(vd , pathelem)

	case 1:
		fmt.Println(WildCardError())
		addfolder = nil

	case 2:
		fmt.Println(TrueDiskError())
		addfolder = nil
	case 3:
		fmt.Println(TrueDiskError())
		addfolder = nil
	}
	vd.Restart()
}


