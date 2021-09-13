package internal

import (
	"fmt"
	"time"
)

func GetCurrentTime() string{
	return time.Now().Format("2006/01/02 15.04")
}

func GetRootDrive()  string{
	return  "c:"
}

func GetSeparatorChar() string{
	return "\\"
}

func PrintPath(path *Path) {
	if path != nil{
		fmt.Println(path.Name)
		fmt.Println("\\")
		path = path.Next
	}
}

