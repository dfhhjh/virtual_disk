package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
type CommandManage struct {
	InputString string
}

func (cm CommandManage) CreateCommand() string{
	inputReader := bufio.NewReader(os.Stdin)//输入的自动类型推导
	input,_:= inputReader.ReadString('\n')//可接受带空格的输入
	cm.InputString = input
	return cm.InputString
}

func (cm CommandManage) CheckSize(cs string) bool{ //检查输入的命令字符数是否超过256字符,返回值为true无错误，返回值为false字数过多
	if len(cs) > 256{
		return false
	}else{
	return true
	}
}

func (cm CommandManage) HaveInvalidCharacters(path string)bool{ //检查命令有没有包含非法字符
	if strings.ContainsAny(path, "<>|"){
		return true
	}
	return false
}

func (cm CommandManage) CheckSeparator(cs string) string{ //检查分隔符是用“/”还是“\”，将输入命令中的分隔符统一转化成“\”
	cschar := []byte(cs)
	if cschar != nil {
		for i :=0; i< len(cschar); i++{
			if( cschar[i] == 47 ){
				cschar[i] += 45
			} else {
				continue
			}
		}
	}
	cs = string(cschar)
	return cs
}

func (cm CommandManage) IsTrueDiskPath(cs string) bool{ //检查是否是真实磁盘的路径，返回值为ture为真实磁盘路径，返回值为false为虚拟磁盘路径
	if  strings.Contains(cs, "@") {
		return true
	}else{
		return false
	}
}

func (cm CommandManage) HaveWildCard(cs string) bool{ //检查是否含有通配符，返回值为ture为含有通配符，返回值为false为不含通配符
	if  strings.ContainsAny(cs, "*|?") {
		return true
	}else{
		return false
	}
}

func (cm CommandManage) GetCommandType(itd bool, hwc bool) int{ //将命令分为4种，非真实磁盘，无通配符回0；非真实磁盘，有通配符，返回1； 真实磁盘，无通配符,返回2；真实磁盘，有通配符，返回3
	if(itd == false && hwc == false ){
		return 0
	} else if(itd == false && hwc == true ){
		return 1
	}else if(itd == true && hwc == false ) {
		return 2
	}else if(itd == true && hwc == true ){
		return 3
	}
	return 8
}

var CommandList = [12]string{dir, md , rd, cd, touch, del, copy, ren, move, save, load, cls}

func (cm CommandManage) SeparateCommand(cs string,vd *VirtualDisk) (string,string){ //返回值为带参数的路径和命令
	cs = strings.TrimSuffix(cs,"\n" )
	var match bool
	for _,v := range CommandList{
		if strings.HasPrefix(cs, v) { //判断字符串是不是以字符串v开头
			cs = strings.TrimPrefix(cs,v)
			cs = strings.TrimPrefix(cs, " ")
			match = true
			return  cs, v//返回去掉v的字符串
		}
	}
	if match == false{
		fmt.Println("命令不正确")
		vd.Restart()
	}
	return "sorry","命令不正确"
}

var ParameterList = map[string][]string{
	dir : {ad,s},
	del : {s},
	rd : {s},
	move : {y},
}

func (cm CommandManage) SeparateParameter(pp string, c string)(string,string){ //返回值为路径和参数
	command := c
	a := len(pp)
	_,t  := ParameterList[command]//先判断这个命令是否含有参数
	if t {
			value, ok := ParameterList[command]
			if ok == true {
				var p string
				var parastr string
				for _, p = range value{//对参数进行循环
					if strings.HasPrefix(pp, GetSeparatorChar()+p){
						pp = strings.TrimPrefix(pp, GetSeparatorChar()+p)
						pp = strings.TrimPrefix(pp, " ")
						parastr += p
					}
				}
				if len(pp) == a{
					return pp, ""
				}else{
					return pp, parastr
				}
			}
		}
	return pp,""
}

func (cm CommandManage) ConvertRelaivePathToAbsolutePath(vd *VirtualDisk, str string) string{ //检查是相对路径还是绝对路径，如果是相对路径就转化为绝对路径
	strlist := SplitPath(str)
	for _,v := range strlist{
		if(v == " "){
			continue //停留在当前目录
		}else if(v == ".") {
			continue
		}else if(v == ".."){
			continue
		}else if(IsExits(v,GetRootDrive())){
			vd.UpdateCurrentFolder(&vd.RootComponent)
			break
		} else{
			vd. GetChildNodeByName(&vd.RootComponent,v)
			break
		}
	}
	var pathlist []string
	for _,value := range strlist{
		if(value == " "){
			continue //停留在当前目录
		}else if(value == ".") {
			continue
		}else if(value == ".."){
		if vd.CurrentFolder.FatherComponent == nil {
			continue
		}
		vd.CurrentFolder = vd.CurrentFolder.FatherComponent //返回上一级目录
		}else if(IsExits(value,GetRootDrive())){
			continue
		} else if value == strlist[len(strlist)-1]{
			pathlist = append(pathlist, value)
		}else{
			pathlist = append(pathlist, value)
			continue
		}
	}
	abstractpath := vd.CurrentFolder.Path
	var pathliststr string
	if len(pathlist) == 0 {
	}else{
		abstractpath = abstractpath + GetSeparatorChar()
		pathliststr = strings.Join(pathlist, GetSeparatorChar())
		abstractpath += pathliststr
	}
	return abstractpath
}

func (cm CommandManage) ConvertRelaivePathToAbsolutePathFile(vd *VirtualDisk, str string) string{ //检查是相对路径还是绝对路径，如果是相对路径就转化为绝对路径
	strlist := SplitPath(str)
	for _,v := range strlist{
		if(v == " "){
			continue //停留在当前目录
		}else if(v == ".") {
			continue
		}else if(v == ".."){
			continue
		}else if(IsExits(v,GetRootDrive())){
			vd.UpdateCurrentFolder(&vd.RootComponent)
			break
		} else{
			vd. GetChildNodeByName(&vd.RootComponent,v)
			break
		}
	}
	var pathlist []string
	for _,value := range strlist{
		if(value == " "){
			continue //停留在当前目录
		}else if(value == ".") {
			continue
		}else if(value == ".."){
			if vd.CurrentFolder.FatherComponent == nil {
				continue
			}
			vd.CurrentFolder = vd.CurrentFolder.FatherComponent //返回上一级目录
		}else if(IsExits(value,GetRootDrive())){
			continue
		} else if value == strlist[len(strlist)-1]{
			pathlist = append(pathlist, value)
		}else{
			pathlist = append(pathlist, value)
			continue
		}
	}
	abstractpath := vd.CurrentFolder.FatherComponent.Path
	var pathliststr string
	if len(pathlist) == 0 {
	}else{
		abstractpath = abstractpath + GetSeparatorChar()
		pathliststr = strings.Join(pathlist, GetSeparatorChar())
		abstractpath += pathliststr
	}
	return abstractpath
}

func (cm CommandManage) DirConvert (vd *VirtualDisk, path string) string{
	var pathlist []string
	strlist := SplitPath(path)
	for _,value := range strlist{
		if(value == ""){
			continue //停留在当前目录
		}else if(value == ".") {
			continue
		}else if(value == ".."){
			if vd.CurrentFolder.FatherComponent == nil {
				continue
			}
			vd.CurrentFolder = vd.CurrentFolder.FatherComponent //返回上一级目录
		}else if(IsExits(value,GetRootDrive())){
			continue
		} else if value == strlist[len(strlist)-1]{
			pathlist = append(pathlist, value)
		}else{
			pathlist = append(pathlist, value)
			continue
		}
	}
	str := strings.Join(pathlist, GetSeparatorChar())
	return str
}

func (cm CommandManage) IsSinglePath(path string) bool{
	if strings.Contains(path, " "){
		return false
	}else{
		return true
	}
}

func (cm CommandManage) HaveParameter(para string) bool{
	if para != ""{
		return true
	}else{
		return false
	}
}









