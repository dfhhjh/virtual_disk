package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func (save Save) CommandExecute(vd *VirtualDisk, path string){
	var sr = &SerializeRecord{}
	if strings.HasPrefix(path,"@"){
		path = strings.TrimPrefix(path,"@")
	}
	vd.UpdateCurrentFolder(&vd.RootComponent)
	f, err := os.Create(path+GetSeparatorChar()+"serialize.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	sr = sr.SetSerializeRecord(vd.CurrentFolder)
	sr.SerializeSonMatch(vd.CurrentFolder)
	b, _ := json.Marshal(sr)
	f.Write(b)
	vd.UpdateCurrentFolder(&vd.RootComponent)
	OutputRootDrive()
	vd.Execute()
	}

