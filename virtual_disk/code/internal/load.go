package internal

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

func (load Load) CommandExecute(vd *VirtualDisk, path string){
	var sr SerializeRecord
	if strings.HasPrefix(path,"@"){
		path = strings.TrimPrefix(path,"@")
	}
	vd.UpdateCurrentFolder(&vd.RootComponent)
	data,_ := ioutil.ReadFile(path)
	json.Unmarshal(data, &sr)
	for i := 0; i < sr.Size; i++{
		var node = &Component{}
		vd.UpdateCurrentFolder(&vd.RootComponent)
		pathelem := SplitPath(sr.Path[i])
		if sr.IsFoLder[i] == true{
			node.AddFolder(vd, pathelem)
		}else{
			node.TouchFile(vd, pathelem)
		}
		node.Content = sr.Content[i]
		node.FileLength = len(node.Content)
	}
	vd.UpdateCurrentFolder(&vd.RootComponent)
	OutputRootDrive()
	vd.Execute()
}
