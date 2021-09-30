package internal

import (
	"fmt"
	"strings"
)

func (move Move) CommandExecute(vd *VirtualDisk, commandmanage CommandManage, typ int, para string, path string) {
	var addcomponent = &Component{}
	var removecomponent = &Component{}
	var repeatcomponent = &Component{}
	pathlist := strings.Split(path," ")
	pathone := pathlist[0]
	pathtwo := pathlist[1]
	abspathone := commandmanage.ConvertRelaivePathToAbsolutePathFile(vd, pathone)
	abspathtwo := commandmanage.ConvertRelaivePathToAbsolutePathFile(vd, pathtwo)
	pathelemone := SplitPath(abspathone)
	pathelemtwo := SplitPath(abspathtwo)
	switch typ{
	case 0:
		if para == ""{
			removecomponent = removecomponent.ChangeNode(vd, pathelemone)
			if removecomponent.IsFolder == true{
				addcomponent.AddFolderTwo(vd, pathelemtwo, removecomponent)
			}else{
				addcomponent.AddFile(vd,pathelemtwo, removecomponent)
			}
			removecomponent.RemoveComponent(vd, pathelemone)
		}else if para == "y"{
			repeatcomponent = removecomponent.ChangeNode(vd, pathelemone)
			if repeatcomponent .IsFolder == true{
				addcomponent.AddFolderTwoY(vd, pathelemtwo, repeatcomponent)
			}else{
				addcomponent.AddFileY(vd,pathelemtwo, repeatcomponent)
			}
			removecomponent.RemoveComponentY(vd, pathelemone)
		}

	case 1:
		fmt.Println(WildCardError())
		addcomponent = nil
	case 2:
		fmt.Println(TrueDiskError())
		addcomponent = nil
	case 3:
		fmt.Println(TrueDiskError())
		addcomponent = nil
	}
	vd.Restart()
}



