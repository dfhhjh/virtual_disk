package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)


func (load Load) CommandExecute(vd *VirtualDisk, path string){
	var sr SerializeRecord
	if strings.HasPrefix(path,"@"){
		path = strings.TrimPrefix(path,"@")
	}
	vd.UpdateCurrentFolder(&vd.RootComponent)
	data,error := ioutil.ReadFile(path)
	if error != nil {
		fmt.Println(error)
		vd.Restart()
	}
	json.Unmarshal(data, &sr)
	for i := 0; i < sr.Size; i++{
		var node = &Component{}
		vd.UpdateCurrentFolder(&vd.RootComponent)
		pathelem := SplitPath(sr.Path[i])
		if sr.IsFoLder[i] == true{
			node.AddFolder(vd, pathelem)
		}else{
			node.LoadFile(vd, pathelem, sr.Content[i])
		}
	}
	vd.Restart()
}
