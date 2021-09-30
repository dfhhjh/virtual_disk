package internal


import "fmt"

func (rd Rd) CommandExecute(vd *VirtualDisk, typ int, para string, path string, commandmanage CommandManage) {
	abspath:= commandmanage.ConvertRelaivePathToAbsolutePathFile(vd, path)
	var delcomponent = &Component{}
	pathelem := SplitPath(abspath)
	switch typ{
	case 0:
		if para == ""{
			delcomponent.RemoveFolder(vd , pathelem)
		}else if para == "s"{
			delcomponent.RemoveFolderS(vd , pathelem)
		}
	case 1:
		fmt.Println(WildCardError())
		delcomponent = nil

	case 2:
		fmt.Println(TrueDiskError())
		delcomponent = nil

	case 3:
		fmt.Println(TrueDiskError())
		delcomponent = nil
	}
	vd.Restart()
}
