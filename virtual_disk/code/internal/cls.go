package internal

import (
	"os"
	"os/exec"
)

/*
#include<stdlib.h>

void cls(){
system("cls");
}
 */
/*
import (
	"C"
)

func (cls Cls) CommandExecute(vd *VirtualDisk){
	C.cls()
	vd.UpdateCurrentFolder(&vd.RootComponent)
	OutputRootDrive()
	vd.Execute()
}*/


func (cls Cls) CommandExecute(vd *VirtualDisk){
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	vd.UpdateCurrentFolder(&vd.RootComponent)
	OutputRootDrive()
	vd.Execute()
}

