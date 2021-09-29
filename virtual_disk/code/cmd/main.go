package main

import (
	"go_code/internal"
)

func main() {
	var vd internal.VirtualDisk
	vd.CreateVirtualDisk()
	internal.OutputRootDrive()
	vd.Execute()
}
