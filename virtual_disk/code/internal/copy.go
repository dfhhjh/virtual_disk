package internal

import (
	"strings"
)

func (copy Copy) CommandExecute(vd *VirtualDisk, commandmanage CommandManage,typ int, path string) {
	var addcomponent = &Component{}
	var copycomponent = &Component{}
	pathlist := strings.Split(path, " ")
	pathone := pathlist[0]
	pathtwo := pathlist[1]
	abspathtwo := commandmanage.ConvertRelaivePathToAbsolutePathFile(vd, pathtwo)
	pathelemtwo := SplitPath(abspathtwo)
	switch typ {
	case 0:
		abspathone := commandmanage.ConvertRelaivePathToAbsolutePathFile(vd, pathone)
		pathelemone := SplitPath(abspathone)
		copycomponent = copycomponent.ChangeNode(vd, pathelemone)
		vd.UpdateCurrentFolder(copycomponent.FatherComponent)
		if copycomponent.IsFolder == false {
			addcomponent.CopyFile(vd, pathelemtwo, copycomponent)
		}
	case 1:
		pathelemone := SplitPath(pathone)
		strone := pathelemone[len(pathelemone)-1]
		pathelemone = pathelemone[:len(pathelemone)-1]
		path = strings.Join(pathelemone, GetSeparatorChar())
		abspath := commandmanage.ConvertRelaivePathToAbsolutePathFile(vd, path)
		pathelemone = SplitPath(abspath)
		copycomponent = copycomponent.ChangeNode(vd, pathelemone)
		copycomponent.HwdAddFile(vd, pathelemtwo, strone)
	case 2:
		if strings.HasPrefix(pathone, "@") {
			pathone = strings.TrimPrefix(pathone, "@")
		}
		addcomponent.TrueDiskAddFile(vd, pathone, pathelemtwo)
	case 3:
		if strings.HasPrefix(pathone, "@") {
			pathone = strings.TrimPrefix(pathone, "@")
		}
		pathelemone := SplitPath(pathone)
		lastone := pathelemone[len(pathelemone)-1]
		pathelemone = pathelemone[:len(pathelemone)-1]
		pathone = strings.Join(pathelemone, GetSeparatorChar())
		addcomponent.HwdTrueDiskAddFile(vd, lastone, pathone, pathelemtwo)
	}
	vd.Restart()
}