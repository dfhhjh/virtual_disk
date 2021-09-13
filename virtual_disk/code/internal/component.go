package internal

import "fmt"

type Path struct{//路径定义为双向链表
	Prior *Path
	Name string
	Next *Path
}

type Component struct{
	Name           string
	CreateTime     string
	ReviseTime     string
	IsFolder       bool //IsFolder = true 为目录Folder，IsFolder= false 为文件File
	Path           *Path
	FatherContents *Folder
}


type ComponentManage interface {
	AddNode()
	RemoveNode()
	GetName() string
	SetName() string
	GetCreateTime() string
	SetCreateTime() string
	GetReviseTime() string
	SetReviseTime() string
	Getpath() string
}

func (c Component) Getpath()  {
	var currentpath []string
	if c.Path != nil {
		for i := range currentpath{
			currentpath[i] = c.Path.Name
			c.Path = c.Path.Next
		}
	}
	PrintPath(c.Path)
}

func (c Component) Setpath() {
	var currentpath []string
		for i := range currentpath{
			if c.FatherContents != nil{
				currentpath[len(currentpath)-1-i] = c.Path.Name
				c.Path.Prior = c.FatherContents.Path
				c.Path = c.Path.Prior
			} else{
				c.Path.Name = GetRootDrive()
				c.Path.Prior = nil
				break
			}
	}

}

func (c Component) AddNode(CommandPath  *Path, WorkingPath  *Path) {
	//Check(CommandPath)
	var pathname *string
	pathname = &WorkingPath.Name
	if CommandPath != nil {
		pathname = &CommandPath.Name
		WorkingPath.Name = *pathname
		CommandPath = CommandPath.Next
		WorkingPath = WorkingPath.Next
	}

}

func (c Component) RemoveNode(CommandPath  *Path, WorkingPath  *Path) {
	//Check(CommandPath)
	if CommandPath != nil {
		if 	WorkingPath.Name == CommandPath.Name{
			WorkingPath.Name = ""
			CommandPath = CommandPath.Next
			WorkingPath = WorkingPath.Next
		} else{
			fmt.Println("抱歉，输入的路径不存在，请重新输入路径：")
		}
	}
}

func (c Component) GetName() string{
	return c.Name
}

func (c Component) SetName(name string) {
	c.Name = name
}

func (c Component) GetCreateTime() string{
	return c.CreateTime
}

func (c *Component) SetCreateTime(time string) {
	c.CreateTime = time
}

func (c Component) GetReviseTime() string{
	return c.ReviseTime
}

func (c *Component) SetReviseTime(time string) {
	c.ReviseTime = time
}

type Folder struct {
	BrotherContents *Folder
	SonContents *Folder
	Component
}


type File struct {
	Content byte
	FileLength int
	Component
}

type FileManage interface {
	GetFileLength()
	SetFileLength()
	GetFileContent()
	SetFileContent()
}

func (f File) GetFileLength() int{
	return f.FileLength
}

func (f File) SetFileLength(length int) {
	f.FileLength = length
}

func (f File) GetFileContent() byte{

	return f.Content
}

func (f File) SetFileContent(content byte) {
	f.Content = content
}

type Link struct {
	Component
}


