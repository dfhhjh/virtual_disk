package internal

import "errors"

func SizeError() error{ //字符过多提醒错误
	 return errors.New("输入的字符数过多，请重新输入")
}

func InvalidCharactersError() error{ //含有非法字符字符，提醒错误
	return errors.New("输入含有非法字符，请重新输入")
}

func WildCardError() error{ //该命令不含有通配符，提醒错误
	return errors.New("该命令不含有通配符，请重新输入")
}

func TrueDiskError() error{//该命令不支持真实磁盘路径，提醒错误
	return errors.New("该命令不支持真实磁盘路径，请重新输入")
}

func PathError() error{//该命令路径有误 ，提醒错误
	return errors.New("该命令路径有误，请重新输入")
}