package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func main() {

	fileData, err := ioutil.ReadFile("E:\\go_code\\virtual_disk\\copy\\pic.png")
	fmt.Println(err)

	fileString := base64.StdEncoding.EncodeToString(fileData)

	b, c := base64.StdEncoding.DecodeString(fileString)
	fmt.Print(b, c)
}





