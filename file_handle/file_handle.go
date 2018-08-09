/*文件处理基本操作*/
package main

import (
	"os"
	"fmt"
	"io/ioutil"
)

func check(e error){
	if e !=nil{
		panic(e)
	}
}
/*文件处理两种方式，os,io/ioutil，处理操作包含文件的读写，增删减改*/
func main (){
	f1,err := os.Create("test.txt")
	check(err)
	defer f1.Close()
	f1.WriteString("mkwdkl\n")
	st :=[]byte{50,97,65,80}
	fmt.Println(st[0])
	f1.Write(st)
	f1.Close()

	f2,err :=os.Open("test.txt")
	check(err)
	ss := make([]byte,100)
	f2.Read(ss)
	fmt.Println(ss)
	fmt.Printf("%s\n",ss)
	f2.Close()

	f3,err :=ioutil.ReadFile("test.txt")
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Println(f3)
	fmt.Println(string(f3))
	fmt.Printf("%s\n",f3)

	os.Create("b.txt")
	ioutil.WriteFile("b.txt",st,0064)

	err =os.Truncate("b.txt",2)
	check(err)

	ff,err := os.OpenFile("test.txt",os.O_APPEND,0666)
	ff.WriteString("\n123456748960")
	ff.Close()
}
