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

	ioutil.WriteFile("b.txt",st,0064)

	cf,err :=os.Create("test2.txt")
	if err !=nil{
		fmt.Println(err)
	}
	//log.Println(cf)
	cf.Close()

	err =os.Truncate("a.txt",20)
	check(err)

	ff,err := os.OpenFile("test.txt",os.O_APPEND,0666)
	ff.WriteString("\n123456748960")
	ff.Close()
}