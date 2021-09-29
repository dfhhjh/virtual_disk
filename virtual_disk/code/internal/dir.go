package internal

import (
	"fmt"
	"strings"
)

func (dir Dir) CommandExecute(vd *VirtualDisk, commandmanage CommandManage, typ int ,para string, singlepath bool, havepara bool, path string) {
	var node = &Component{}
	switch typ{
	case 0:
		if singlepath && !havepara{
			abspath := commandmanage.ConvertRelaivePathToAbsolutePath(vd, path)
			pathelem := SplitPath(abspath)
			node = node.ChangeNode(vd,pathelem)
			if vd.CurrentFolder != &vd.RootComponent{
				node.DirCurrentAndFatherFolder()
			}
			node.TraverseDir(vd)
		}else if !singlepath && !havepara{
			pathlist := strings.Split(path," ")
			pathone := pathlist[0]
			pathtwo := pathlist[1]
			abspathone := commandmanage.ConvertRelaivePathToAbsolutePath(vd, pathone)
			abspathtwo := commandmanage.ConvertRelaivePathToAbsolutePath(vd, pathtwo)
			pathelemone := SplitPath(abspathone)
			pathelemtwo := SplitPath(abspathtwo)
			node = node.ChangeNode(vd, pathelemone)
			if vd.CurrentFolder != &vd.RootComponent{
				node.DirCurrentAndFatherFolder()
			}
			node.TraverseDir(vd)
			node = node.ChangeNode(vd, pathelemtwo)
			if vd.CurrentFolder != &vd.RootComponent{
				node.DirCurrentAndFatherFolder()
			}
			node.TraverseDir(vd)
		}else if singlepath && havepara{
			abspath := commandmanage.ConvertRelaivePathToAbsolutePath(vd, path)
			pathelem := SplitPath(abspath)
			node = node.ChangeNode(vd,pathelem)
			if vd.CurrentFolder != &vd.RootComponent{
				node.DirCurrentAndFatherFolder()
			}
			if para == s{
				node.TraverseDirS(vd)
			}else if para == ad{
				node.TraverseDirAd(vd)
			}else if para == "ads"{
				node.TraverseDirAd(vd)
				node = node.ChangeNode(vd,pathelem)
				node.TraverseDirS(vd)
			}
		}else{
			pathlist := strings.Split(path," ")
			pathone := pathlist[0]
			pathtwo := pathlist[1]
			abspathone := commandmanage.ConvertRelaivePathToAbsolutePath(vd, pathone)
			abspathtwo := commandmanage.ConvertRelaivePathToAbsolutePath(vd, pathtwo)
			pathelemone := SplitPath(abspathone)
			pathelemtwo := SplitPath(abspathtwo)
			node = node.ChangeNode(vd, pathelemone)
			if vd.CurrentFolder != &vd.RootComponent{
				node.DirCurrentAndFatherFolder()
			}
			if para == s{
				node.TraverseDirS(vd)
			}else if para == ad{
				node.TraverseDirAd(vd)
			}else if para == "ads"{
				node.TraverseDirAd(vd)
				node = node.ChangeNode(vd,pathelemone)
				node.TraverseDirS(vd)
			}
			node = node.ChangeNode(vd, pathelemtwo)
			if vd.CurrentFolder != &vd.RootComponent{
				node.DirCurrentAndFatherFolder()
			}
			if para == s{
				node.TraverseDirS(vd)
			}else if para == ad{
				node.TraverseDirAd(vd)
			}else if para == "ads"{
				node.TraverseDirAd(vd)
				node = node.ChangeNode(vd,pathelemtwo)
				node.TraverseDirS(vd)
			}
		}
		vd.UpdateCurrentFolder(&vd.RootComponent)
		OutputRootDrive()
		vd.Execute()
	case 1:
		if singlepath && !havepara{
			pathelem := SplitPath(path)
			str := pathelem[len(pathelem)-1]
			pathelem = pathelem[:len(pathelem)-1]
			path = strings.Join(pathelem, GetSeparatorChar())
			abspath := commandmanage.ConvertRelaivePathToAbsolutePath(vd, path)
			pathelem = SplitPath(abspath)
			node = node.ChangeNode(vd,pathelem)
			if vd.CurrentFolder != &vd.RootComponent{
				node.DirCurrentAndFatherFolder()
			}
			node.HwdTraverseDir(vd,str)
		}else if !singlepath && !havepara{
			pathlist := strings.Split(path," ")
			pathone := pathlist[0]
			pathelemone := SplitPath(pathone)
			strone := pathelemone[len(pathelemone)-1]
			pathelemone = pathelemone[:len(pathelemone)-1]
			pathone = strings.Join(pathelemone , GetSeparatorChar())
			abspathone := commandmanage.ConvertRelaivePathToAbsolutePath(vd, pathone)
			pathelemone = SplitPath(abspathone)
			pathtwo := pathlist[1]
			pathelemtwo := SplitPath(pathtwo)
			strtwo := pathelemtwo[len(pathelemtwo)-1]
			pathelemtwo = pathelemtwo[:len(pathelemtwo)-1]
			pathtwo = strings.Join(pathelemtwo , GetSeparatorChar())
			abspathtwo := commandmanage.ConvertRelaivePathToAbsolutePath(vd, pathtwo)
			pathelemtwo = SplitPath(abspathtwo)
			node = node.ChangeNode(vd,pathelemone)
			if vd.CurrentFolder != &vd.RootComponent{
				node.DirCurrentAndFatherFolder()
			}
			node.HwdTraverseDir(vd,strone)
			node = node.ChangeNode(vd,pathelemtwo)
			if vd.CurrentFolder != &vd.RootComponent{
				node.DirCurrentAndFatherFolder()
			}
			node.HwdTraverseDir(vd,strtwo)
		}else if singlepath && havepara{
			pathelem := SplitPath(path)
			str := pathelem[len(pathelem)-1]
			pathelem = pathelem[:len(pathelem)-1]
			path = strings.Join(pathelem, GetSeparatorChar())
			abspath := commandmanage.ConvertRelaivePathToAbsolutePath(vd, path)
			pathelem = SplitPath(abspath)
			node = node.ChangeNode(vd,pathelem)
			if vd.CurrentFolder != &vd.RootComponent{
				node.DirCurrentAndFatherFolder()
			}
			if para == s{
				node.HwdTraverseDirS(vd, str)
			}else if para == ad{
				node.HwdTraverseDirAd(vd, str)
			}else if para == "ads"{
				node.HwdTraverseDirAd(vd, str)
				node = node.ChangeNode(vd,pathelem)
				node.HwdTraverseDirS(vd, str)
			}
		}else{
			pathlist := strings.Split(path," ")
			pathone := pathlist[0]
			pathelemone := SplitPath(pathone)
			strone := pathelemone[len(pathelemone)-1]
			pathelemone = pathelemone[:len(pathelemone)-1]
			pathone = strings.Join(pathelemone , GetSeparatorChar())
			abspathone := commandmanage.ConvertRelaivePathToAbsolutePath(vd, pathone)
			pathelemone = SplitPath(abspathone)
			pathtwo := pathlist[1]
			pathelemtwo := SplitPath(pathtwo)
			strtwo := pathelemtwo[len(pathelemtwo)-1]
			pathelemtwo = pathelemtwo[:len(pathelemtwo)-1]
			pathtwo = strings.Join(pathelemtwo , GetSeparatorChar())
			abspathtwo := commandmanage.ConvertRelaivePathToAbsolutePath(vd, pathtwo)
			pathelemtwo = SplitPath(abspathtwo)
			node = node.ChangeNode(vd, pathelemone)
			if vd.CurrentFolder != &vd.RootComponent{
				node.DirCurrentAndFatherFolder()
			}
			if para == s{
				node.HwdTraverseDirS(vd, strone)
			}else if para == ad{
				node.HwdTraverseDirAd(vd, strone)
			}else if para == "ads"{
				node.HwdTraverseDirAd(vd, strone)
				node = node.ChangeNode(vd,pathelemone)
				node.HwdTraverseDirS(vd, strone)
			}
			node = node.ChangeNode(vd, pathelemtwo)
			if vd.CurrentFolder != &vd.RootComponent{
				node.DirCurrentAndFatherFolder()
			}
			if para == s{
				node.HwdTraverseDirS(vd, strtwo)
			}else if para == ad{
				node.HwdTraverseDirAd(vd, strtwo)
			}else if para == "ads"{
				node.HwdTraverseDirAd(vd, strtwo)
				node = node.ChangeNode(vd,pathelemtwo)
				node.HwdTraverseDirS(vd, strtwo)
			}
		}
		vd.UpdateCurrentFolder(&vd.RootComponent)
		OutputRootDrive()
		vd.Execute()
	case 2:
		fmt.Println(TrueDiskError())
		node = nil
		OutputRootDrive()
		vd.Execute()
	case 3:
		fmt.Println(TrueDiskError())
		node = nil
		OutputRootDrive()
		vd.Execute()
	}

}