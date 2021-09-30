package internal

import (
	"fmt"
	"strings"
)

func (dir Dir) CommandExecute(vd *VirtualDisk, commandmanage CommandManage, typ int ,para string, singlepath bool, havepara bool, path string) {
	var node = &Component{}
	node = vd.CurrentFolder
	path = commandmanage.DirConvert(vd, path)
	switch typ{
	case 0:
		if singlepath && !havepara{
			if path == ""{
				node.TraverseDir(vd)
			}else{
				pathelem := SplitPath(path)
				node = node.DirChangeNode(vd,pathelem)
				node.TraverseDir(vd)
			}
		}else if !singlepath && !havepara{
			if path == ""{
				node.TraverseDir(vd)
			}else{
				pathlist := strings.Split(path," ")
				pathone := pathlist[0]
				pathtwo := pathlist[1]
				abspathone := commandmanage.DirConvert(vd, pathone)
				abspathtwo := commandmanage.DirConvert(vd, pathtwo)
				pathelemone := SplitPath(abspathone)
				pathelemtwo := SplitPath(abspathtwo)
				node = node.DirChangeNode(vd, pathelemone)
				node.TraverseDir(vd)
				node = node.DirChangeNode(vd, pathelemtwo)
				node.TraverseDir(vd)
			}
		}else if singlepath && havepara{
			if path == ""{
				if para == s{
					node.TraverseDirS(vd)
				}else if para == ad{
					node.TraverseDirAd(vd)
				}else if para == "ads"{
					node.TraverseDirAd(vd)
					node.TraverseDirS(vd)
				}
			}else{
				abspath := commandmanage.DirConvert(vd, path)
				pathelem := SplitPath(abspath)
				node = node.DirChangeNode(vd,pathelem)
				if para == s{
					node.TraverseDirS(vd)
				}else if para == ad{
					node.TraverseDirAd(vd)
				}else if para == "ads"{
					node.TraverseDirAd(vd)
					node = node.DirChangeNode(vd,pathelem)
					node.TraverseDirS(vd)
				}
			}
		}else{
			if path == ""{
				if para == s{
					node.TraverseDirS(vd)
				}else if para == ad{
					node.TraverseDirAd(vd)
				}else if para == "ads"{
					node.TraverseDirAd(vd)
					node.TraverseDirS(vd)
				}
			}else{
				pathlist := strings.Split(path," ")
				pathone := pathlist[0]
				pathtwo := pathlist[1]
				abspathone := commandmanage.DirConvert(vd, pathone)
				abspathtwo := commandmanage.DirConvert(vd, pathtwo)
				pathelemone := SplitPath(abspathone)
				pathelemtwo := SplitPath(abspathtwo)
				node = node.DirChangeNode(vd, pathelemone)
				if para == s{
					node.TraverseDirS(vd)
				}else if para == ad{
					node.TraverseDirAd(vd)
				}else if para == "ads"{
					node.TraverseDirAd(vd)
					node = node.DirChangeNode(vd,pathelemone)
					node.TraverseDirS(vd)
				}
				node = node.DirChangeNode(vd, pathelemtwo)
				if para == s{
					node.TraverseDirS(vd)
				}else if para == ad{
					node.TraverseDirAd(vd)
				}else if para == "ads"{
					node.TraverseDirAd(vd)
					node = node.DirChangeNode(vd,pathelemtwo)
					node.TraverseDirS(vd)
				}
			}

		}
	case 1:
		if singlepath && !havepara{
			pathelem := SplitPath(path)
			str := pathelem[len(pathelem)-1]
			pathelem = pathelem[:len(pathelem)-1]
			path = strings.Join(pathelem, GetSeparatorChar())
			abspath := commandmanage.DirConvert(vd, path)
			pathelem = SplitPath(abspath)
			node = node.DirChangeNode(vd,pathelem)
			node.HwdTraverseDir(vd,str)
		}else if !singlepath && !havepara{
			pathlist := strings.Split(path," ")
			pathone := pathlist[0]
			pathelemone := SplitPath(pathone)
			strone := pathelemone[len(pathelemone)-1]
			pathelemone = pathelemone[:len(pathelemone)-1]
			pathone = strings.Join(pathelemone , GetSeparatorChar())
			abspathone := commandmanage.DirConvert(vd, pathone)
			pathelemone = SplitPath(abspathone)
			pathtwo := pathlist[1]
			pathelemtwo := SplitPath(pathtwo)
			strtwo := pathelemtwo[len(pathelemtwo)-1]
			pathelemtwo = pathelemtwo[:len(pathelemtwo)-1]
			pathtwo = strings.Join(pathelemtwo , GetSeparatorChar())
			abspathtwo := commandmanage.DirConvert(vd, pathtwo)
			pathelemtwo = SplitPath(abspathtwo)
			node = node.DirChangeNode(vd,pathelemone)
			node.HwdTraverseDir(vd,strone)
			node = node.DirChangeNode(vd,pathelemtwo)
			node.HwdTraverseDir(vd,strtwo)
		}else if singlepath && havepara{
			pathelem := SplitPath(path)
			str := pathelem[len(pathelem)-1]
			pathelem = pathelem[:len(pathelem)-1]
			path = strings.Join(pathelem, GetSeparatorChar())
			abspath := commandmanage.DirConvert(vd, path)
			pathelem = SplitPath(abspath)
			node = node.DirChangeNode(vd,pathelem)
			if para == s{
				node.HwdTraverseDirS(vd, str)
			}else if para == ad{
				node.HwdTraverseDirAd(vd, str)
			}else if para == "ads"{
				node.HwdTraverseDirAd(vd, str)
				node = node.DirChangeNode(vd,pathelem)
				node.HwdTraverseDirS(vd, str)
			}
		}else{
			pathlist := strings.Split(path," ")
			pathone := pathlist[0]
			pathelemone := SplitPath(pathone)
			strone := pathelemone[len(pathelemone)-1]
			pathelemone = pathelemone[:len(pathelemone)-1]
			pathone = strings.Join(pathelemone , GetSeparatorChar())
			abspathone := commandmanage.DirConvert(vd, pathone)
			pathelemone = SplitPath(abspathone)
			pathtwo := pathlist[1]
			pathelemtwo := SplitPath(pathtwo)
			strtwo := pathelemtwo[len(pathelemtwo)-1]
			pathelemtwo = pathelemtwo[:len(pathelemtwo)-1]
			pathtwo = strings.Join(pathelemtwo , GetSeparatorChar())
			abspathtwo := commandmanage.DirConvert(vd, pathtwo)
			pathelemtwo = SplitPath(abspathtwo)
			node = node.DirChangeNode(vd, pathelemone)
			if para == s{
				node.HwdTraverseDirS(vd, strone)
			}else if para == ad{
				node.HwdTraverseDirAd(vd, strone)
			}else if para == "ads"{
				node.HwdTraverseDirAd(vd, strone)
				node = node.DirChangeNode(vd,pathelemone)
				node.HwdTraverseDirS(vd, strone)
			}
			node = node.DirChangeNode(vd, pathelemtwo)
			if vd.CurrentFolder != &vd.RootComponent{
				node.DirCurrentAndFatherFolder()
			}
			if para == s{
				node.HwdTraverseDirS(vd, strtwo)
			}else if para == ad{
				node.HwdTraverseDirAd(vd, strtwo)
			}else if para == "ads"{
				node.HwdTraverseDirAd(vd, strtwo)
				node = node.DirChangeNode(vd,pathelemtwo)
				node.HwdTraverseDirS(vd, strtwo)
			}
		}
	case 2:
		fmt.Println(TrueDiskError())
		node = nil
	case 3:
		fmt.Println(TrueDiskError())
		node = nil
	}
	vd.Restart()
}