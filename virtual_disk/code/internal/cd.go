package internal

import "fmt"

func (cd Cd) CommandExecute(vd *VirtualDisk, typ int, abspath string) {
	var changefolder = &Component{}
	pathelem := SplitPath(abspath)
	switch typ{
	case 0:
		changefolder = changefolder.ChangeNode(vd, pathelem)
		fmt.Print(changefolder.Path)
		fmt.Print(GetSeparatorChar())
		vd.Execute()
	case 1:
		fmt.Println(WildCardError())
		vd.Execute()
	case 2:
		fmt.Println(TrueDiskError())
		vd.Execute()
	case 3:
		fmt.Println(TrueDiskError())
		vd.Execute()
	}
}