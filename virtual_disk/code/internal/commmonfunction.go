package internal

import (
	"fmt"
	"strings"
	"time"
)

func GetCurrentTime() string{
	return time.Now().Format("2006/01/02 15.04")
}

func GetRootDrive()  string{
	return  "C:"
}

func GetSeparatorChar() string{
	return "\\"
}

func OutputRootDrive(){
	fmt.Print(GetRootDrive(), GetSeparatorChar())
}

func  ConvertCase (cs string) string{ //检查输入字符字母的大小写,将大小写字母统一转化成小写字母
	cschar := []byte(cs)
	if cschar != nil {
		for i :=0; i< len(cschar); i++{
			if( cschar[i] > 64 && cschar[i] < 91){
				cschar[i] +=32
			} else {
				continue
			}
		}
	}
	cs = string(cschar)
	return cs
}

func IsExits (str1 string, str2 string) bool{
	return  strings.EqualFold(str1, str2)
}

func GenerateSpace (a int) string{
	var str string
	for i:=0; i<a ;i++{
		str += " "
	}
	return str
}

func  SplitPath(cs string) []string{
	csl := strings.Split(cs,GetSeparatorChar())//字符串转字符串数组
	return csl
}
