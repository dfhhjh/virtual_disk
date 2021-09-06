package Class


type VirtualDiskManageSystem struct {
	virtualdisk VirtualDisk
}//虚拟磁盘管理系统类

func (vdms VirtualDiskManageSystem) VirtualDiskManage() {

}

type VirtualDisk struct {
	CommandManageSystem CommandManage
	RootContents Contents
	WorkingPath []string
}//虚拟磁盘类

func (vd VirtualDisk) Execute() {

}

func (vd VirtualDisk)  VirtualDiskInit(){

}

type Component struct{
	Name string
	CreateTime string
	ReviseTime string
	ctype int
	path []string
	FatherContents *Contents
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
}

func (c Component) AddNode() {

}

func (c Component) RemoveNode() {

}

func (c Component) GetName() {

}

func (c Component) SetName() {

}

func (c Component) GetCreateTime() {

}

func (c Component) SetCreateTime() {

}

func (c Component) GetReviseTime() {

}

func (c Component) SetReviseTime() {

}

type Contents struct {
	BrotherContents *Contents
	SonContents *Contents
	Component
}

type File struct {
	Content string
	FileLength int
	Component
}

type FileManage interface {
	GetFileLength()
	SetFileLength()
	GetFileContent()
	SetFileContent()
}

func (f File) GetFileLength() {

}

func (f File) SetFileLength() {

}

func (f File) GetFileContent() {

}

func (f File) SetFileContent() {

}

type Link struct {
	Component
}

type CommandManage struct {
	CommandString []string
	Split
	TypeAnalyze
	ParameterAnalyze
	PathAnalyze
	FormatAnalyze
}

func (cn CommandManage) CreateCommand() {

}

type Split struct {

}

func (s Split) CreateCommand() {

}