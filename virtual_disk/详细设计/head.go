package main
import "fmt"

type file struct{
	name string
	filelen int
	filetime int
}

type content struct {
	name string
	contime int
	soncon *content
	brocon *content
	facon *content
}

