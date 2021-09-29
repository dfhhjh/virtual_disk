package internal

import "fmt"

func (del Del) CommandExecute(vd *VirtualDisk, typ int, para string, abspath string) {
	var delcomponent = &Component{}
	pathelem := SplitPath(abspath)
	switch typ{
	case 0:
		if para == ""{
			delcomponent.RemoveComponent(vd , pathelem)
		}else if para == "s"{
			delcomponent.RemoveComponentS(vd , pathelem)
		}
		vd.UpdateCurrentFolder(&vd.RootComponent)
		OutputRootDrive()
		vd.Execute()
	case 1:
		fmt.Println(WildCardError())
		delcomponent = nil
		OutputRootDrive()
		vd.Execute()
	case 2:
		fmt.Println(TrueDiskError())
		delcomponent = nil
		OutputRootDrive()
		vd.Execute()
	case 3:
		fmt.Println(TrueDiskError())
		delcomponent = nil
		OutputRootDrive()
		vd.Execute()
	}

}

