/*程序默认本地压缩文件为test.zip，用于修改压缩文件的内容，并生成符合要求的新型压缩文件*/
package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
	)
//解压zip压缩文件，输出解压后的文件名字
func unzipfile(pathname string)[]string{
	red, err := zip.OpenReader(pathname)
	if err !=nil{
		fmt.Println(err)
	}
	var ss []string
	for _,fil := range red.File{
	    ss = append(ss,fil.Name)
		f,_ :=fil.Open()
		str := make([]byte,100)
		f.Read(str)
		ff,err :=os.Create(fil.Name)
		if err != nil{
			fmt.Println(err)
		}
		ff.Write(str)
		ff.Close()
	}
	return ss
}
//对本地的文件进行字符串替换处理
func filecontent(filename []string){
	for _,name :=range filename{
		str,err :=ioutil.ReadFile(name)
		if err !=nil{
			fmt.Println(err)
		}
		for i := 0; i < len(str); i++ {
			if str[i] == 114 && str[i+1] == 32 && str[i+2] >= 48 && str[i+2] <= 57 {
				str[i] = 42
				str[i+1] = 42
				str[i+2] = 42
				i = i + 2
			}
		}
		ioutil.WriteFile(name,str,0064)
	}
}
//压缩本地文件，并输出压缩文件名
func zipfile(filestring []string)string{
	f,_ := os.Create("test3.zip")
	w :=zip.NewWriter(f)
	for _,name :=range filestring{
		fw,err :=w.Create(name)
		if err!=nil{
			fmt.Println(err)
		}
		n,err :=ioutil.ReadFile(name)
		if err!=nil{
			fmt.Println(err)
		}
		fw.Write(n)
	}
	w.Flush()
	w.Close()
	f.Close()
	return f.Name()
}
func main (){
	filename1 :=unzipfile("test.zip")
	filecontent(filename1)
	name :=zipfile(filename1)
	fmt.Printf("修改后的压缩文件为：%s\n",name)
}
