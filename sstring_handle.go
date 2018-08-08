package main

import (
	"io/ioutil"
	"fmt"
	"strings"
	"strconv"
)

func check(err error){
	if err!= nil{
		panic(err)
	}
}
func main(){
	b,err :=ioutil.ReadFile("a.txt")
	if err!= nil{
		fmt.Println(err)

	}
	str := string(b)
	fmt.Println(str)
	fmt.Printf("pre string %t \n",strings.HasPrefix(str,":vfk"))
	fmt.Printf("the suffix is %t\n",strings.HasSuffix(str,"fmm"))
	fmt.Printf("the position %d\n",strings.Index(str,"cfn"))
	fmt.Printf("the contain %t \n",strings.Contains(str,"cfn"))
	b2 := strings.Replace(str,"cfn","y",20)
	fmt.Println(b2)
	fmt.Println(strings.Count(str,"f"))
	fmt.Println(strings.Repeat(str[0:8],3))
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(strings.ToUpper(str)))
	fmt.Println(len(strings.Fields(str)))
	fmt.Println(strings.Split(str,"cfn"))
	fmt.Println(strings.Trim(str, "vm" ))
	fmt.Println("the size of int is :",strconv.IntSize)
	fmt.Println(strconv.Itoa(3))
	fmt.Println(strconv.Atoi("23890"))
	stt := strings.Split(str," ")
	fmt.Println(stt)
	 ss :=strings.Join(stt," yes ")
	fmt.Println(ss)
	 sss :="asj"+"dqn"+"iii"
	 fmt.Println(sss)
	 fmt.Println("**********************")
	 ioutil.WriteFile("a.txt",[]byte(ss),0064)
}
