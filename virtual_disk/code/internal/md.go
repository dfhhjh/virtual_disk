package internal

import "fmt"

func (md Md) CommandExecute(vd *VirtualDisk, typ int, abspath string) {
	var addfolder = &Component{}
	pathelem := SplitPath(abspath)
	switch typ{
	case 0:
		addfolder.AddFolder(vd , pathelem)
		vd.UpdateCurrentFolder(&vd.RootComponent)
		OutputRootDrive()
		vd.Execute()
	case 1:
		fmt.Println(WildCardError())
		addfolder = nil
		OutputRootDrive()
		vd.Execute()
	case 2:
		fmt.Println(TrueDiskError())
		addfolder = nil
		OutputRootDrive()
		vd.Execute()
	case 3:
		fmt.Println(TrueDiskError())
		addfolder = nil
		OutputRootDrive()
		vd.Execute()
	}

}


