package internal

import "fmt"

func (touch Touch) CommandExecute(vd *VirtualDisk, typ int,abspath string) {
	var addfile = &Component{}
	pathelem := SplitPath(abspath)
	switch typ{
	case 0:
		addfile.TouchFile(vd , pathelem)
		vd.UpdateCurrentFolder(&vd.RootComponent)
		OutputRootDrive()
		vd.Execute()
	case 1:
		fmt.Println(WildCardError())
		addfile = nil
		OutputRootDrive()
		vd.Execute()
	case 2:
		fmt.Println(TrueDiskError())
		addfile = nil
		OutputRootDrive()
		vd.Execute()
	case 3:
		fmt.Println(TrueDiskError())
		addfile = nil
		OutputRootDrive()
		vd.Execute()
	}

}
