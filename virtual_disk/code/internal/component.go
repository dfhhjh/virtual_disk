package internal

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Component struct{
	Name           	string
	Time     		string
	IsFolder		bool //IsFolder = true 为目录Folder，IsFolder= false 为文件File
	Path         	string
	FatherComponent	*Component
	SonComponent 	[]*Component
	Content 		[]byte
	FileLength 		int
}

func (currentnode Component) Getpath() string{
	return currentnode.Path
}

func (currentnode Component) GetFatherComponentPath() string{
	return currentnode.FatherComponent.Path
}

func (currentnode Component) GetName() string{
	return currentnode.Name
}

func (currentnode Component) GetTime() string{
	return currentnode.Time
}

func (currentnode *Component) GetFileLength() int{
	return currentnode.FileLength
}

func (currentnode *Component) GetFileContent() []byte{
	return currentnode.Content
}

func (currentnode *Component) DirCurrentAndFatherFolder(){
	str1 := currentnode.Time +  GenerateSpace(4) + "<DIR>" +  GenerateSpace(10) + "."
	str2 := currentnode.FatherComponent.Time  +  GenerateSpace(4) + "<DIR>" +  GenerateSpace(10) + ".."
	fmt.Println(str1)
	fmt.Println(str2)
}

func (currentnode *Component) DirFolder(){
	str1 := currentnode.Time +  GenerateSpace(4) + "<DIR>" +  GenerateSpace(10) + currentnode.Name
	fmt.Println(str1)
}

func (currentnode *Component) DirFile(){
	fmt.Print(currentnode.Time)
  	fmt.Print(GenerateSpace(7) )
  	fmt.Printf("% 10d",currentnode.FileLength)
  	fmt.Print(GenerateSpace(2))
  	fmt.Println(currentnode.Name)
}

func (currentnode *Component)  MatchSonComponent(vd *VirtualDisk, name string) (bool,*Component) {
	var node = &Component{}
	for _, val := range currentnode.SonComponent {
		if IsExits(val.Name, name) {
			node = val
			vd.UpdateCurrentFolder(val)
			return true, node
		}
	}
	return false, node
}

func (currentnode *Component) TraverseDir(vd *VirtualDisk){
	fmt.Println(currentnode.Path,"的目录")
	if vd.CurrentFolder != &vd.RootComponent{
		currentnode.DirCurrentAndFatherFolder()
	}
	for _,v := range currentnode.SonComponent{
		if v.IsFolder{
			v.DirFolder()
		}else{
			v.DirFile()
		}
	}
}

func (currentnode *Component) TraverseDirAd(vd *VirtualDisk){
	fmt.Println(vd.CurrentFolder.Path,"的目录")
	if vd.CurrentFolder != &vd.RootComponent{
		currentnode.DirCurrentAndFatherFolder()
	}
	for _,v := range vd.CurrentFolder.SonComponent{
		if v.IsFolder{
			v.DirFolder()
			vd.UpdateCurrentFolder(v)
		}
	}
}

func (currentnode *Component) TraverseDirS(vd *VirtualDisk)  {
	currentnode = vd.CurrentFolder
	if currentnode.SonComponent == nil{
		currentnode.TraverseDir(vd)
	}else{
		currentnode.TraverseDir(vd)
		for _,v := range currentnode.SonComponent{
			if v.IsFolder{
				vd.UpdateCurrentFolder(v)
				v.TraverseDirS(vd)
			}
		}
	}
}

func (currentnode *Component) HwdTraverseDir(vd *VirtualDisk, str string){
	if strings.Contains(str, "*"){
		strlist := strings.Split(str, "*")
		strpre := strlist[0]
		strsuffix := strlist[1]
		for _,v := range vd.CurrentFolder.SonComponent{
			if strings.HasPrefix(v.Name, strpre) && strings.HasSuffix(v.Name, strsuffix){
				vd.UpdateCurrentFolder(v)
				currentnode.TraverseDir(vd)
			}
		}
	}else if strings.Contains(str, "?"){
		posi := strings.Index(str, "?")
		prestr := str[0:posi]
		posilast := strings.LastIndex(str,"?")
		suffixstr := str[posilast +1:]
		for _,v := range vd.CurrentFolder.SonComponent{
			if strings.HasPrefix(v.Name, prestr) && strings.HasSuffix(v.Name,suffixstr) && len(v.Name) == len(str){
				vd.UpdateCurrentFolder(v)
				currentnode.TraverseDir(vd)
			}
		}
	}

}

func (currentnode *Component) HwdTraverseDirAd(vd *VirtualDisk, str string){
	if strings.Contains(str, "*"){
		strlist := strings.Split(str, "*")
		strpre := strlist[0]
		strsuffix := strlist[1]
		for _,v := range vd.CurrentFolder.SonComponent{
			if strings.HasPrefix(v.Name, strpre) && strings.HasSuffix(v.Name, strsuffix){
				vd.UpdateCurrentFolder(v)
				currentnode.TraverseDirAd(vd)

			}
		}
	}else if strings.Contains(str, "?"){
		posi := strings.Index(str, "?")
		prestr := str[0:posi]
		posilast := strings.LastIndex(str,"?")
		suffixstr := str[posilast +1:len(str)]
		for _,v := range vd.CurrentFolder.SonComponent{
			if strings.HasPrefix(v.Name, prestr) && strings.HasSuffix(v.Name,suffixstr) && len(v.Name) == len(str){
				vd.UpdateCurrentFolder(v)
				currentnode.TraverseDirAd(vd)
			}
		}
	}
}

func (currentnode *Component) HwdTraverseDirS(vd *VirtualDisk, str string) {
	if strings.Contains(str, "*") {
		strlist := strings.Split(str, "*")
		strpre := strlist[0]
		strsuffix := strlist[1]
		for _, v := range vd.CurrentFolder.SonComponent {
			if strings.HasPrefix(v.Name, strpre) && strings.HasSuffix(v.Name, strsuffix) {
				vd.UpdateCurrentFolder(v)
				currentnode.TraverseDirS(vd)
			}
		}
	} else if strings.Contains(str, "?") {
		posi := strings.Index(str, "?")
		prestr := str[0:posi]
		posilast := strings.LastIndex(str, "?")
		suffixstr := str[posilast+1 : len(str)]
		for _, v := range vd.CurrentFolder.SonComponent {
			if strings.HasPrefix(v.Name, prestr) && strings.HasSuffix(v.Name, suffixstr) && len(v.Name) == len(str) {
				vd.UpdateCurrentFolder(v)
				currentnode.TraverseDirS(vd)
			}
		}
	}
}

func (currentnode *Component) ChangeNode(vd *VirtualDisk, pathelement  []string) *Component {
	var changenode *Component
	var match bool
	for _,v := range pathelement {
		if v == GetRootDrive() {
			changenode = &vd.RootComponent
			continue
		}
		match,changenode = changenode.MatchSonComponent(vd, v)
		if match && v == pathelement[len(pathelement)-1]{
			return  changenode
		} else if(match){
			continue
		} else{
			PathError()
			vd.Execute()
		}
		return changenode
	}
	return &vd.RootComponent
}

func (currentnode *Component) DirChangeNode(vd *VirtualDisk, pathelement  []string) *Component {
	var changenode = &Component{}
	changenode = vd.CurrentFolder
	var match bool
	for _,v := range pathelement {
		if v == GetRootDrive() {
			changenode = &vd.RootComponent
			continue
		}
		if changenode.SonComponent == nil && v != pathelement[len(pathelement)-1] {
			fmt.Println(PathError())
			vd.Restart()
		}
		match,changenode = changenode.MatchSonComponent(vd, v)
		if match && v == pathelement[len(pathelement)-1]{
			return  changenode
		} else if(match){
			continue
		} else{
			fmt.Println(PathError())
			vd.Restart()
		}
		return changenode
	}
	return &vd.RootComponent
}

func (currentnode *Component) AddFolder(vd *VirtualDisk, pathelement  []string) {
	var fathernode *Component
	var addnode *Component
	var match bool
	for _,v := range pathelement {
		if v == GetRootDrive() {
			addnode = &vd.RootComponent
			vd.UpdateCurrentFolder(&vd.RootComponent)
			continue
		}
		match,addnode = addnode.MatchSonComponent(vd, v)
		if(match){
			continue
		} else if v == pathelement[len(pathelement)-1]{
			fathernode = vd.CurrentFolder
			addnode = addnode.SetFolder(fathernode ,v)
			vd.UpdateCurrentFolder(addnode)
		}else{
			fathernode = vd.CurrentFolder
			addnode = addnode.SetFolder(fathernode ,v)
			vd.UpdateCurrentFolder(addnode)
			continue
		}

	}
}

func (currentnode *Component) AddFolderTwo(vd *VirtualDisk, pathelement  []string, node *Component){
	var fathernode *Component
	var addnode *Component
	var match bool
	for _,v := range pathelement {
		if v == GetRootDrive() {
			addnode = &vd.RootComponent
			vd.UpdateCurrentFolder(&vd.RootComponent)
			continue
		}
		match,addnode = addnode.MatchSonComponent(vd, v)
		 if match && v == pathelement[len(pathelement)-1]{
			fathernode = vd.CurrentFolder
			addnode = addnode.SetFolder(fathernode, node.Name)
			vd.UpdateCurrentFolder(addnode)
		}else if(match){
			 continue
		}else{
			fmt.Println(PathError())
		}
	}
}

func (currentnode *Component) AddFolderTwoY(vd *VirtualDisk, pathelement  []string, node *Component) {
	var fathernode *Component
	var addnode *Component
	var match bool
	for _,v := range pathelement {
		if v == GetRootDrive() {
			addnode = &vd.RootComponent
			vd.UpdateCurrentFolder(&vd.RootComponent)
			continue
		}
		match,addnode = addnode.MatchSonComponent(vd, v)
		if match && v == pathelement[len(pathelement)-1]{
			if v == node.Name{
				for k,val := range addnode.FatherComponent.SonComponent{
					if val == addnode{
						slice := addnode.FatherComponent.SonComponent
						addnode.FatherComponent.SonComponent = append(slice[:k], slice[k+1:]...)
					}
				}
				addnode = node
				addnode.FatherComponent = vd.CurrentFolder.FatherComponent
				addnode.FatherComponent.SonComponent = append(addnode.FatherComponent.SonComponent, node)
				vd.CurrentFolder.Time = GetCurrentTime()
				addnode.Time = GetCurrentTime()

			}else{
				fathernode = vd.CurrentFolder
				addnode = addnode.SetFolder(fathernode, node.Name)
				vd.UpdateCurrentFolder(addnode)
			}
		}else if(match){
			continue
		}else{
			fmt.Println(PathError())
		}
	}
}

func (currentnode *Component) SetFolder( fathernode *Component, name string) *Component{
	var setnode = &Component{}
	setnode.Name = name
	setnode.FatherComponent = fathernode
	fathernode.SonComponent = append(fathernode.SonComponent, setnode)
	fathernode.Time = GetCurrentTime()
	setnode.SonComponent = nil
	setnode.Path = fathernode.Path + GetSeparatorChar() + name
	setnode.IsFolder = true
	setnode.Time = GetCurrentTime()
	return setnode
}

func (currentnode *Component) TouchFile(vd *VirtualDisk, pathelement  []string) {
	var fathernode *Component
	var touchnode *Component
	var match bool
	for _,v := range pathelement {
		if v == GetRootDrive() {
			touchnode = &vd.RootComponent
			vd.UpdateCurrentFolder(&vd.RootComponent)
			continue
		}
		match, touchnode = touchnode.MatchSonComponent(vd, v)
		if match && v == pathelement[len(pathelement)-1] {
			touchnode.Time = GetCurrentTime()
		}else if match {
			continue
		} else {
			fathernode = vd.CurrentFolder
			touchnode = touchnode.SetFile(fathernode ,v)
		}
	}
}

func (currentnode *Component) LoadFile(vd *VirtualDisk, pathelement  []string, content []byte) {
	var fathernode *Component
	var touchnode *Component
	var match bool
	for _,v := range pathelement {
		if v == GetRootDrive() {
			touchnode = &vd.RootComponent
			vd.UpdateCurrentFolder(&vd.RootComponent)
			continue
		}
		match, touchnode = touchnode.MatchSonComponent(vd, v)
		if match && v == pathelement[len(pathelement)-1] {
			touchnode.Time = GetCurrentTime()
		}else if match {
			continue
		} else {
			fathernode = vd.CurrentFolder
			touchnode = touchnode.SetLoadFile(fathernode , v, content)
		}
	}
}

func (currentnode *Component) SetLoadFile(fathernode *Component, name string, content []byte) *Component{
	var setnode = &Component{}
	setnode.Name = name
	setnode.FatherComponent = fathernode
	setnode.SonComponent = nil
	fathernode.SonComponent = append(fathernode.SonComponent, setnode)
	fathernode.Time = GetCurrentTime()
	setnode.Path = fathernode.Path + GetSeparatorChar() + name
	setnode.IsFolder = false
	setnode.Time = GetCurrentTime()
	setnode.Content = content
	setnode.FileLength = len(setnode.Content)
	return setnode
}

func (currentnode *Component) AddFile(vd *VirtualDisk, pathelement  []string,  node *Component){
	var fathernode *Component
	var addnode *Component
	var match bool
	for _,v := range pathelement {
		if v == GetRootDrive() {
			addnode = &vd.RootComponent
			vd.UpdateCurrentFolder(&vd.RootComponent)
			continue
		}
		match, addnode = addnode.MatchSonComponent(vd, v)
		 if match && v == pathelement[len(pathelement)-1]{
			fathernode = vd.CurrentFolder
			addnode = addnode.SetFile(fathernode, node.Name)
		}else if(match){
			continue
		} else{
			fmt.Println(PathError())
		}
	}
}

func (currentnode *Component) CopyFile(vd *VirtualDisk, pathelement  []string,  node *Component){
	var fathernode *Component
	var addnode *Component
	var match bool
	for _,v := range pathelement {
		if v == GetRootDrive() {
			addnode = &vd.RootComponent
			vd.UpdateCurrentFolder(&vd.RootComponent)
			continue
		}
		match, addnode = addnode.MatchSonComponent(vd, v)
		if match && v == pathelement[len(pathelement)-1] && addnode.IsFolder == true{
			fathernode = vd.CurrentFolder
			addnode = addnode.SetFile(fathernode, node.Name)
		}else if match && v == pathelement[len(pathelement)-1] && addnode.IsFolder == false && v == node.Name {
			vd.CurrentFolder.Time = GetCurrentTime()
			addnode.Time = GetCurrentTime()
		}else if match && v == pathelement[len(pathelement)-1] && addnode.IsFolder == false && v!= node.Name{
			fmt.Println("是否进行覆盖(y\\n\\all)")
			var input string
			fmt.Scan(&input)
			if input == "y"{
				addnode.Name = node.Name
				vd.CurrentFolder.Time = GetCurrentTime()
				addnode.Time = GetCurrentTime()
			}else if input == "n"{
				fathernode = vd.CurrentFolder
				addnode = addnode.SetFile(fathernode, node.Name)
			}else if input == "all"{
				addnode.Name = node.Name
				vd.CurrentFolder.Time = GetCurrentTime()
				addnode.Time = GetCurrentTime()
			}else{
				fmt.Println("输入错误")
				vd.Restart()
			}

		}else if match {
			continue
		} else{
			fmt.Println(PathError())
		}
	}
}

func (currentnode *Component) HwdAddFile(vd *VirtualDisk, pathelement  []string, name string) {
	var addnode *Component
	if strings.Contains(name, "*"){
		strlist := strings.Split(name, "*")
		strpre := strlist[0]
		strsuffix := strlist[1]
		for i, v := range vd.CurrentFolder.SonComponent{
			if strings.HasPrefix(v.Name, strpre) && strings.HasSuffix(v.Name, strsuffix){
				if i != 0 {
					var node *Component
					node.AddFile(vd , pathelement ,v)
				}else{
					addnode.AddFile(vd , pathelement ,v)
				}
			}
		}
	}else if strings.Contains(name, "?"){
		posi := strings.Index(name, "?")
		prestr := name[0:posi]
		posilast := strings.LastIndex(name,"?")
		suffixstr := name[posilast +1:len(name)]
		for i,v := range vd.CurrentFolder.SonComponent{
			if strings.HasPrefix(v.Name, prestr) && strings.HasSuffix(v.Name,suffixstr) && len(v.Name) == len(name){
				if i != 0{
					var node *Component
					node.AddFile(vd , pathelement ,v)
				}else{
					addnode.AddFile(vd , pathelement ,v)
				}
			}
		}
	}
}

func (currentnode *Component) AddFileY(vd *VirtualDisk, pathelement  []string,  node *Component) {
	var fathernode *Component
	var addnode *Component
	var match bool
	for _,v := range pathelement {
		if v == GetRootDrive() {
			addnode = &vd.RootComponent
			vd.UpdateCurrentFolder(&vd.RootComponent)
			continue
		}
		match, addnode = addnode.MatchSonComponent(vd, v)
		if match && v == pathelement[len(pathelement)-1]{
			if v == node.Name{
				for k,val := range addnode.FatherComponent.SonComponent{
					if val == addnode{
						slice := addnode.FatherComponent.SonComponent
						addnode.FatherComponent.SonComponent = append(slice[:k], slice[k+1:]...)
					}
				}
				addnode = node
				addnode.FatherComponent = vd.CurrentFolder.FatherComponent
				addnode.FatherComponent.SonComponent = append(addnode.FatherComponent.SonComponent, node)
				vd.CurrentFolder.Time = GetCurrentTime()
				addnode.Time = GetCurrentTime()
				currentnode.RemoveComponent(vd, pathelement)
			}else{
				fathernode = vd.CurrentFolder
				addnode = addnode.SetFile(fathernode, node.Name)
				vd.UpdateCurrentFolder(addnode)
			}
		}else if(match){
			continue
		} else{
			fmt.Println(PathError())
		}
	}
}

func (currentnode *Component) TrueDiskAddFile(vd *VirtualDisk, pathone string, pathelementtwo  []string){
	var fathernode *Component
	var addnode *Component
	var match bool
	pathelementone := SplitPath(pathone)
	data, err := ioutil.ReadFile(pathone)
	if err != nil {
		fmt.Println("File reading error", err)
		vd.UpdateCurrentFolder(&vd.RootComponent)
		OutputRootDrive()
		vd.Execute()
	}
	name := pathelementone[len(pathelementone)-1]
	for _,v := range pathelementtwo {
		if v == GetRootDrive() {
			addnode = &vd.RootComponent
			vd.UpdateCurrentFolder(&vd.RootComponent)
			continue
		}
		match, addnode = addnode.MatchSonComponent(vd, v)
		if match && v == pathelementtwo[len(pathelementtwo)-1] && addnode.IsFolder == true{
			fathernode = vd.CurrentFolder
			addnode = addnode.SetFile(fathernode, name)
			addnode.Content = data
			addnode.FileLength = len(data)
		}else if match && v == pathelementtwo[len(pathelementtwo)-1] && addnode.IsFolder == false && v == name {
			vd.CurrentFolder.Time = GetCurrentTime()
			addnode.Time = GetCurrentTime()
			addnode.Content = data
			addnode.FileLength = len(data)
		}else if match && v == pathelementtwo[len(pathelementtwo)-1] && addnode.IsFolder == false && v!= name{
			addnode.Name = name
			vd.CurrentFolder.Time = GetCurrentTime()
			addnode.Time = GetCurrentTime()
			addnode.Content = data
			addnode.FileLength = len(data)
		}else if match {
			continue
		} else{
			fmt.Println(PathError())
		}
	}
}

func (currentnode *Component) HwdTrueDiskAddFile(vd *VirtualDisk, elemone string, pathone string, pathelementtwo  []string){
	var fathernode *Component
	var match bool
	err := filepath.Walk(pathone, func(path string, f os.FileInfo, err error) error {
		if ( f == nil ) {return err}
		if f.IsDir() {return nil}
		pathelem := strings.Split(path, "E")
		for _, v := range pathelem{
			var addnode *Component
			v = "E" + v
			velem := strings.Split(v, GetSeparatorChar())
			lastelem := velem[len(velem)-1]
			if strings.Contains(elemone, "*"){
				strlist := strings.Split(elemone, "*")
				strpre := strlist[0]
				strsuffix := strlist[1]
			if strings.HasPrefix(lastelem, strpre) && strings.HasSuffix(lastelem, strsuffix){
				data, erro := ioutil.ReadFile(v)
				if erro != nil {
					fmt.Println("File reading error", erro)
					vd.UpdateCurrentFolder(&vd.RootComponent)
					OutputRootDrive()
					vd.Execute()
				}
				for _,val := range pathelementtwo {
					if val == GetRootDrive() {
						addnode = &vd.RootComponent
						vd.UpdateCurrentFolder(&vd.RootComponent)
						continue
					}
					match, addnode = addnode.MatchSonComponent(vd, val)
					if match && val == pathelementtwo[len(pathelementtwo)-1] && addnode.IsFolder == true{
						fathernode = vd.CurrentFolder
						addnode = addnode.SetFile(fathernode, lastelem)
						addnode.Content = data
						addnode.FileLength = len(data)
					}else if match && val == pathelementtwo[len(pathelementtwo)-1] && addnode.IsFolder == false && val == lastelem {
						vd.CurrentFolder.Time = GetCurrentTime()
						addnode.Time = GetCurrentTime()
						addnode.Content = data
						addnode.FileLength = len(data)
					}else if match && val == pathelementtwo[len(pathelementtwo)-1] && addnode.IsFolder == false && val != lastelem{
						addnode.Name = lastelem
						vd.CurrentFolder.Time = GetCurrentTime()
						addnode.Time = GetCurrentTime()
						addnode.Content = data
						addnode.FileLength = len(data)
					}else if match {
						continue
					} else{
						fmt.Println(PathError())
					}
				}

			}

			}else if  strings.Contains(elemone, "?"){
				posi := strings.Index(elemone, "?")
				prestr := elemone[0:posi]
				posilast := strings.LastIndex(elemone,"?")
				suffixstr := elemone[posilast +1:len(elemone)]
				if strings.HasPrefix(lastelem, prestr) && strings.HasSuffix(lastelem,suffixstr) && len(lastelem) == len(elemone) {
					data, erro := ioutil.ReadFile(v)
					if erro != nil {
						fmt.Println("File reading error", erro)
						vd.UpdateCurrentFolder(&vd.RootComponent)
						OutputRootDrive()
						vd.Execute()
					}
					for _, val := range pathelementtwo {
						if val == GetRootDrive() {
							addnode = &vd.RootComponent
							vd.UpdateCurrentFolder(&vd.RootComponent)
							continue
						}
						match, addnode = addnode.MatchSonComponent(vd, val)
						if match && val == pathelementtwo[len(pathelementtwo)-1] && addnode.IsFolder == true {
							fathernode = vd.CurrentFolder
							addnode = addnode.SetFile(fathernode, lastelem)
							addnode.Content = data
							addnode.FileLength = len(data)
						} else if match && val == pathelementtwo[len(pathelementtwo)-1] && addnode.IsFolder == false && val == lastelem {
							vd.CurrentFolder.Time = GetCurrentTime()
							addnode.Time = GetCurrentTime()
							addnode.Content = data
							addnode.FileLength = len(data)
						} else if match && val == pathelementtwo[len(pathelementtwo)-1] && addnode.IsFolder == false && val != lastelem {
							addnode.Name = lastelem
							vd.CurrentFolder.Time = GetCurrentTime()
							addnode.Time = GetCurrentTime()
							addnode.Content = data
							addnode.FileLength = len(data)
						} else if match {
							continue
						} else {
							fmt.Println(PathError())
						}
					}
				}
			}
		}
		return nil
		})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func (currentnode *Component) SetFile(fathernode *Component, name string) *Component{
	var setnode = &Component{}
	setnode.Name = name
	setnode.FatherComponent = fathernode
	setnode.SonComponent = nil
	fathernode.SonComponent = append(fathernode.SonComponent, setnode)
	fathernode.Time = GetCurrentTime()
	setnode.Path = fathernode.Path + GetSeparatorChar() + name
	setnode.IsFolder = false
	setnode.Time = GetCurrentTime()
	setnode.FileLength = 0
	return setnode
}

func (currentnode *Component) RemoveComponent(vd *VirtualDisk, pathelement  []string){
	var match bool
	for _,v := range pathelement {
		if v == GetRootDrive() {
			currentnode  = &vd.RootComponent
			vd.UpdateCurrentFolder(&vd.RootComponent)
			continue
		}
		match, currentnode = currentnode .MatchSonComponent(vd, v)
		if(match){
			if(v == pathelement[len(pathelement)-1]){
				for k,val := range currentnode.FatherComponent.SonComponent{
					if val == currentnode{
						slice := currentnode.FatherComponent.SonComponent
						currentnode.FatherComponent.SonComponent = append(slice[:k], slice[k+1:]...)
					}
				}
				currentnode = nil
				vd.UpdateCurrentFolder(nil)
			}else{
				continue
			}
		}else{
			PathError()
			vd.Execute()
		}
	}
}

func (currentnode *Component) RemoveComponentY(vd *VirtualDisk, pathelement  []string) {
	var match bool
	for _, v := range pathelement {
		if v == GetRootDrive() {
			currentnode = &vd.RootComponent
			vd.UpdateCurrentFolder(&vd.RootComponent)
			continue
		}
		match, currentnode = currentnode.MatchSonComponent(vd, v)
		if (match) {
			if match {
				if (v == pathelement[len(pathelement)-2]) {
					for k, val := range currentnode.SonComponent {
						if val.Name == pathelement[len(pathelement)-1] {
							val = nil
							slice := currentnode.SonComponent
							currentnode.SonComponent = append(slice[:k], slice[k+1:]...)
							vd.UpdateCurrentFolder(nil)
						}
					}
				}
				continue
				currentnode = nil
				vd.UpdateCurrentFolder(nil)
			} else {
				PathError()
				vd.Execute()
			}
		}
	}
}

func (currentnode *Component) RemoveComponentS(vd *VirtualDisk, pathelement  []string){
	var match bool
	for _,v := range pathelement {
		if v == GetRootDrive() {
			currentnode  = &vd.RootComponent
			vd.UpdateCurrentFolder(&vd.RootComponent)
			continue
		}
		match, currentnode = currentnode .MatchSonComponent(vd, v)
		if match {
			if(v == pathelement[len(pathelement)-1]){
				for _, value := range currentnode.SonComponent{
					if value != nil{
						valueelem := SplitPath(value.Path)
						value.RemoveComponentS(vd, valueelem)
					}
				}
				for k,val := range currentnode.FatherComponent.SonComponent{
					if val == currentnode{
						slice := currentnode.FatherComponent.SonComponent
						currentnode.FatherComponent.SonComponent = append(slice[:k], slice[k+1:]...)
					}
				}
				currentnode = nil
				vd.UpdateCurrentFolder(nil)
			}
			continue
		}else{
			PathError()
			vd.Execute()
		}
	}
}

func (currentnode *Component) RemoveFolder(vd *VirtualDisk, pathelement  []string){
	var match bool
	for _,v := range pathelement {
		if v == GetRootDrive() {
			currentnode  = &vd.RootComponent
			vd.UpdateCurrentFolder(&vd.RootComponent)
			continue
		}
		match, currentnode = currentnode .MatchSonComponent(vd, v)
		if(match){
			if(v == pathelement[len(pathelement)-1] && currentnode.IsFolder == true){
				for k,val := range currentnode.FatherComponent.SonComponent{
					if val == currentnode{
						slice := currentnode.FatherComponent.SonComponent
						currentnode.FatherComponent.SonComponent = append(slice[:k], slice[k+1:]...)
					}
				}
				currentnode = nil
				vd.UpdateCurrentFolder(nil)
			}
			continue
		}else{
			PathError()
			vd.Execute()
		}
	}
}

func (currentnode *Component) RemoveFolderS(vd *VirtualDisk, pathelement  []string){
	var match bool
	for _,v := range pathelement {
		if v == GetRootDrive() {
			currentnode  = &vd.RootComponent
			vd.UpdateCurrentFolder(&vd.RootComponent)
			continue
		}
		match, currentnode = currentnode .MatchSonComponent(vd, v)
		if(match){
			if(v == pathelement[len(pathelement)-1] && currentnode.IsFolder == true){
				for _, value := range currentnode.SonComponent{
					if value != nil{
						valueelem := SplitPath(value.Path)
						value.RemoveComponentS(vd, valueelem)
					}
				}
				for k,val := range currentnode.FatherComponent.SonComponent{
					if val == currentnode{
						slice := currentnode.FatherComponent.SonComponent
						currentnode.FatherComponent.SonComponent = append(slice[:k], slice[k+1:]...)
					}
				}
				currentnode = nil
				vd.UpdateCurrentFolder(nil)
			}
			continue
		}else{
			PathError()
			vd.Execute()
		}
	}
}

func (currentnode *Component) RenComponent (vd *VirtualDisk, pathelement  []string, name string){
	var match bool
	for _,v := range pathelement {
		if v == GetRootDrive() {
			currentnode = &vd.RootComponent
			continue
		}
		match, currentnode = currentnode.MatchSonComponent(vd, v)
		if match && v == pathelement[len(pathelement)-1]{
			currentnode.Name = name
			currentnode.Path = currentnode.FatherComponent.Path + GetSeparatorChar() + currentnode.Name
		} else if(match){
			continue
		} else{
			PathError()
			vd.Execute()
		}
	}
}


