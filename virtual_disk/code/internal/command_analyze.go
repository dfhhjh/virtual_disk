package internal

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const  (
	dir   	= "dir"
	md    	= "md"
	rd    	= "rd"
	cd    	= "cd"
	touch  	= "touch"
	del 	= "del"
	copy 	= "copy"
	ren 	= "ren"
	move	= "move"
	mklink	= "mklink"
	save	= "save"
	load	= "load"
	cls		= "cls"
)

const (
	ad	= "\\ad"
	s	= "\\s"
	y	= "\\y"
)

func New(str string) error{
	return &CommandSizeError{str}
}

type CommandSizeError struct{
	CommandName string
}

func (cse *CommandSizeError) Error() string{
	return cse.CommandName
}//以上三个结构体、函数和方法用于处理输入命令数值过大

type CommandManage struct {
	InputString string
	CommandAnalyze
	ParameterAnalyze
	TypeAnalyze
	Split
}

func (ca CommandManage) CreateCommand() string{
	inputReader := bufio.NewReader(os.Stdin)//输入的自动类型推导
	input, err := inputReader.ReadString('\n')//可接受带空格的输入
	if err == nil{
		ca.InputString = input
	}
	return ca.InputString
}

func (ca CommandManage) CheckSize(cs string) (int,error){ //检查输入的命令字符数是否超过256字符，字符过多提醒错误
	if len(cs) > 256{
		return 0, errors.New("输入的字符数过多，请重新输入")
	}
	return 1,nil
}

func (ca CommandManage) CheckCase(cs string) string{ //检查输入字母的大小写,将大小写字母统一转化成小写字母
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

func (ca CommandManage) CheckSeparator(cs string) string{ //检查分隔符是用“/”还是“\”，将输入命令中的分隔符统一转化成“\”
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



var CommandList = [13]string{dir, md , rd, cd, touch, del, copy, ren, move, mklink, save, load, cls}

type CommandAnalyze struct {
	ParameterPath string
	Command string
}

func (sc CommandAnalyze) SeparateCommand(cs string) (string,string){
	for _,v := range CommandList{
		if strings.HasPrefix(cs, v) { //判断字符串是不是以字符串v开头
			csr := strings.Replace(cs, " ", "",-1)
			return strings.TrimPrefix(csr,v ),v //返回去掉v的字符串
		}
	}
	return "1","1"
}

type ParameterAnalyze struct { //包含了参数/s,/ad,/y
	path string
	Parameter string
}

var ParameterList = map[string][]string{
	dir : {ad,s},
	del : {s},
	rd : {s},
	move : {y},
}

func (pa ParameterAnalyze) SeparateParameter(pp string, c string)(string,string) {
	command := c
	_,t  := ParameterList[command]
	if t {
		for _, value := range ParameterList {
			var ok bool
			value, ok = ParameterList[command]
			if ok == true {
				for _,v := range value{
					if strings.HasPrefix(pp,v){
						ppr := strings.Replace(pp, " ", "",-1)
						return strings.TrimPrefix(ppr, v), v
						break
					}
				}
				return pp,"没有参数"
			}
		}
	}
	return "sorry","这个命令没有参数"
}

func (ca CommandManage) IsRelaivePath(cs string) bool{ //检查是相对路径还是绝对路径，返回值为ture为相对路径，返回值为false为绝对路径
	drive :=fmt.Sprint(GetRootDrive()+GetSeparatorChar())
	if  strings.Contains(cs,drive) {//判断是否包含盘符
		return false
	}else{
		return true
	}
}

func (ca CommandManage) IsTrueDiskPath(cs string) bool{ //检查是否是真实磁盘的路径，返回值为ture为真实磁盘路径，返回值为false为虚拟磁盘路径
	if  strings.Contains(cs, "@") {
		return true
	}else{
		return false
	}
}

func (ca CommandManage) HaveWildCard(cs string) bool{ //检查是否含有通配符，返回值为ture为含有通配符，返回值为false为不含通配符
	if  strings.ContainsAny(cs, "*|?") {
		return true
	}else{
		return false
	}
}


func (ca CommandManage) CheckFormat(cs string) (bool, bool, bool){
	ca.CheckSize(cs)
	ca.CheckCase(cs)
	ca.CheckSeparator(cs)
	itp := ca.IsTrueDiskPath(cs)
	irp := ca.IsRelaivePath(cs)
	hwc := ca.HaveWildCard(cs)
	return itp, irp, hwc
}

type TypeAnalyze struct {
	//将命令分为8种，非真实磁盘，无通配符，绝对路径径返回0；非真实磁盘，无通配符，相对路径，返回1；非真实磁盘，有通配符，绝对路径，返回2；非真实磁盘，有通配符，相对路径，返回3；
	//，真实磁盘，无通配符，绝对路径径返回4；真实磁盘，无通配符，相对路径，返回5；真实磁盘，有通配符，绝对路径，返回6非真实磁盘，有通配符，相对路径，返回7
	Commandtype int
}

func (ta TypeAnalyze) GetCommandType(itp bool, hwc bool,irp bool) int{
	if(itp == false && hwc == false && irp == false){
		return 0
	} else if(itp == false && hwc == false && irp == true){
		return 1
	}else if(itp == false && hwc == true && irp == false){
		return 2
	}else if(itp == false && hwc == true && irp == true){
		return 3
	}else if(itp == true && hwc == false && irp == false){
		return 4
	}else if(itp == true && hwc == false && irp == true){
		return 5
	}else if(itp == true && hwc == true && irp == false){
		return 6
	}else if(itp == true && hwc == true && irp == true){
		return 7
	}
	return 1
}

type Split struct {
	PathElement []string
}

func (s Split) SplitCommand(cs string) []string{
	csl := strings.Split(cs,GetSeparatorChar())
	return csl
}







